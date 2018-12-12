// 推送模块
// 目前仅整合APNS（IOS）、极光推送（Android）
package notification

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type D map[string]interface{}

const (
	ApnsPriorityLow  = 5
	ApnsPriorityHigh = 10
)

type PlatformType string

var (
	PlatformAndroid PlatformType = "android"
	PlatformIos     PlatformType = "ios"
)

type Alert struct {
	Action       string   `json:"action,omitempty"`
	ActionLocKey string   `json:"action-loc-key,omitempty"`
	Body         string   `json:"body,omitempty"`
	LaunchImage  string   `json:"launch-image,omitempty"`
	LocArgs      []string `json:"loc-args,omitempty"`
	LocKey       string   `json:"loc-key,omitempty"`
	Title        string   `json:"title,omitempty"`
	Subtitle     string   `json:"subtitle,omitempty"`
	TitleLocArgs []string `json:"title-loc-args,omitempty"`
	TitleLocKey  string   `json:"title-loc-key,omitempty"`
}

type RequestPush struct {
	Notifications []PushNotification `json:"notifications" binding:"required"`
}
type PushNotification struct {
	// Common
	Tokens           []string     `json:"tokens" binding:"required"`
	Platform         PlatformType `json:"platform" binding:"required"`
	Message          string       `json:"message,omitempty"`
	Title            string       `json:"title,omitempty"`
	Priority         string       `json:"priority,omitempty"`
	ContentAvailable bool         `json:"content_available,omitempty"`
	MutableContent   bool         `json:"mutable_content,omitempty"`
	Sound            interface{}  `json:"sound,omitempty"`
	Data             D            `json:"data,omitempty"`
	Retry            int          `json:"retry,omitempty"`
	wg               *sync.WaitGroup

	// iOS
	Expiration  int64    `json:"expiration,omitempty"`
	ApnsID      string   `json:"apns_id,omitempty"`
	CollapseID  string   `json:"collapse_id,omitempty"`
	Topic       string   `json:"topic,omitempty"`
	Badge       *int     `json:"badge,omitempty"`
	Category    string   `json:"category,omitempty"`
	ThreadID    string   `json:"thread-id,omitempty"`
	URLArgs     []string `json:"url-args,omitempty"`
	Alert       Alert    `json:"alert,omitempty"`
	Production  bool     `json:"production,omitempty"`
	Development bool     `json:"development,omitempty"`
	SoundName   string   `json:"name,omitempty"`
	SoundVolume float32  `json:"volume,omitempty"`

	// Android
	// 公共部分足以
}

func (p *PushNotification) WaitDone() {
	if p.wg != nil {
		p.wg.Done()
	}
}

func (p *PushNotification) AddWaitCount() {
	if p.wg != nil {
		p.wg.Add(1)
	}
}

// push messages to app
func PushToApp(req PushNotification) bool {
	switch req.Platform {
	case PlatformIos:
		if err := NewAPNSClient(); err != nil {
			logrus.Error(err)
			return false
		}

		return PushToIOS(req)
	case PlatformAndroid:
		if err := NewJgClient(); err != nil {
			logrus.Error(err)
			return false
		}

		return PushToAndroid(req)
	default:
		logrus.Error("未知的推送终端")
		return false
	}
}
