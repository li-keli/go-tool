package notification

import (
	"github.com/sideshow/apns2"
	"github.com/ylywyn/jpush-api-go-client"
)

// 推送配置
type NotificationConfig struct {
	Android SectionAndroid
	Ios     SectionIos
}

type SectionAndroid struct {
	Enabled bool   `yaml:"enabled"`
	AppKey  string `yaml:"appkey"`
	Secret  string `yaml:"secret"`
}

type SectionIos struct {
	Enabled    bool   `yaml:"enabled" json:"enabled"`
	KeyPath    string `yaml:"key_path" json:"key_path"`
	KeyBase64  string `yaml:"key_base64" json:"key_base64"`
	KeyType    string `yaml:"key_type" json:"key_type"`
	Password   string `yaml:"password" json:"password"`
	Production bool   `yaml:"production" json:"production"`
	MaxRetry   int    `yaml:"max_retry" json:"max_retry"`
	KeyID      string `yaml:"key_id" json:"key_id"`
	TeamID     string `yaml:"team_id" json:"team_id"`
}

var (
	Config        NotificationConfig      // 配置文件 支持yaml、json 导入
	ApnsClient    *apns2.Client           // Apns(IOS)推送实例
	AndroidClient *jpushclient.PushClient // 极光(Android)推送实例
)
