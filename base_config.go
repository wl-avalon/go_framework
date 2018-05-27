//基础配置文件
package go_framework

import (
	"gopkg.in/ini.v1"
	"go_framework/request/rpc"
)

type Config struct {
	AppConfig AppConfig
	LogConfig LogConfig
}

//APP相关配置
type AppConfig struct {
	AppName string `ini:"appName"`
	RunMode string `ini:"runMode"`
}

//日志配置
type LogConfig struct {
	LogFileDir string `ini:"logFileDir"`
	LogLevel   string `ini:"logLevel"`
}

var BaseConfig *Config

func InitConfig(path string) error {
	config 		:= new(Config)
	conf, err := ini.Load(path) //加载配置文件
	if err != nil {
		return err
	}

	err = conf.Section("log").MapTo(&(config.LogConfig))
	if err != nil {
		return err
	}

	err = conf.Section("app").MapTo(&(config.AppConfig))
	if err != nil {
		return err
	}

	domain := conf.Section("domain").KeysHash()
	rpc.InitRpc(domain)
	BaseConfig = config
	return nil
}
