package define

import "sync"

type SymbolInfo struct {
	SymbolName    string
	SymbolID      int32
	BaseCoinName  string
	BaseCoinID    int32
	QuoteCoinName string
	QuoteCoinID   int32
	BaseCoinPrec  int32
	QuoteCoinPrec int32
}

type CoinInfo struct {
	CoinID   int32
	CoinName string
	Prec     int32
}

type SymbolCoinConfig[Key string | int32, Value *SymbolInfo | *CoinInfo] map[Key]Value

func (s SymbolCoinConfig[Key, Value]) CastToSyncMap() *sync.Map {
	m := &sync.Map{}
	for k, v := range s {
		m.Store(k, v)
	}
	return m
}
