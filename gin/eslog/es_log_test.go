package eslog

import (
	"testing"
	"time"
)

func TestNewGinLog(t *testing.T) {
	t.Run("es_log测试", func(t *testing.T) {
		log := NewGinLog("172.16.7.20", "http://172.16.7.20:9200", "mylog")
		log.WithField("Key", "fps").Debug("调试错误")
		log.WithField("Key", "fps").Info("简单异常")
		log.WithField("Key", "fps").Error("出错误了")
	})

	time.Sleep(time.Second * 1)
}
