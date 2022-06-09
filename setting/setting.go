package setting

// 解析ini文件的工具包
import (
	"gopkg.in/ini.v1"
)


//通过函数new()为Type分配内存
var Conf = new(AppConfig)

// 总配置结构体
type AppConfig struct {
	Release bool `ini:"release"`
	Port    int  `ini:"port"`
	*MySQLConfig `ini:"mysql"`
	*RedisConfig `ini:"redis"`
	*TokenConfig `ini:"token"`
}

// mysql 配置结构体
type MySQLConfig struct {
	User string `ini:"user"`
	Password string `ini:"password"`
	DB string `ini:"db"`
	Host string `ini:"host"`
	Port int `ini:"port"`
}

// redis 配置
type RedisConfig struct {
	Host string `ini:"host"`
	Port int `ini:"port"`
	DB   int `ini:"db"`
}

type TokenConfig struct {
	SigningKey string `ini:"singing_key"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
