package common

import (
	"io/ioutil"
	"os"
	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		Server   ServerConfig
		Database DatabaseConfig
		Redis    RedisConfig
		Etcd     EtcdConfig
	}

	ServerConfig struct {
		Host       string
		Port       int
		Env        string
		Cron       bool
		GeneralLog string `toml:"general_log"`
		ErrorLog   string `toml:"error_log"`
		AccessLog  string `toml:"access_log"`
	}

	DatabaseConfig struct {
		Host     string
		Port     int
		User     string
		Password string
		Db       string
		MaxOpen  int
		MaxIdle  int
	}

	RedisConfig struct {
		Host     string
		Port     int
		Password string
	}

	EtcdConfig struct {
		Host string
		Port int
	}
)

var configInfo Config

func ParseConfigFile(filePath *string) {
	fi, err := os.Open(*filePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if _, err := toml.Decode(string(fd), &configInfo); err != nil {
		panic(err)
	}
}

func GetConfig() Config {
	return configInfo
}
