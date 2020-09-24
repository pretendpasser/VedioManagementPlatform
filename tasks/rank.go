package tasks

import "VMP/cac

// RestartDailyRank 重启一天的排名
func RestartDailyRank() error {
	return cache.RedisClient.Del("rank:daily").Err()
}