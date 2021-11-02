package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type TypeAppConfig struct {
	Address string `json:"address"`
}

type TypePGConfig struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type TypeSMTPConfig struct {
	Server    string `json:"server"`
	Port      int    `json:"port"`
	Address   string `json:"address"`
	Password  string `json:"password"`
	VerifyURL string `json:"verify_url"`
}

type TypeJWTConfig struct {
	Secret string `json:"secret"`
	Expire int    `json:"expire"`
}

type TypeWXConfig struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type TypeConfig struct {
	App  TypeAppConfig  `json:"app"`
	PG   TypePGConfig   `json:"pg"`
	SMTP TypeSMTPConfig `json:"smtp"`
	JWT  TypeJWTConfig  `json:"jwt"`
	WX   TypeWXConfig   `json:"wx"`
}

var Config TypeConfig

// InitConfig 读入配置
func InitConfig() {
	configFilename := "config.json"

	configFile, err := ioutil.ReadFile("./config/" + configFilename)

	if err != nil {
		log.Println("config: read file error " + configFilename)
		log.Panic(err)
	}

	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		log.Println("config: json unmarshal error " + configFilename)
		log.Panic(err)
	}

	log.Println("config: config " + configFilename + " loaded")
}
