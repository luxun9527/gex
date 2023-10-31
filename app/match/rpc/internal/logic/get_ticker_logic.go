package logic

import (
	"context"
	"encoding/json"
	"github.com/luxun9527/gex/app/match/rpc/internal/dao/model"
	"github.com/luxun9527/gex/app/match/rpc/internal/svc"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/luxun9527/gex/common/proto/define"
	"github.com/luxun9527/gex/common/utils"

	"github.com/luxun9527/gex/app/match/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTickerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTickerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTickerLogic {
	return &GetTickerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取ticker
func (l *GetTickerLogic) GetTicker(in *pb.GetTickerReq) (*pb.GetTickerResp, error) {
	quoteCoinPrec := l.svcCtx.Config.SymbolInfo.QuoteCoinPrec
	baseCoinPrec := l.svcCtx.Config.SymbolInfo.BaseCoinPrec
	resp := &pb.GetTickerResp{}
	if in.Symbol == "" {
		data, err := l.svcCtx.RedisClient.Hgetall(string(define.Ticker))
		respData := make([]*pb.GetTickerResp_Ticker, 0, len(data))
		if err != nil {
			logx.Errorw("query from redis failed", logger.ErrorField(err))
			return nil, errs.RedisFailed
		}
		for _, v := range data {
			var tickerRedisData model.TickerRedisData
			if err := json.Unmarshal([]byte(v), &tickerRedisData); err != nil {
				return nil, errs.Internal
			}
			d := &pb.GetTickerResp_Ticker{
				LatestPrice: utils.PrecCut(tickerRedisData.Price, quoteCoinPrec),
				High:        utils.PrecCut(tickerRedisData.High, quoteCoinPrec),
				Low:         utils.PrecCut(tickerRedisData.Low, quoteCoinPrec),
				Volume:      utils.PrecCut(tickerRedisData.Price, quoteCoinPrec),
				Amount:      utils.PrecCut(tickerRedisData.Turnover, baseCoinPrec),
				Last24Price: utils.PrecCut(tickerRedisData.Last24, quoteCoinPrec),
				PriceRange:  tickerRedisData.Range,
				Symbol:      tickerRedisData.Symbol,
			}
			respData = append(respData, d)
		}
		resp.TickerList = respData
	} else {
		var tickerRedisData model.TickerRedisData
		data, err := l.svcCtx.RedisClient.Hget(string(define.Ticker), in.Symbol)
		if err != nil {
			logx.Errorw("query from redis failed", logger.ErrorField(err))
			return nil, errs.RedisFailed
		}
		respData := make([]*pb.GetTickerResp_Ticker, 0, len(data))
		if err := json.Unmarshal([]byte(data), &tickerRedisData); err != nil {
			return nil, errs.Internal
		}
		d := &pb.GetTickerResp_Ticker{
			LatestPrice: utils.PrecCut(tickerRedisData.Price, quoteCoinPrec),
			High:        utils.PrecCut(tickerRedisData.High, quoteCoinPrec),
			Low:         utils.PrecCut(tickerRedisData.Low, quoteCoinPrec),
			Volume:      utils.PrecCut(tickerRedisData.Price, quoteCoinPrec),
			Amount:      utils.PrecCut(tickerRedisData.Turnover, baseCoinPrec),
			Last24Price: utils.PrecCut(tickerRedisData.Last24, quoteCoinPrec),
			PriceRange:  tickerRedisData.Range,
			Symbol:      tickerRedisData.Symbol,
		}
		respData = append(respData, d)
		resp.TickerList = respData
	}

	return resp, nil
}
