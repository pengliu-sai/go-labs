package labs01

import (
	"fmt"
	"math/rand"
	"testing"

	"gopkg.in/redis.v5"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	initBaseData()
}

func initBaseData() {
	times := 1000000 // 50W
	for times > 0 {
		member := fmt.Sprintf("KENAN_%d", times)
		zadd(member)
		times--
	}
}

func zadd(member string) {
	cmd := client.ZAdd("test_rankings", redis.Z{
		Score:  float64(rand.Intn(2000-1200) + 1200),
		Member: member,
	})
	_, err := cmd.Result()
	if err != nil {
		panic(err.Error())
	}
}

func zrevrange() {
	cmd := client.ZRevRangeWithScores("test_rankings", 0, 100)
	rs, err := cmd.Result()
	if err != nil {
		panic(err.Error())
	}
	_ = rs
}

func zscore(member string) {
	cmd := client.ZScore("test_rankings", member)
	rs, err := cmd.Result()
	_ = rs
	_ = err
}

func Benchmark_ZADD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		member := fmt.Sprintf("KENAN2_%d", i)
		zadd(member)
	}
}

func Benchmark_ZREVRANGE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zrevrange()
	}
}

func Benchmark_ZSCORE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		member := fmt.Sprintf("KENAN_%d", i)
		zscore(member)
	}
}

func Benchmark_ZADD_REPEAT(b *testing.B) {
	times := 1000000 // 50W
	for times > 0 {
		member := fmt.Sprintf("KENAN_%d", times)
		zadd(member)
		times--
	}
}
