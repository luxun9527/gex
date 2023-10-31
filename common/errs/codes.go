package errs

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Code codes.Code

/*
	错误返回设计
	1、错误码能够体现出相对详细的错误，message暴露出更少的信息，详细的错误信息一般在日志中。
	2、定义在同一的地方，所有的服务去依赖。
	3、兼顾易用性，提供易用灵活的api。
*/

// 通用错误
const (
	Start Code = 100000
	// Internal 内部错误
	internal Code = 110000

	// ExecSqlFailed Sql执行失败
	execSqlFailed Code = 120000
	// RecordNotFound 在指定条件下查找有记录没找到。
	recordNotFound Code = 121000
	// RedisFailed 使用redis错误
	redisFailed Code = 130000
	// MongoFailed 使用mongo错误 。
	mongoFailed Code = 131000
	// KafkaFailed kafka错误 。
	kafkaFailed Code = 132000
	// EtcdFailed 使用Etcd错误 。
	etcdFailed Code = 133000
	// DTMFailed  使用dtm错误 。
	dTMFailed Code = 135000
	// PulsarFailed pulsar错误 。
	pulsarFailed Code = 136000

	// AuthFailed 认证失败
	authFailed Code = 140000
	// Timeout 超时
	timeout Code = 150000

	// ParamValidateFailed 请求参数校验失败
	paramValidateFailed Code = 171000
	// requestHeadValidateFailed 通用请求头校验失败
	requestHeadValidateFailed Code = 172000

	//======================================2开头为account的错误====================================

	//UserNotFound 用户不存在
	userNotFound Code = 200001
	// AmountInsufficient 用户余额不足
	amountInsufficient Code = 200002
	// TokenValidateFailed token验证失败
	tokenValidateFailed Code = 200003
	// TokenExpire Token到期
	tokenExpire Code = 200004
	// LoginFailed 登录账户密码验证失败
	loginFailed Code = 200005

	//============================3开头为order的错误==============================

	// OrderNotFound 订单为找到
	orderNotFound Code = 300001
	// OrderHasResolved 订单已经成交或已经取消
	orderHasResolved Code = 300002
	// LoOrderCancelFailed 市价单不允许手动取消
	loOrderCancelFailed Code = 300003
	// NotBids 订单簿没有买单
	notBids Code = 300004
	// NotAsks 订单簿没有卖单
	notAsks Code = 300005
)

func (c Code) Translate(lang string) string {
	return translator.translate(lang, c)
}

func (c Code) NewError() error {
	return status.New(codes.Code(c), c.String()).Err()
}

func (c Code) newError(msg string) error {
	return status.New(codes.Code(c), msg).Err()
}

func (c Code) String() string {
	return ""
}

func (c Code) dtmErrorMsg() string {
	return fmt.Sprintf("=%d=", int32(c))
}
