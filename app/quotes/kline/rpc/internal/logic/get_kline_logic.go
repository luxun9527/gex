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
		Where(kline.Symbol.Eq(l.svcCtx.Config.SymbolInfo.SymbolName), kline.KlineType.Eq(int32(in.KlineType)), kline.StartTime.Between(in.StartTime, in.EntTime)).
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
			Open:      utils.PrecCut(klineList[i].Open, l.svcCtx.Config.SymbolInfo.QuoteCoinPrec),
			High:      utils.PrecCut(klineList[i].High, l.svcCtx.Config.SymbolInfo.QuoteCoinPrec),
			Low:       utils.PrecCut(klineList[i].Low, l.svcCtx.Config.SymbolInfo.QuoteCoinPrec),
			Close:     utils.PrecCut(klineList[i].Close, l.svcCtx.Config.SymbolInfo.QuoteCoinPrec),
			Volume:    utils.PrecCut(klineList[i].Volume, l.svcCtx.Config.SymbolInfo.QuoteCoinPrec),
			Amount:    utils.PrecCut(klineList[i].Amount, l.svcCtx.Config.SymbolInfo.BaseCoinPrec),
			Range:     klineList[i].Range,
		}
		klineResp = append(klineResp, d)
	}
	return &pb.GetKlineResp{KlineList: klineResp}, nil
}
