package model

import (
	"Go-User-System/config"
	"database/sql"
	"log"
	"strconv"

	"github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"
)

var PG *sql.DB
var RD redis.Conn

// InitModel 连接数据库
func InitModel() {
	connStr := "host=" + config.Config.PG.Address + " port=" + strconv.Itoa(config.Config.PG.Port)
	connStr = connStr + " user=" + config.Config.PG.User + " password=" + config.Config.PG.Password
	connStr = connStr + " database=" + config.Config.PG.Database + " sslmode=disable"
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
