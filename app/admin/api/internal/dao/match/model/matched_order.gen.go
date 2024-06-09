// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMatchedOrder = "matched_order"

// MatchedOrder mapped from table <matched_order>
type MatchedOrder struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:雪花算法id" json:"id"`                        // 雪花算法id
	MatchID      string `gorm:"column:match_id;not null;comment:撮合id" json:"match_id"`                                   // 撮合id
	MatchSubID   string `gorm:"column:match_sub_id;not null;comment:本次匹配的id，一次撮合会多次匹配" json:"match_sub_id"`              // 本次匹配的id，一次撮合会多次匹配
	SymbolID     int32  `gorm:"column:symbol_id;not null;comment:交易对id" json:"symbol_id"`                                // 交易对id
	SymbolName   string `gorm:"column:symbol_name;not null;comment:交易对名称" json:"symbol_name"`                            // 交易对名称
	TakerUserID  int64  `gorm:"column:taker_user_id;not null;comment:taker用户id" json:"taker_user_id"`                    // taker用户id
	TakerOrderID string `gorm:"column:taker_order_id;not null;comment:taker订单id" json:"taker_order_id"`                  // taker订单id
	MakerOrderID string `gorm:"column:maker_order_id;not null;comment:maker订单id" json:"maker_order_id"`                  // maker订单id
	MakerUserID  int64  `gorm:"column:maker_user_id;not null;comment:maker用户id" json:"maker_user_id"`                    // maker用户id
	TakerIsBuyer int32  `gorm:"column:taker_is_buyer;not null;default:2;comment:taker是否是买单 1是 2否" json:"taker_is_buyer"` // taker是否是买单 1是 2否
	Price        string `gorm:"column:price;not null;comment:价格" json:"price"`                                           // 价格
	Qty          string `gorm:"column:qty;not null;comment:数量(基础币)" json:"qty"`                                          // 数量(基础币)
	Amount       string `gorm:"column:amount;not null;comment:金额（计价币）" json:"amount"`                                    // 金额（计价币）
	MatchTime    int64  `gorm:"column:match_time;not null;comment:撮合时间" json:"match_time"`                               // 撮合时间
	CreatedAt    int64  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`                               // 创建时间
	UpdatedAt    int64  `gorm:"column:updated_at;not null;comment:修改时间" json:"updated_at"`                               // 修改时间
}

// TableName MatchedOrder's table name
func (*MatchedOrder) TableName() string {
	return TableNameMatchedOrder
}