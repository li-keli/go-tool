// 配合gin读取config配置
// 支持多环境配置

package config

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := MyConfig{}

	NewConfig(&config)

	fmt.Printf("service name is :%s\n", config.Service.Name)
	fmt.Printf("service port is :%d\n", config.Service.Port)
}

type MyConfig struct {
	Service struct {
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
	} `yaml:"service"`
}
