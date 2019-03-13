package wechat

import (
	"github.com/li-keli/go-tool/wechat/cache"
	"github.com/li-keli/go-tool/wechat/context"
	"github.com/li-keli/go-tool/wechat/kf"
	"github.com/li-keli/go-tool/wechat/material"
	"github.com/li-keli/go-tool/wechat/server"
	"github.com/li-keli/go-tool/wechat/user"

	"net/http"
	"sync"
)

// 微信上下文模型
type Wechat struct {
	Context *context.Context
}

// 配置
type Config struct {
	AppID               string
	AppSecret           string
	Token               string
	EncodingAESKey      string
	SelfFuncAccessToken func() (resAccessToken context.ResAccessToken, err error) // 自定义授权token获取方法

	Cache cache.Cache
}

// 模型初始化
func NewWechat(cfg *Config) *Wechat {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &Wechat{context}
}

func copyConfigToContext(cfg *Config, context *context.Context) {
	context.AppID = cfg.AppID
	context.AppSecret = cfg.AppSecret
	context.Token = cfg.Token
	context.EncodingAESKey = cfg.EncodingAESKey
	context.SelfFuncAccessToken = cfg.SelfFuncAccessToken
	context.Cache = cfg.Cache
	context.SetAccessTokenLock(new(sync.RWMutex))
	context.SetJsAPITicketLock(new(sync.RWMutex))
}

// GetMaterial 素材管理
func (wc *Wechat) GetMaterial() *material.Material {
	return material.NewMaterial(wc.Context)
}

// GetServer 消息管理
func (wc *Wechat) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return server.NewServer(wc.Context)
}

//GetAccessToken 获取access_token
func (wc *Wechat) GetAccessToken() (string, error) {
	return wc.Context.GetAccessToken()
}

// GetUser 用户管理接口
func (wc *Wechat) GetUser() *user.User {
	return user.NewUser(wc.Context)
}

//GetKf 客服管理接口
func (wc *Wechat) GetKf() *kf.Kf {
	return kf.NewCustomerServer(wc.Context)
}
