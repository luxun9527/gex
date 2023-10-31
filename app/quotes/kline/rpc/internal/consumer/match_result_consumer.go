package consumer

import (
	"context"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/internal/model"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/internal/svc"
	"github.com/luxun9527/gex/common/pkg/logger"
	matchMq "github.com/luxun9527/gex/common/proto/mq/match"
	"github.com/luxun9527/gex/common/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
)

func InitConsumer(sc *svc.ServiceContext) <-chan *model.MatchData {
	md := make(chan *model.MatchData)
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
				matchData := &model.MatchData{
					MessageID:  message.ID(),
					MatchTime:  r.MatchResult.MatchTime / 1e9,
					Volume:     utils.NewFromStringMaxPrec(r.MatchResult.Amount).Mul(utils.NewFromStringMaxPrec("2")),
					Amount:     utils.NewFromStringMaxPrec(r.MatchResult.Qty).Mul(utils.NewFromStringMaxPrec("2")),
					StartPrice: utils.NewFromStringMaxPrec(r.MatchResult.BeginPrice),
					EndPrice:   utils.NewFromStringMaxPrec(r.MatchResult.EndPrice),
					Low:        utils.NewFromStringMaxPrec(r.MatchResult.LowPrice),
					High:       utils.NewFromStringMaxPrec(r.MatchResult.HighPrice),
				}
				md <- matchData
			}
			//message.ID()
			//sc.MatchConsumer.AckIDCumulative()
			//if err := sc.MatchConsumer.Ack(message); err != nil {
			//	logx.Errorw("ack message failed", logger.ErrorField(err))
			//}
		}

	}()
	return md
}
