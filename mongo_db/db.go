package mongo_db

import (
	"github.com/li-keli/mgo"
	"github.com/sirupsen/logrus"
)

var mongoDbSession *mgo.Session

// mongodb conn init
func NewMongo(url string) {
	session, err := mgo.Dial(url)
	if err != nil {
		logrus.Fatal("mongodb connection error: ", err, url)
	}
	session.SetMode(mgo.Monotonic, true)
	mongoDbSession = session
}

// 获取mongodb会话
// 请注意获取到的session，一定要defer session.close()
func GetMongoSession() *mgo.Session {
	if mongoDbSession == nil {
		panic("mongodb链接未初始化")
	}
	return mongoDbSession.Copy()
}
