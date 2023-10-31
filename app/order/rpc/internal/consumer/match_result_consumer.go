package consumer

import (
	"context"
	"github.com/luxun9527/gex/app/order/rpc/internal/logic"
	"github.com/luxun9527/gex/app/order/rpc/internal/svc"
	"github.com/luxun9527/gex/common/pkg/logger"
	matchMq "github.com/luxun9527/gex/common/proto/mq/match"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
)

func InitConsumer(sc *svc.ServiceContext) {
	matchResultHandler := logic.NewHandleMatchResultLogic(sc)
	go func() {
		for {
			message, err := sc.MatchConsumer.Receive(context.Background())
			if err != nil {
				logx.Errorw("consumer message match result failed", logger.ErrorField(err))
				continue
			}

			var m matchMq.MatchResp
			if err := proto.Unmarshal(message.Payload(), &m); err != nil {
				logx.Errorw("unmarshal match result failed", logger.ErrorField(err))
				if err := sc.MatchConsumer.Ack(message); err != nil {
					logx.Errorw("consumer message failed", logger.ErrorField(err))
				}
				continue
			}

			//todo 防重复提交校验。
			switch r := m.Resp.(type) {
			case *matchMq.MatchResp_MatchResult:

				logx.Infow("receive match result data ", logx.Field("data", r))
				//修改订单状态，插入到成交表中，修改用户的资产
				if err := matchResultHandler.HandleMatchResult(r); err != nil {
					logx.Severef("handle match result failed", logger.ErrorField(err), logx.Field("data", r))
				}

			case *matchMq.MatchResp_Cancel:
				logx.Infow("receive match cancel data ", logx.Field("data", r))
				if err := matchResultHandler.CancelOrder(r); err != nil {
					logx.Severef("[consumer] handle cancel order message failed", logger.ErrorField(err), logx.Field("data", r))
				}
			}
			if err := sc.MatchConsumer.Ack(message); err != nil {
				logx.Errorw("ack message failed", logger.ErrorField(err))
			}

		}

	}()
}
