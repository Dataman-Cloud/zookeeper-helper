package config

import (
	"log"
	"runtime"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

const (
	DefaultTimeout   = 24 * 3600
	DefaultHost      = ""
	DefaultPort      = 6500
	DefaultDebugging = true
	DefaultLogLevel  = "debug"
)

type Config struct {
	NumCPU    int
	Host      string
	Port      uint
	Debugging bool
	Log       *LogConfig
}

type LogConfig struct {
	Console   bool
	File      string
	FileNum   int
	FileSize  int
	Level     string
	Formatter string
}

var pairs Config

func Pairs() Config {
	return pairs
}

func init() {
	jww.SetLogThreshold(jww.LevelTrace)
	jww.SetStdoutThreshold(jww.LevelInfo)

	log.Println("initing config ...")

	viper.SetConfigName("zookeeper-helper")
	viper.AddConfigPath("./")
	viper.AddConfigPath("$HOME/.omega/")
	viper.AddConfigPath("/etc/omega/")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Panicln("can't read config file:", err)
	}

	initDefault()

	if err := viper.Unmarshal(&pairs); err != nil {
		log.Panicln("can't covert to config pairs: ", err)
	}

	if !pairs.Debugging {
		jww.SetLogThreshold(jww.LevelError)
		jww.SetStdoutThreshold(jww.LevelError)
	}

	log.Printf("initialized config pairs: %+v\n", pairs)
}

func initDefault() {
	viper.SetDefault("numCPU", runtime.NumCPU())
	viper.SetDefault("host", DefaultHost)
	viper.SetDefault("port", DefaultPort)
	viper.SetDefault("debugging", true)
}

func Get(name string) interface{} {
	return viper.Get(name)
}

func GetString(name string) string {
	return viper.GetString(name)
}

func GetInt(name string) int {
	return viper.GetInt(name)
}
