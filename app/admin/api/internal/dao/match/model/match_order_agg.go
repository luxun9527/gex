package model

import (
	"database/sql/driver"
	"encoding/json"
)

type MatchedOrderAgg struct {
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:雪花算法id" `                                 // 雪花算法id
	MatchID      string         `gorm:"column:match_id;not null;comment:撮合id" `                                                  // 撮合id
	SymbolID     int32          `gorm:"column:symbol_id;not null;comment:交易对id"`                                                 // 交易对id
	SymbolName   string         `gorm:"column:symbol_name;not null;comment:交易对名称" `                                              // 交易对名称
	TotalAmount  string         `gorm:"column:total_amount;not null;comment:总金额" `                                               // 交易对名称
	TotalQty     string         `gorm:"column:total_qty;not null;comment:总数量" `                                                  // 交易对名称
	TakerIsBuyer int32          `gorm:"column:taker_is_buyer;not null;default:2;comment:taker是否是买单 1是 2否" json:"taker_is_buyer"` // taker是否是买单 1是 2否
	CreatedAt    int64          `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`                               // 创建时间
	SubMatchList SubMatchedList `gorm:"column:sub_match_list;not null;comment:子撮合" `                                             // 交易对名称
}

type SubMatchedOrder struct {
	TakerUserID  int64  `gorm:"column:taker_user_id;not null;comment:taker用户id" json:"taker_user_id"`   // taker用户id
	TakerOrderID string `gorm:"column:taker_order_id;not null;comment:taker订单id" json:"taker_order_id"` // taker订单id
	MakerOrderID string `gorm:"column:maker_order_id;not null;comment:maker订单id" json:"maker_order_id"` // maker订单id
	MakerUserID  int64  `gorm:"column:maker_user_id;not null;comment:maker用户id" json:"maker_user_id"`   // maker用户id
	Price        string `gorm:"column:price;not null;comment:价格" json:"price"`                          // 价格
	Qty          string `gorm:"column:qty;not null;comment:数量(基础币)" json:"qty"`                         // 数量(基础币)
	Amount       string `gorm:"column:amount;not null;comment:金额（计价币）" json:"amount"`                   // 金额（计价币）
	MatchTime    int64  `gorm:"column:match_time;not null;" json:"match_time"`                          // 金额（计价币）
}

type SubMatchedList []*SubMatchedOrder

func (data *SubMatchedList) Scan(value interface{}) (err error) {
	return json.Unmarshal(value.([]byte), data)
}

func (data *SubMatchedList) Value() (driver.Value, error) {
	return json.Marshal(data)
}
