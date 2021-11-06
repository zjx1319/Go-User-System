package model

import (
	"Go-User-System/config"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"
)

var PG *sql.DB
var RD redis.Conn

// InitModel 连接数据库
func InitModel() {
	pgConfig := make(map[string]string)
	pgConfig["host"] = config.Config.PG.Address
	pgConfig["port"] = strconv.Itoa(config.Config.PG.Port)
	pgConfig["user"] = config.Config.PG.User
	pgConfig["password"] = config.Config.PG.Password
	pgConfig["database"] = config.Config.PG.Database
	connStr := ""
	for key, value := range pgConfig {
		connStr += fmt.Sprintf("%s=%s ", key, value)
	}
	connStr += " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}
	PG = db

	setDB := redis.DialDatabase(config.Config.RD.Database)
	setPasswd := redis.DialPassword(config.Config.RD.Password)
	RD, err = redis.Dial("tcp", config.Config.RD.Address+":"+strconv.Itoa(config.Config.RD.Port), setDB, setPasswd)
	if err != nil {
		log.Panic(err)
	}

	initModelUser()
}
