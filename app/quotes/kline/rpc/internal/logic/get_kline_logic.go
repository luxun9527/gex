package logic

import (
	"context"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/internal/svc"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/pb"
	"github.com/luxun9527/gex/common/errs"
	"github.com/luxun9527/gex/common/pkg/logger"
	"github.com/luxun9527/gex/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetKlineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetKlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetKlineLogic {
	return &GetKlineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取k线
func (l *GetKlineLogic) GetKline(in *pb.GetKlineReq) (*pb.GetKlineResp, error) {
	kline := l.svcCtx.Query.Kline
	klineList, err := kline.WithContext(l.ctx).
		Where(kline.Symbol.Eq(l.svcCtx.SymbolInfo.SymbolName), kline.KlineType.Eq(int32(in.KlineType)), kline.StartTime.Between(in.StartTime, in.EntTime)).
		Limit(300).
		Order(kline.StartTime.Desc()).
		Find()
	if err != nil {
		logx.Errorw("query  kline from mysql failed", logger.ErrorField(err))
		return nil, errs.ExecSqlFailed
	}
	klineResp := make([]*pb.GetKlineResp_Kline, 0, len(klineList)+1)
	for i := len(klineList) - 1; i >= 0; i-- {
		d := &pb.GetKlineResp_Kline{
			StartTime: klineList[i].StartTime,
			EndTime:   klineList[i].EndTime,
			Symbol:    klineList[i].Symbol,
			Open:      utils.PrecCut(klineList[i].Open, l.svcCtx.SymbolInfo.QuoteCoinPrec.Load()),
			High:      utils.PrecCut(klineList[i].High, l.svcCtx.SymbolInfo.QuoteCoinPrec.Load()),
			Low:       utils.PrecCut(klineList[i].Low, l.svcCtx.SymbolInfo.QuoteCoinPrec.Load()),
			Close:     utils.PrecCut(klineList[i].Close, l.svcCtx.SymbolInfo.QuoteCoinPrec.Load()),
			Volume:    utils.PrecCut(klineList[i].Volume, l.svcCtx.SymbolInfo.QuoteCoinPrec.Load()),
			Amount:    utils.PrecCut(klineList[i].Amount, l.svcCtx.SymbolInfo.BaseCoinPrec.Load()),
			Range:     klineList[i].Range,
		}
		klineResp = append(klineResp, d)
	}
	return &pb.GetKlineResp{KlineList: klineResp}, nil
}
