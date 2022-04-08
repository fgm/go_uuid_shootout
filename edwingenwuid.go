package uuid_shootout

import (
	"context"
	"strconv"

	"github.com/edwingeng/wuid/redis/v8/wuid"
	"github.com/go-redis/redis/v8"
)

var g = wuid.NewWUID("default", nil)

func EdwingenWuidGen1() string {
	return strconv.Itoa(int(g.Next()))
}

func EdwingenWuidNewClient() (client redis.UniversalClient, autoDisconnect bool, err error) {
	client = redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	err = client.Ping(context.Background()).Err()
	return client, false, err
}

func init() {
	if err := g.LoadH28FromRedis(EdwingenWuidNewClient, "wuid"); err != nil {
		panic(err)
	}
}
