package utils

import "github.com/spf13/cast"

var (
	ShardingCount = 10 //分表数量
)

func WithShardingSuffix(tableName string, userId int64) string {
	s := userId % int64(ShardingCount)
	suffix := cast.ToString(s)
	if s < 10 {
		suffix = "0" + suffix
	}
	return tableName + "_" + suffix
}
