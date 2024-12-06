package conf

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kr/pretty"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"sync"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/04 下午11:36
 * @FilePath: conf
 * @Description: 实现配置加载
 */

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env      string
	Kitex    Kitex    `yaml:"kitex"`
	MySQL    MySQL    `yaml:"mysql"`
	Redis    Redis    `yaml:"redis"`
	Registry Registry `yaml:"registry"`
	Email    Email    `yaml:"email"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Kitex struct {
	Service         string `yaml:"service"`
	Address         string `yaml:"address"`
	EnablePprof     bool   `yaml:"enable_pprof"`
	EnableGzip      bool   `yaml:"enable_gzip"`
	EnableAccessLog bool   `yaml:"enable_access_log"`
	LogLevel        string `yaml:"log_level"`
	LogFileName     string `yaml:"log_file_name"`
	LogMaxSize      int    `yaml:"log_max_size"`
	LogMaxBackup    int    `yaml:"log_max_backup"`
	LogMaxAge       int    `yaml:"log_max_age"`
}

type Registry struct {
	RegistryAddress string `yaml:"registry_address"`
}

type Email struct {
	Host       string `yaml:"host"`
	TLSAddr    string `yaml:"tls_addr"`
	ServerName string `yaml:"server_name"`
	Password   string `yaml:"password"`
}

// GetConf 获取配置实例
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	prefix := "micro_services/user_server/conf"
	configFileRelPath := filepath.Join(prefix, filepath.Join(GetEnv(), "conf.yaml"))
	content, err := os.ReadFile(configFileRelPath)
	if err != nil {
		panic(err)
	}
	conf = &Config{}
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		klog.Errorf("parse yaml error: %v", err)
		panic(err)
	}
	if err := validator.Validate(conf); err != nil {
		klog.Errorf("validate config error: %v", err)
		panic(err)
	}
	conf.Env = GetEnv()
	_, _ = pretty.Printf("%+v\n", conf)
}

func GetEnv() string {
	e := os.Getenv("GO_ENV")
	if len(e) <= 0 {
		return "dev"
	}
	return e
}

func LogLevel() klog.Level {
	level := GetConf().Kitex.LogLevel
	switch level {
	case "trace":
		return klog.LevelTrace
	case "debug":
		return klog.LevelDebug
	case "info":
		return klog.LevelInfo
	case "notice":
		return klog.LevelNotice
	case "warn":
		return klog.LevelWarn
	case "error":
		return klog.LevelError
	case "fatal":
		return klog.LevelFatal
	default:
		return klog.LevelInfo
	}
}
