package web

import (
	// . "example/homework/chatapp/utils"
	"context"
	"example/homework/chatapp/service"
	"sync"
)

const (
	servercrt = "server.crt"
	serverkey = "server.key"
)

var (
	mu         sync.RWMutex
	WebManager *webManager
)

func InitWeb(ctx context.Context) {
	appCfg := service.LoadAppCfg()
	mngr := &webManager{
		appCfg: appCfg,
		ctx:    ctx,
	}
	mngr.Run()
	WebManager = mngr
}

func GetWebManager() *webManager {
	mu.RLocker().Lock()
	defer mu.RLocker().Unlock()
	return WebManager
}
