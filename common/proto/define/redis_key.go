package define

type RedisKey string

const (
	Ticker RedisKey = "ticker"
)

func (key RedisKey) WithSymbol(symbol string) string {
	return string(key) + "_" + symbol
}
