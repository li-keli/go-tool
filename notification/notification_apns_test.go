package notification

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPushToIOS(t *testing.T) {
	Config = NotificationConfig{}

	Config.Ios.Enabled = true
	Config.Ios.Production = false
	Config.Ios.KeyPath = "../apns_debug.p12"
	Config.Ios.Password = "1"
	Config.Ios.KeyID = "com.jsj.AirRailwayMannagerForEmployee"

	err := NewAPNSClient()
	assert.Nil(t, err)

	var badge = 33
	req := PushNotification{
		Tokens:   []string{"059f2e6023baa8efdef3ad68bf1c37bb411c7102e10081edc50b02512a5d2f1c"},
		Platform: PlatformIos,
		Badge:    &badge,
		Alert: Alert{
			Title: "测试通道",
			Body:  "test01",
		},
		Topic: Config.Ios.KeyID,
		Data:  map[string]interface{}{},
	}

	isError := PushToIOS(req)
	assert.False(t, isError)
}

func TestApnsHostFromRequest(t *testing.T) {
	Config = NotificationConfig{}

	Config.Ios.Enabled = true
	Config.Ios.KeyPath = "../apns_debug.p12"

	err := NewAPNSClient()
	assert.Nil(t, err)
	//err = InitAppStatus()
	//assert.Nil(t, err)

	req := PushNotification{
		Production: true,
	}
	client := getApnsClient(req)
	assert.Equal(t, apns2.HostProduction, client.Host)

	req = PushNotification{
		Development: true,
	}
	client = getApnsClient(req)
	assert.Equal(t, apns2.HostDevelopment, client.Host)

	req = PushNotification{}
	Config.Ios.Production = true
	client = getApnsClient(req)
	assert.Equal(t, apns2.HostProduction, client.Host)

	Config.Ios.Production = false
	client = getApnsClient(req)
	assert.Equal(t, apns2.HostDevelopment, client.Host)
}

func Test_Apns_Push(t *testing.T) {
	var (
		client *apns2.Client
		err    error
	)
	cert, err := certificate.FromP12File("../apns_debug.p12", "1")
	if err != nil {
		logrus.Error("Cert Error:", err)
	}
	client = apns2.NewClient(cert).Development()
	client.Host = apns2.HostDevelopment // 测试地址

	notification := &apns2.Notification{}
	notification.DeviceToken = "01aa2364596a353dbaffdecd42f2c58ea63f49d6de2d80017b0c439ed0aa7c15"
	notification.Topic = "com.jsj.AirRailwayMannagerForEmployee"
	notification.Payload = []byte(`{"aps":{"alert":"Hello!"}}`)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := client.PushWithContext(ctx, notification)

	fmt.Println(res)
	assert.Equal(t, 200, res.StatusCode)
}
