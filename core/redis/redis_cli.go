package redis

import (
	"../../settings"
	"github.com/garyburd/redigo/redis"
)

type RedisCli struct {
	conn redis.Conn
}

type RedisPool struct {
	pool *redis.Pool
}

var instanceRedisCli *RedisCli = nil
var instanceRedisPool *RedisPool = nil
var ErrNil = redis.ErrNil

func Pconnect(db int) (conn *RedisCli) {
	instanceRedisCli = new(RedisCli)
	instanceRedisCli.conn = instanceRedisPool.pool.Get()
	_, err := instanceRedisCli.conn.Do("SELECT", db)
	if err != nil {
		panic(err)
	}

	return instanceRedisCli
}

func Pinit() {
	if instanceRedisPool == nil {
		instanceRedisPool = new(RedisPool)
		host := settings.Get().RedisHost
		instanceRedisPool.pool = redis.NewPool(func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host+":6379")

			if err != nil {
				return nil, err
			}

			return c, err
		}, 30)
	}
}

func Pclose() {
	if instanceRedisPool != nil {
		instanceRedisPool.pool.Close()
	}
}

func (redisCli *RedisCli) Close() {
	redisCli.conn.Close()
}

func (redisCli *RedisCli) SetValue(key string, value string, expiration ...interface{}) error {
	_, err := redisCli.conn.Do("SET", key, value)

	if err == nil && expiration != nil {
		redisCli.conn.Do("EXPIRE", key, expiration[0])
	}

	return err
}

func (redisCli *RedisCli) HappendValue(hash string, key string, value string) error {
	now, err := redis.String(redisCli.conn.Do("HGET", hash, key))
	if err != nil {
		now = ""
	}

	_, serr := redisCli.conn.Do("HSET", hash, key, now+value)

	return serr
}

func (redisCli *RedisCli) HsetValue(hash string, key string, value string) error {
	_, err := redisCli.conn.Do("HSET", hash, key, value)

	return err
}

func (redisCli *RedisCli) HgetValue(hash string, key string) (string, error) {
	return redis.String(redisCli.conn.Do("HGET", hash, key))
}

func (redisCli *RedisCli) GetValue(key string) (string, error) {
	return redis.String(redisCli.conn.Do("GET", key))
}

func (redisCli *RedisCli) HgetAll(hash string) (map[string]string, error) {
	return redis.StringMap(redisCli.conn.Do("HGETALL", hash))
}
