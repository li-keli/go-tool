// 配合gin读取config配置
// 支持多环境配置

package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewConfig(conf interface{}) {

	viper.AddConfigPath(".")

	viper.SetDefault("PRODUCTION", false)
	if viper.GetBool("PRODUCTION") {
		viper.SetConfigName("config-prod")
	} else {
		viper.SetConfigName("config-dev")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := viper.Unmarshal(conf); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
