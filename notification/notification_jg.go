// Android 推送模块
package notification

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/ylywyn/jpush-api-go-client"
)

type JgClient struct {
	MasterSecret string
	AppKey       string
	AuthCode     string
	BaseUrl      string
}

// Jg Client
func NewJgClient() error {
	AndroidClient = jpushclient.NewPushClient(Config.Android.Secret, Config.Android.AppKey)

	return nil
}

func PushToAndroid(req PushNotification) bool {
	var pf jpushclient.Platform
	_ = pf.Add(jpushclient.ANDROID)

	var ad jpushclient.Audience
	ad.SetID(req.Tokens)

	var notice jpushclient.Notice
	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: req.Title})

	var msg jpushclient.Message
	msg.Title = req.Alert.Title
	msg.Content = req.Alert.Body
	msg.Extras = req.Data

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetMessage(&msg)
	payload.SetNotice(&notice)

	bytes, _ := payload.ToBytes()

	fmt.Println(string(bytes))

	str, err := AndroidClient.Send(bytes)

	logrus.Info(str)

	return err != nil
}
