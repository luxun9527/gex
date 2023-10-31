package errs

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	Internal                  = internal.newError("")
	ExecSqlFailed             = execSqlFailed.newError("")
	RecordNotFound            = recordNotFound.newError("")
	RedisFailed               = redisFailed.newError("")
	MongoFailed               = mongoFailed.newError("")
	KafkaFailed               = kafkaFailed.newError("")
	EtcdFailed                = etcdFailed.newError("")
	DTMFailed                 = dTMFailed.newError("")
	PulsarFailed              = pulsarFailed.newError("")
	AuthFailed                = authFailed.newError("")
	Timeout                   = timeout.newError("")
	ParamValidateFailed       = paramValidateFailed.newError("")
	RequestHeadValidateFailed = requestHeadValidateFailed.newError("")
	UserNotFound              = userNotFound.newError("")
	AmountInsufficient        = amountInsufficient.newError("")
	TokenValidateFailed       = tokenValidateFailed.newError("")
	TokenExpire               = tokenExpire.newError("")
	LoginFailed               = loginFailed.newError("")
	OrderNotFound             = orderNotFound.newError("")
	OrderHasResolved          = orderHasResolved.newError("")
	LoOrderCancelFailed       = loOrderCancelFailed.newError("")
	NotBids                   = notBids.newError("")
	NotAsks                   = notAsks.newError("")
)

func WarpMessage(err error, msg string) error {
	s, ok := status.FromError(err)
	if ok {
		msg = s.Message() + ":" + msg
		return status.New(s.Code(), msg).Err()
	}
	return err
}

// CastToDtmError dtm规定 saga失败如果要补偿的话grpc返回abort 其他错误则为重试 我们的错误只能在msg中体现 refer https://dtm.pub/practice/workflow.html#%E5%88%86%E6%94%AF%E6%93%8D%E4%BD%9C%E7%BB%93%E6%9E%9C
func CastToDtmError(err error) error {
	s, ok := status.FromError(err)
	if ok {
		return status.Error(codes.Aborted, Code(s.Code()).dtmErrorMsg())
	}
	return status.Error(codes.Aborted, internal.dtmErrorMsg())
}
