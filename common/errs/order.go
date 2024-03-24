//go:generate errgen -p order.go

package errs

const (

	// OrderNotFoundCode 订单为未找到
	OrderNotFoundCode Code = OrderCodeInit + iota + 1
	// OrderHasResolvedCode 订单已经成交或已经取消
	OrderHasResolvedCode
	// LoOrderCancelFailedCode 市价单不允许手动取消
	LoOrderCancelFailedCode
	// NotBidsCode 订单簿没有买单
	NotBidsCode
	// NotAsksCode 订单簿没有卖单
	NotAsksCode

	ErrPrecCode
)

var (
	OrderNotFound       = OrderNotFoundCode.Error("")
	OrderHasResolved    = OrderHasResolvedCode.Error("")
	LoOrderCancelFailed = LoOrderCancelFailedCode.Error("")
	NotBids             = NotBidsCode.Error("")
	NotAsks             = NotAsksCode.Error("")
	ErrPrec             = ErrPrecCode.Error("")
)
