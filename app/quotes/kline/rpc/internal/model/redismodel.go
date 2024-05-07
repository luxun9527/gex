package model

import "github.com/luxun9527/gex/app/quotes/kline/rpc/internal/dao/model"

type RedisModel struct {
	model.Kline
	MatchID int64 `json:"match_id"`
}
