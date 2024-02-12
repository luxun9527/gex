package errs

import (
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Code codes.Code

const (
	CommonCodeInit Code = 100000 * (iota + 1)
	AccountCodeInit
	AdminCodeInit
	MatchCodeInit
	OrderCodeInit
	QuoteCodeInit
)

func WarpMessage(err error, msg string) error {
	s, ok := status.FromError(err)
	if ok {
		msg = s.Message() + ":" + msg
		return status.New(s.Code(), msg).Err()
	}
	return errors.Wrap(err, msg)
}

// CastToDtmError dtm规定 saga失败如果要补偿的话grpc返回abort 其他错误则为重试 我们的错误只能在msg中体现 refer https://dtm.pub/practice/workflow.html#%E5%88%86%E6%94%AF%E6%93%8D%E4%BD%9C%E7%BB%93%E6%9E%9C
func CastToDtmError(err error) error {
	s, ok := status.FromError(err)
	if ok {
		return status.Error(codes.Aborted, Code(s.Code()).DtmErrorMsg())
	}
	return status.Error(codes.Aborted, InternalCode.DtmErrorMsg())
}

func (c Code) Translate(lang string) string {
	return translator.translate(lang, c)
}

func (c Code) Error(msg string) error {
	return status.New(codes.Code(c), msg).Err()
}

func (c Code) String() string {
	return ""
}

func (c Code) DtmErrorMsg() string {
	return fmt.Sprintf("=%d=", int32(c))
}
