package model

import (
	"Go-User-System/config"
	"github.com/garyburd/redigo/redis"
)

func BanToken(token string) (err error) {
	_, err = RD.Do("SET", token, true, "EX", config.Config.JWT.Expire*60)
	return
}

func IsTokenBanned(token string) (is bool, err error) {
	is, err = redis.Bool(RD.Do("GET", token))
	if err == redis.ErrNil {
		err = nil
	}
	return
}
