package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/internal/consumer"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/internal/dao/query"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/internal/model"
	"github.com/luxun9527/gex/app/quotes/kline/rpc/internal/svc"
	"github.com/luxun9527/gex/common/proto/define"
	commonWs "github.com/luxun9527/gex/common/proto/ws"
	"github.com/luxun9527/gex/common/utils"
	gpush "github.com/luxun9527/gpush/proto"
	logger "github.com/luxun9527/zaplog"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm/clause"
	"time"
)

// KlineHandler 基于utc时间
type KlineHandler struct {
	//存储k线
	Klines []*model.Kline
	//落库chan
	storeLatestKline chan model.StoreKline
	//发送
	sendChan chan model.Kline
	//定时写入和发送的定时器
	ticker *time.Ticker
	//是否改变
	changed bool
	//提交方式

	cron            *utils.WrapCron
	svcCtx          *svc.ServiceContext
	latestMessageID pulsar.MessageID
	latestMatchId   int64
}

func NewKlineHandler(svcCtx *svc.ServiceContext) *KlineHandler {
	klineHandler := &KlineHandler{
		storeLatestKline: make(chan model.StoreKline),
		sendChan:         make(chan model.Kline),
		ticker:           time.NewTicker(300 * time.Millisecond),
		svcCtx:           svcCtx,
	}

	wrapCron, err := utils.NewWrapCron("1 * * * * ?")
	if err != nil {
		logx.Severef("init cron failed %v", err)
	}
	klineHandler.cron = wrapCron
	return klineHandler
}
func InitKlineHandler(svcCtx *svc.ServiceContext) {
	handler := NewKlineHandler(svcCtx)
	handler.readInitData()
	handler.cron.Start()
	matchDataChan := consumer.InitConsumer(handler.svcCtx)
	go handler.update(matchDataChan)
	go handler.store()
	go handler.send()
}
func (kl *KlineHandler) readInitData() {
	klines := make([]*model.Kline, 0, len(model.KlineTypes))

	for _, v := range model.KlineTypes {

		data, err := kl.svcCtx.RedisClient.Hget(define.Kline.WithParams(), kl.svcCtx.Config.SymbolInfo.SymbolName+"_"+v.String())
		if err != nil {
			if errors.Is(err, redis.Nil) {
				kline := &model.Kline{
					StartTime: 0,
					EndTime:   0,
					KlineType: v,
					Amount:    utils.DecimalZeroMaxPrec,
					Volume:    utils.DecimalZeroMaxPrec,
					Open:      utils.DecimalZeroMaxPrec,
					High:      utils.DecimalZeroMaxPrec,
					Low:       utils.DecimalZeroMaxPrec,
					Close:     utils.DecimalZeroMaxPrec,
					Range:     "0",
				}
				klines = append(klines, kline)
				continue
			}
			logx.Severef("read init kline data failed err=%v", err)
		}
		var d model.RedisModel
		if err := json.Unmarshal([]byte(data), &d); err != nil {
			logx.Severef("unmarshal kline data failed err=%v", err)
		}

		kline := &model.Kline{
			KlineType: model.KlineType(d.KlineType),
			StartTime: d.StartTime,
			EndTime:   d.EndTime,
			Amount:    utils.NewFromStringMaxPrec(d.Volume),
			Volume:    utils.NewFromStringMaxPrec(d.Amount),
			Open:      utils.NewFromStringMaxPrec(d.Open),
			High:      utils.NewFromStringMaxPrec(d.High),
			Low:       utils.NewFromStringMaxPrec(d.Low),
			Close:     utils.NewFromStringMaxPrec(d.Close),
			Range:     d.Range,
		}
		klines = append(klines, kline)
	}
	kl.Klines = klines

}
func (kl *KlineHandler) update(matchData <-chan *model.MatchData) {
	for {
		select {
		case md := <-matchData:
			kl.updateLatestKline(md, false)
			kl.changed = true
			kl.latestMessageID = md.MessageID
			kl.latestMatchId = md.MatchID
		case <-kl.ticker.C:
			if kl.changed {
				kl.snapshot()
			}
			kl.changed = false
		case <-kl.cron.C:
			kl.updateLatestKline(nil, true)
			kl.changed = true
		}
	}
}

// 存储历史k线和最新的k线
func (kl *KlineHandler) store() {
	for klineData := range kl.storeLatestKline {
		//存储历史k线
		if klineData.IsHistory {
			err := kl.svcCtx.Query.Transaction(func(tx *query.Query) error {
				for _, v := range klineData.Klines {
					mysqlData := v.CastToMysqlData(kl.svcCtx.Config.SymbolInfo)
					logx.Infow("store history kline data", logx.Field("data", mysqlData))
					if err := kl.svcCtx.Query.Kline.WithContext(context.Background()).
						Clauses(clause.Insert{Modifier: "IGNORE"}).
						Create(mysqlData); err != nil {
						return err
					}
				}

				return nil
			})
			if err != nil {
				logx.Severef("store message to mysql failed err = %v message id %v", err, kl.latestMessageID)
			}

		} else {
			if err := kl.svcCtx.Query.Transaction(func(tx *query.Query) error {
				for _, v := range klineData.Klines {
					data := v.CastToRedisModelData(kl.svcCtx.Config.SymbolInfo, klineData.MatchID)
					d, _ := json.Marshal(data)
					if err := kl.svcCtx.RedisClient.Hset(define.Kline.WithParams(), data.Symbol+"_"+v.KlineType.String(), string(d)); err != nil {
						logx.Errorw("update last kline failed", logger.ErrorField(err))
						return err
					}
				}

				if klineData.MessageID != nil {
					if err := kl.svcCtx.MatchConsumer.AckIDCumulative(kl.latestMessageID); err != nil {
						logx.Errorw("consumer message failed", logger.ErrorField(err), logx.Field("messageID", kl.latestMessageID))
						return err
					}
				}

				return nil
			}); err != nil {
				logx.Severef("store last kline failed err=%v", err)
			}
		}

	}
}

func (kl *KlineHandler) snapshot() {
	latestKline := make([]*model.Kline, 0, len(kl.Klines))
	for _, v := range kl.Klines {
		t := *v
		kl.sendChan <- t
		latestKline = append(latestKline, &t)

	}
	//定时存储最新的一根k线
	l := model.StoreKline{
		Klines:    latestKline,
		MessageID: kl.latestMessageID,
		MatchID:   kl.latestMatchId,
	}
	kl.storeLatestKline <- l
}
func (kl *KlineHandler) send() {
	for data := range kl.sendChan {
		msg := commonWs.Message[commonWs.Kline]{
			Topic:   commonWs.KlinePrefix.WithParam(kl.svcCtx.Config.SymbolInfo.SymbolName) + "@" + data.KlineType.String(),
			Payload: data.CastToWsData(kl.svcCtx.Config.SymbolInfo),
		}
		if _, err := kl.svcCtx.WsClient.PushData(context.Background(), &gpush.Data{
			Uid:   "",
			Topic: commonWs.KlinePrefix.WithParam(kl.svcCtx.Config.SymbolInfo.SymbolName) + "@" + data.KlineType.String(),
			Data:  msg.ToBytes(),
		}); err != nil {
			logx.Errorw("push kline websocket data failed", logger.ErrorField(err), logx.Field("data", msg))
		}
	}
}

// 更新最新的k线
func (kl *KlineHandler) updateLatestKline(data *model.MatchData, isBegin bool) {
	logx.Infow("receive match data ", logx.Field("data", data))
	for _, klineData := range kl.Klines {
		logx.Infow("before update ", logx.Field("klineData", klineData.CastToMysqlData(kl.svcCtx.Config.SymbolInfo)))
		//如果是mock撮合用最新的价格计算
		if isBegin {
			//价格为零不计算
			if klineData.Close.Equal(utils.DecimalZeroMaxPrec) {
				return
			}
			data = &model.MatchData{}
			data.MatchTime = time.Now().Unix()
			data.Amount = utils.DecimalZeroMaxPrec
			data.Volume = utils.DecimalZeroMaxPrec
			data.High = klineData.Close
			data.Low = klineData.Close
			data.StartPrice = klineData.Close
			data.EndPrice = klineData.Close
		}

		var (
			startTime,
			endTime int64
		)
		//修正交易时间为一个新的区间,如5分钟k线，交易时间为 06:23 则修改其为 05:00
		switch klineData.KlineType {
		case model.Week1:
			startTime = utils.BeginOfWeek(data.MatchTime)
			endTime = startTime + int64(klineData.KlineType.GetCycle())
		case model.Month1:
			startTime = utils.BeginOfMonth(data.MatchTime)
			endTime = utils.NextMonth(startTime)
		default:
			//去掉时间戳的余数
			startTime = data.MatchTime / int64(klineData.KlineType.GetCycle()) * int64(klineData.KlineType.GetCycle())
			endTime = startTime + int64(klineData.KlineType.GetCycle())
		}
		//初始化k线
		if klineData.Open.Equal(utils.DecimalZeroMaxPrec) {
			klineData.Open = data.StartPrice
			klineData.StartTime = startTime
			klineData.EndTime = endTime
			klineData.High = data.High
			klineData.Low = data.Low
			klineData.Close = data.EndPrice
			klineData.Amount = data.Amount
			klineData.Volume = data.Volume
			klineData.Range = "0"
		}
		//如果k线在一个新的周期
		if startTime > klineData.StartTime && startTime > 0 {
			//存储历史k线
			//发送到发送和写的chan
			historyKline := *klineData

			//返回修改为最新的k线
			klineData.Open = data.StartPrice
			klineData.StartTime = startTime
			klineData.EndTime = endTime
			klineData.High = data.High
			klineData.Low = data.Low
			klineData.Close = data.EndPrice
			klineData.Amount = data.Amount
			klineData.Volume = data.Volume
			if !klineData.Open.Equal(utils.DecimalZeroMaxPrec) {
				klineData.Range = data.EndPrice.Sub(klineData.Open).Div(klineData.Open).Mul(utils.NewFromStringMaxPrec("100")).StringFixedBank(3)
			}
			newKline := *klineData
			//k线
			sk := model.StoreKline{
				Klines:    []*model.Kline{&historyKline, &newKline},
				MessageID: data.MessageID,
				IsHistory: true,
			}
			kl.sendChan <- historyKline
			kl.sendChan <- newKline
			kl.storeLatestKline <- sk
			continue
		}
		//比较高低，累加成交量成交额
		klineData.StartTime = startTime
		klineData.EndTime = endTime
		klineData.Amount = klineData.Amount.Add(data.Amount)
		klineData.Volume = klineData.Volume.Add(data.Volume)
		if data.High.GreaterThan(klineData.High) {
			klineData.High = data.High
		}
		if data.Low.LessThan(klineData.Low) {
			klineData.Low = data.Low
		}
		if !klineData.Open.Equal(utils.DecimalZeroMaxPrec) {
			klineData.Range = data.EndPrice.Sub(klineData.Open).Div(klineData.Open).Mul(utils.NewFromStringMaxPrec("100")).StringFixedBank(3)
		}
		klineData.Close = data.EndPrice
		logx.Debugw("after update ", logx.Field("klineData", klineData.CastToMysqlData(kl.svcCtx.Config.SymbolInfo)))

	}
}
