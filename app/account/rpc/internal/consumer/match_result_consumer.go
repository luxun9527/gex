package consumer

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/luxun9527/gex/app/account/rpc/internal/logic"
	"github.com/luxun9527/gex/app/account/rpc/internal/svc"
	"github.com/luxun9527/gex/common/pkg/logger"
	matchMq "github.com/luxun9527/gex/common/proto/mq/match"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
)

type MatchResultConsumer struct {
	sc *svc.ServiceContext
}

func InitConsumer(sc *svc.ServiceContext) {
	for _, consumer := range sc.MatchConsumerList {
		go func(c pulsar.Consumer) {
			for {
				message, err := c.Receive(context.Background())
				if err != nil {
					logx.Errorw("consumer message match result failed", logger.ErrorField(err))
					continue
				}
				var m matchMq.MatchResp
				if err := proto.Unmarshal(message.Payload(), &m); err != nil {
					logx.Errorw("unmarshal match result failed", logger.ErrorField(err))
					if err := c.Ack(message); err != nil {
						logx.Errorw("consumer message failed", logger.ErrorField(err))
					}
					continue
				}
				//todo 防重复提交校验。
				switch r := m.Resp.(type) {
				case *matchMq.MatchResp_MatchResult:
					logx.Infow("receive match result data ", logx.Field("data", r))
					//修改订单状态，插入到成交表中，修改用户的资产
					if err := logic.NewHandleMatchResultLogic(sc).HandleMatchResult(r); err != nil {
						logx.Severef("[consumer]handle match result failed", logger.ErrorField(err), logx.Field("data", r))
					}
				case *matchMq.MatchResp_Cancel:
					logx.Infow("receive match cancel data ", logx.Field("data", r))
					if err := logic.NewHandleMatchResultLogic(sc).HandleCancelOrder(r); err != nil {
						logx.Severef("[consumer]  match result cancel order failed", logger.ErrorField(err), logx.Field("data", r))
					}
					//解冻用户资产
				}
				if err := c.Ack(message); err != nil {
					logx.Errorw("ack message failed", logger.ErrorField(err))
				}
			}

		}(consumer)
	}

}
