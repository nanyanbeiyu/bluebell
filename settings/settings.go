package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// 全局变量，存储所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	Port         int    `mapstructure:"port"`
	*MysqlConfig `mapstructure:"mysql"`
	*LogConfig   `mapstructure:"log"`
	*RedisConfig `mapstructure:"redis"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db_name"`
	MaxOpenConns int    `mapstructure:"max_open_conn"`
	MaxIdleConns int    `mapstructure:"max_idle_conn"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Post     int    `mapstructure:"Post"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// 使用viper读取配置

func Init() (err error) {
	viper.SetConfigFile("./config/config.yaml") // 指定配置文件路径
	err = viper.ReadInConfig()                  // 读取配置信息
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
		return // 读取配置信息失败
	}
	// 将读取的配置信息保存至全局变量Conf
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
	}
	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改")
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("Fatal error config file: %s \n", err)
		}
	})

	return
}
