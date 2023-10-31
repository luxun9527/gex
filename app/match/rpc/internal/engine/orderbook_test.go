package engine

import (
	"fmt"
	"github.com/luxun9527/gex/app/match/rpc/internal/config"
	"github.com/luxun9527/gex/common/proto/define"
	enum "github.com/luxun9527/gex/common/proto/enum"
	"github.com/luxun9527/gex/common/utils"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/stringx"
	"math/rand"
	"testing"
	"time"
)

func TestMatchOrderBook(t *testing.T) {
	var c config.Config
	c.SymbolInfo = define.SymbolInfo{
		SymbolName:    "BTC_USDT",
		SymbolID:      1,
		BaseCoinName:  "BTC",
		BaseCoinID:    1,
		QuoteCoinName: "USDT",
		QuoteCoinID:   1,
		BaseCoinPrec:  3,
		QuoteCoinPrec: 4,
	}
	me := NewMatchEngine(&c, nil, nil)
	//buy price 12.12 qty 100
	buyOrder1 := &Order{
		OrderID:        stringx.Randn(10),
		SequenceId:     time.Now().UnixNano(),
		CreateTime:     0,
		IsCancel:       false,
		Uid:            0,
		Price:          utils.NewFromStringMaxPrec("10"),
		Qty:            utils.NewFromStringMaxPrec("100"),
		OrderType:      enum.OrderType_LO,
		Amount:         utils.NewFromStringMaxPrec("1000"),
		Side:           enum.Side_Buy,
		OrderStatus:    0,
		UnfilledQty:    utils.NewFromStringMaxPrec("100"),
		FilledQty:      utils.DecimalZeroMaxPrec,
		UnfilledAmount: utils.NewFromStringMaxPrec("1000"),
	}
	buyOrder1.UnfilledAmount = buyOrder1.Qty.Mul(buyOrder1.Price)
	//buy price 12.12 qty 100
	time.Sleep(time.Millisecond)
	buyOrder2 := &Order{
		OrderID:        stringx.Randn(10),
		SequenceId:     time.Now().UnixNano(),
		CreateTime:     0,
		IsCancel:       false,
		Uid:            0,
		Price:          utils.NewFromStringMaxPrec("12.12"),
		Qty:            utils.NewFromStringMaxPrec("90"),
		OrderType:      enum.OrderType_LO,
		Amount:         decimal.Decimal{},
		Side:           enum.Side_Buy,
		OrderStatus:    0,
		UnfilledQty:    utils.NewFromStringMaxPrec("90"),
		FilledQty:      utils.DecimalZeroMaxPrec,
		UnfilledAmount: decimal.Decimal{},
	}
	buyOrder2.UnfilledAmount = buyOrder2.Qty.Mul(buyOrder2.Price)
	//buy price 12.12 qty 90
	time.Sleep(time.Millisecond)
	buyOrder3 := &Order{
		OrderID:     stringx.Randn(10),
		SequenceId:  time.Now().UnixNano(),
		CreateTime:  0,
		IsCancel:    false,
		Price:       utils.NewFromStringMaxPrec("12.13"),
		Qty:         utils.NewFromStringMaxPrec("120"),
		OrderType:   enum.OrderType_LO,
		Side:        enum.Side_Buy,
		OrderStatus: 0,
		UnfilledQty: utils.NewFromStringMaxPrec("120"),
		FilledQty:   utils.DecimalZeroMaxPrec,
	}
	buyOrder3.UnfilledAmount = buyOrder3.Qty.Mul(buyOrder3.Price)
	//buy price 12.13 qty 120
	time.Sleep(time.Millisecond)
	sellOrder1 := &Order{
		OrderID:     stringx.Randn(10),
		SequenceId:  time.Now().UnixNano(),
		CreateTime:  0,
		IsCancel:    false,
		Price:       utils.NewFromStringMaxPrec("12.34"),
		Qty:         utils.NewFromStringMaxPrec("120"),
		OrderType:   enum.OrderType_LO,
		Side:        enum.Side_Sell,
		OrderStatus: 0,
		UnfilledQty: utils.NewFromStringMaxPrec("120"),
		FilledQty:   utils.DecimalZeroMaxPrec,
	}
	sellOrder1.UnfilledAmount = sellOrder1.Qty.Mul(sellOrder1.Price)
	//sell price 12.34 qty 120
	time.Sleep(time.Millisecond)
	sellOrder2 := &Order{
		OrderID:     stringx.Randn(10),
		SequenceId:  time.Now().UnixNano(),
		CreateTime:  0,
		IsCancel:    false,
		Price:       utils.NewFromStringMaxPrec("12.89"),
		Qty:         utils.NewFromStringMaxPrec("120"),
		OrderType:   enum.OrderType_LO,
		Side:        enum.Side_Sell,
		OrderStatus: 0,
		UnfilledQty: utils.NewFromStringMaxPrec("120"),
		FilledQty:   utils.DecimalZeroMaxPrec,
	}
	sellOrder2.UnfilledAmount = sellOrder2.Qty.Mul(sellOrder2.Price)
	//sell price 12.89 qty 120
	time.Sleep(time.Millisecond)
	sellOrder3 := &Order{
		OrderID:     stringx.Randn(10),
		SequenceId:  time.Now().UnixNano(),
		CreateTime:  0,
		IsCancel:    false,
		Price:       utils.NewFromStringMaxPrec("12.87"),
		Qty:         utils.NewFromStringMaxPrec("1"),
		OrderType:   enum.OrderType_LO,
		Side:        enum.Side_Sell,
		OrderStatus: 0,
		UnfilledQty: utils.NewFromStringMaxPrec("1"),
		FilledQty:   utils.DecimalZeroMaxPrec,
	}
	sellOrder3.UnfilledAmount = sellOrder3.Qty.Mul(sellOrder3.Price)
	//sell price 12.87 1
	time.Sleep(time.Millisecond)
	me.HandleOrder(buyOrder1)
	me.HandleOrder(buyOrder2)
	me.HandleOrder(buyOrder3)
	me.HandleOrder(sellOrder1)
	me.HandleOrder(sellOrder2)
	me.HandleOrder(sellOrder3)
	fmt.Println(me.asks)
	fmt.Println(me.bids)
	time.Sleep(time.Millisecond * 100)
	depth := me.GetDepth(200)

	for _, v := range depth.Asks {
		fmt.Printf("[asks] Price=%v Qty=%v\n", v.Price, v.Qty)
	}
	for _, v := range depth.Bids {
		fmt.Printf("[bids] Price=%v Qty=%v\n", v.Price, v.Qty)
	}
	time.Sleep(time.Second * 10)
}

func TestMatchLO(t *testing.T) {
	var c config.Config
	c.SymbolInfo = define.SymbolInfo{
		SymbolName:    "BTC_USDT",
		SymbolID:      1,
		BaseCoinName:  "BTC",
		BaseCoinID:    1,
		QuoteCoinName: "USDT",
		QuoteCoinID:   1,
		BaseCoinPrec:  3,
		QuoteCoinPrec: 4,
	}
	me := NewMatchEngine(&c, nil, nil)
	buyOrder1 := &Order{
		OrderID:        stringx.Randn(16),
		SequenceId:     time.Now().UnixNano(),
		CreateTime:     0,
		IsCancel:       false,
		Uid:            0,
		Price:          utils.NewFromStringMaxPrec("10"),
		Qty:            utils.NewFromStringMaxPrec("100"),
		OrderType:      enum.OrderType_LO,
		Amount:         utils.NewFromStringMaxPrec("1000"),
		Side:           enum.Side_Buy,
		OrderStatus:    0,
		UnfilledQty:    utils.NewFromStringMaxPrec("100"),
		FilledQty:      utils.DecimalZeroMaxPrec,
		UnfilledAmount: utils.NewFromStringMaxPrec("1000"),
	}
	buyOrder1.UnfilledAmount = buyOrder1.Qty.Mul(buyOrder1.Price)
	time.Sleep(time.Millisecond)
	sellOrder1 := &Order{
		OrderID:        stringx.Randn(16),
		SequenceId:     time.Now().UnixNano(),
		CreateTime:     0,
		IsCancel:       false,
		Uid:            0,
		Price:          utils.NewFromStringMaxPrec("10"),
		Qty:            utils.NewFromStringMaxPrec("50"),
		OrderType:      enum.OrderType_LO,
		Amount:         utils.NewFromStringMaxPrec("1000"),
		Side:           enum.Side_Sell,
		OrderStatus:    0,
		UnfilledQty:    utils.NewFromStringMaxPrec("50"),
		FilledQty:      utils.DecimalZeroMaxPrec,
		UnfilledAmount: utils.NewFromStringMaxPrec("1000"),
	}
	sellOrder1.UnfilledAmount = sellOrder1.Qty.Mul(sellOrder1.Price)
	time.Sleep(time.Millisecond)
	sellOrder2 := &Order{
		OrderID:        stringx.Randn(16),
		SequenceId:     time.Now().UnixNano(),
		CreateTime:     0,
		IsCancel:       false,
		Uid:            0,
		Price:          utils.NewFromStringMaxPrec("10"),
		Qty:            utils.NewFromStringMaxPrec("5"),
		OrderType:      enum.OrderType_LO,
		Amount:         utils.NewFromStringMaxPrec("50"),
		Side:           enum.Side_Sell,
		OrderStatus:    0,
		UnfilledQty:    utils.NewFromStringMaxPrec("5"),
		FilledQty:      utils.DecimalZeroMaxPrec,
		UnfilledAmount: utils.NewFromStringMaxPrec("50"),
	}
	sellOrder2.UnfilledAmount = sellOrder2.Qty.Mul(sellOrder2.Price)
	time.Sleep(time.Millisecond * 100)
	me.HandleOrder(buyOrder1)
	me.HandleOrder(sellOrder1)
	me.HandleOrder(sellOrder2)
	fmt.Println(me.asks)
	fmt.Println(me.bids)
	time.Sleep(time.Millisecond * 10)
	depth := me.GetDepth(200)
	for _, v := range depth.Asks {
		fmt.Printf("[asks] Price=%v Qty=%v\n", v.Price, v.Qty)
	}
	for _, v := range depth.Bids {
		fmt.Printf("[bids] Price=%v Qty=%v\n", v.Price, v.Qty)
	}
	time.Sleep(time.Hour)
}

// 卖单撮合，吃买单
func TestMatchMO(t *testing.T) {
	var c config.Config
	c.SymbolInfo = define.SymbolInfo{
		SymbolName:    "BTC_USDT",
		SymbolID:      1,
		BaseCoinName:  "BTC",
		BaseCoinID:    1,
		QuoteCoinName: "USDT",
		QuoteCoinID:   1,
		BaseCoinPrec:  3,
		QuoteCoinPrec: 4,
	}
	me := NewMatchEngine(&c, nil, nil)
	sellOrder1 := &Order{
		OrderID:     stringx.Randn(10),
		SequenceId:  time.Now().UnixNano() + rand.Int63n(time.Now().UnixNano()),
		CreateTime:  0,
		IsCancel:    false,
		Price:       utils.NewFromStringMaxPrec("12.12"),
		Qty:         utils.NewFromStringMaxPrec("90"),
		OrderType:   enum.OrderType_LO,
		Side:        enum.Side_Sell,
		OrderStatus: 0,
		UnfilledQty: utils.NewFromStringMaxPrec("90"),
		FilledQty:   utils.DecimalZeroMaxPrec,
	}
	sellOrder1.UnfilledAmount = sellOrder1.Qty.Mul(sellOrder1.Price)
	sellOrder1.Amount = sellOrder1.Qty.Mul(sellOrder1.Price)
	time.Sleep(time.Millisecond)
	sellOrder2 := &Order{
		OrderID:     stringx.Randn(10),
		SequenceId:  time.Now().UnixNano() + rand.Int63n(time.Now().UnixNano()),
		CreateTime:  0,
		IsCancel:    false,
		Price:       utils.NewFromStringMaxPrec("12.12"),
		Qty:         utils.NewFromStringMaxPrec("5"),
		OrderType:   enum.OrderType_LO,
		Side:        enum.Side_Sell,
		OrderStatus: 0,
		UnfilledQty: utils.NewFromStringMaxPrec("5"),
		FilledQty:   utils.DecimalZeroMaxPrec,
	}
	sellOrder2.UnfilledAmount = sellOrder2.Qty.Mul(sellOrder2.Price)
	sellOrder2.Amount = sellOrder2.Qty.Mul(sellOrder1.Price)
	time.Sleep(time.Millisecond)
	buyOrder1 := &Order{
		OrderID:     stringx.Randn(10),
		SequenceId:  time.Now().UnixNano(),
		CreateTime:  0,
		IsCancel:    false,
		Price:       utils.DecimalZeroMaxPrec,
		Qty:         utils.NewFromStringMaxPrec("0"),
		OrderType:   enum.OrderType_MO,
		Side:        enum.Side_Buy,
		OrderStatus: 0,
		UnfilledQty: utils.NewFromStringMaxPrec("0"),
		FilledQty:   utils.DecimalZeroMaxPrec,
	}
	buyOrder1.Amount = utils.NewFromStringMaxPrec("1000")
	buyOrder1.UnfilledAmount = utils.NewFromStringMaxPrec("1000")
	me.HandleOrder(sellOrder1)
	me.HandleOrder(sellOrder2)
	me.HandleOrder(buyOrder1)
	fmt.Println(me.asks)
	fmt.Println(me.bids)
	time.Sleep(time.Millisecond * 10)
	depth := me.GetDepth(200)
	for _, v := range depth.Asks {
		fmt.Printf("[asks] Price=%v Qty=%v\n", v.Price, v.Qty)
	}
	for _, v := range depth.Bids {
		fmt.Printf("[bids] Price=%v Qty=%v\n", v.Price, v.Qty)
	}
	time.Sleep(time.Second * 3)
}
func TestGetDepth(t *testing.T) {
	var c config.Config
	c.SymbolInfo = define.SymbolInfo{
		SymbolName:    "BTC_USDT",
		SymbolID:      1,
		BaseCoinName:  "BTC",
		BaseCoinID:    1,
		QuoteCoinName: "USDT",
		QuoteCoinID:   1,
		BaseCoinPrec:  3,
		QuoteCoinPrec: 4,
	}
	me := NewMatchEngine(&c, nil, nil)

	sellOrder1 := &Order{
		OrderID:     stringx.Randn(10),
		SequenceId:  time.Now().UnixNano() + rand.Int63n(time.Now().UnixNano()),
		CreateTime:  0,
		IsCancel:    false,
		Price:       utils.NewFromStringMaxPrec("12.12"),
		Qty:         utils.NewFromStringMaxPrec("90"),
		OrderType:   enum.OrderType_LO,
		Side:        enum.Side_Sell,
		OrderStatus: 0,
		UnfilledQty: utils.NewFromStringMaxPrec("90"),
		FilledQty:   utils.DecimalZeroMaxPrec,
	}
	sellOrder1.UnfilledAmount = sellOrder1.Qty.Mul(sellOrder1.Price)
	sellOrder1.Amount = sellOrder1.Qty.Mul(sellOrder1.Price)

	sellOrder2 := &Order{
		OrderID:     stringx.Randn(10),
		SequenceId:  time.Now().UnixNano() + rand.Int63n(time.Now().UnixNano()),
		CreateTime:  0,
		IsCancel:    false,
		Price:       utils.NewFromStringMaxPrec("12.12"),
		Qty:         utils.NewFromStringMaxPrec("5"),
		OrderType:   enum.OrderType_LO,
		Side:        enum.Side_Sell,
		OrderStatus: 0,
		UnfilledQty: utils.NewFromStringMaxPrec("5"),
		FilledQty:   utils.DecimalZeroMaxPrec,
	}
	sellOrder2.UnfilledAmount = sellOrder2.Qty.Mul(sellOrder2.Price)
	sellOrder2.Amount = sellOrder2.Qty.Mul(sellOrder1.Price)
	buyOrder1 := &Order{
		OrderID:     stringx.Randn(10),
		SequenceId:  time.Now().UnixNano(),
		CreateTime:  0,
		IsCancel:    false,
		Price:       utils.NewFromStringMaxPrec("5.12"),
		Qty:         utils.NewFromStringMaxPrec("12"),
		OrderType:   enum.OrderType_LO,
		Side:        enum.Side_Buy,
		OrderStatus: 0,
		UnfilledQty: utils.NewFromStringMaxPrec("10"),
		FilledQty:   utils.DecimalZeroMaxPrec,
	}
	buyOrder2 := &Order{
		OrderID:     stringx.Randn(10),
		SequenceId:  time.Now().UnixNano(),
		CreateTime:  0,
		IsCancel:    false,
		Price:       utils.NewFromStringMaxPrec("6.12"),
		Qty:         utils.NewFromStringMaxPrec("12"),
		OrderType:   enum.OrderType_LO,
		Side:        enum.Side_Buy,
		OrderStatus: 0,
		UnfilledQty: utils.NewFromStringMaxPrec("10"),
		FilledQty:   utils.DecimalZeroMaxPrec,
	}
	buyOrder2.Amount = buyOrder2.Qty.Mul(buyOrder2.Price)
	buyOrder2.UnfilledAmount = buyOrder2.Qty.Mul(buyOrder2.Price)
	me.HandleOrder(sellOrder1)
	me.HandleOrder(sellOrder2)
	me.HandleOrder(buyOrder1)
	me.HandleOrder(buyOrder2)
	fmt.Println(me.bids)
	fmt.Println(me.asks)
	depth := me.GetDepth(200)
	for _, v := range depth.Bids {
		fmt.Printf("bids Price=%v Qty=%v\n", v.Price, v.Qty)
	}
	for _, v := range depth.Asks {
		fmt.Printf("asks Price=%v Qty=%v\n", v.Price, v.Qty)
	}
	time.Sleep(time.Second * 3)
}
