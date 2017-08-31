package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type (
	configStruct struct {
		Port        string            `json:"port"`
		DBConfig    map[string]string `json:"postgre"`
		RedisConfig map[string]string `json:"redis"`
		MongoConfig map[string]string `json:"mongo"`
	}
)

var (
	// Config return variable
	Config configStruct
	// Env return variable
	Env string
)

func init() {
	Config, _ = initConfig()
}

func initConfig() (cfg configStruct, err error) {
	if os.Getenv("APPENV") == "" {
		Env = "development"
	}
	cfgPath := fmt.Sprintf("/etc/gonyast/config.%s.json", Env)
	if _, err := os.Stat(cfgPath); err != nil {
		if os.IsNotExist(err) {
			cfgPath = fmt.Sprintf(".%s", cfgPath)
		}
	}
	file, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		log.Fatal("Config file missing :", err)
	}
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatal("Config parese failed :", err)
	}
	return
}
