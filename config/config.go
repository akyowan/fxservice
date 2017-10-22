package config

import (
	"aposervice/domain"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Mysql              domain.DBConf
	Redis              domain.DBConf
	MongoDB            domain.DBConf
	Server             domain.ServerConf
	CommentDayMaxCount int64
}

var (
	Conf     Config
	LogLevel string
	IsDebug  bool
)

var (
	configFilePath = flag.String("c", "/etc/fxservice/apocenter/config.json", "config file path")
	isDebug        = flag.Bool("d", false, "debug mode")
	isLocal        = flag.Bool("local", false, "local develop mode")
)

func init() {
	flag.Parse()
	if isDebug != nil {
		IsDebug = *isDebug
	}
	file, err := os.Open(*configFilePath)
	if err != nil {
		fmt.Println("starting with default config...")
	} else {
		fmt.Println("reading config from:", *configFilePath)
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Conf)
		if err != nil {
			panic(err)
		}
	}
}
