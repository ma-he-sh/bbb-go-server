package redis

import (
	"time"

	"github.com/go-redis/redis/v7"
	env "github.com/devmarka/bbb-go-server/env"
	"github.com/pkg/errors"
)

func GetClient() redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     env.REDISHost(),
		Password: env.REDISPass(),
		DB:       0,
	})
	return *client
}

func SetString(key string, value string, exptime time.Duration) string {
	var client = GetClient()
	err := client.Set(key, value, exptime*time.Second).Err()
	if err != nil {
		panic(err)
	}

	value, err = client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return value
}

func SetObject(key string, value interface{}, exptime time.Duration) interface{} {
	var client = GetClient()
	err := client.Set(key, value, exptime*time.Second).Err()
	if err != nil {
		panic(err)
	}

	value, err = client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return value
}

func GetString(key string) string {
	var client = GetClient()
	value, err := client.Get(key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		panic(err)
	}
	return value
}

func GetObject(key string) (string, error) {
	var client = GetClient()
	value, err := client.Get(key).Result()
	if err == redis.Nil {
		return value, errors.New("entry_not_found")
	} else if err != nil {
		return value, err
	}
	return value, err
}

func ResetKey(key string) {
	var client = GetClient()
	client.Do("EXPIRE", key)
}

func RemoveKey(key string) {
	var client = GetClient()
	client.Do("DELETE", key)
}