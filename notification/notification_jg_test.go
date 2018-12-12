// Android 推送模块
package notification

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushToAndroid(t *testing.T) {
	Config = NotificationConfig{}

	Config.Android.Enabled = true

	err := NewJgClient()
	assert.Nil(t, err)

	req := PushNotification{
		Tokens:   []string{"140fe1da9e858ebd37d"},
		Platform: PlatformIos,
		Alert: Alert{
			Title: "测试通道",
			Body:  "test01",
		},
		Data: map[string]interface{}{},
	}

	isError := PushToAndroid(req)
	assert.False(t, isError)
}
