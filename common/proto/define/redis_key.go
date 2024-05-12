package define

type RedisKey string

const (
	Ticker                   RedisKey = "gex:ticker"
	Kline                    RedisKey = "gex:kline"
	AccountToken             RedisKey = "gex:account:token"
	AccountConsumedMessageId RedisKey = "gex:account:consumed:messageId"
	OrderConsumedMessageId   RedisKey = "gex:order:consumed:messageId"
)

func (key RedisKey) WithSymbol(symbol string) string {
	return string(key) + "_" + symbol
}

func (key RedisKey) WithParams(params ...string) string {
	if len(params) == 0 {
		return string(key)
	}
	k := string(key)
	for _, v := range params {
		k += ":" + v
	}
	return k
}
