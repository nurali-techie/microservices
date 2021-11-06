package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var (
	REDIS_URL = "search_redis:6379"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(
		&redis.Options{
			Addr: REDIS_URL,
		},
	)
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
	return client
}

type Cache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) *Cache {
	return &Cache{
		client: client,
	}
}

func (c *Cache) GetRestaurant(restoID string) (*Restaurant, error) {
	restoVal, err := c.client.Get(restoID).Result()
	if err != nil {
		return nil, err
	}

	var resto *Restaurant
	if err := json.NewDecoder(bytes.NewReader([]byte(restoVal))).Decode(&resto); err != nil {
		return nil, err
	}
	return resto, nil
}

func (c *Cache) SetRestaurant(resto *Restaurant) error {
	restoVal, err := json.Marshal(resto)
	if err != nil {
		return nil
	}
	expiry := endOfDay()
	fmt.Println(expiry)
	return c.client.Set(resto.ID, restoVal, endOfDay()).Err()
}

func endOfDay() time.Duration {
	t := time.Now()
	return time.Duration(time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location()).UnixNano())
}
