package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jinzhu/configor"
)

var Config = struct {
	Url      string `env:"Url" default:"http://locahost"`
	Name     string `env:"Name" default:"local"`
	Host     string `env:"Host" default:"localhost"`
	Port     string `env:"Port" default:"3306"`
	User     string `env:"User" default:"root"`
	Password string `env:"Password" default:"123456"`
}{}

func init() {
	configPath := filepath.Join(CWD(), "application.yml")
	fmt.Println(configPath)
	if err := configor.Load(&Config, configPath); err != nil {
		fmt.Println(fmt.Sprintf("Config Path:%s ,Error:%s", configPath, err.Error()))
	}

}

func CWD() string {
	// 兼容 travis 集成测试
	if os.Getenv("TRAVIS_BUILD_DIR") != "" {
		return os.Getenv("TRAVIS_BUILD_DIR")
	}

	path, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(path)
}
