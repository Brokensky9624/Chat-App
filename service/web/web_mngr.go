package web

import (
	"context"
	"example/homework/chatapp/service"

	"github.com/gin-gonic/gin"
)

type webManager struct {
	appCfg    *service.AppConfig
	isInited  bool
	ctx       context.Context
	webEngine *gin.Engine
}

func (man *webManager) Run() {
	router := gin.Default()
	man.webEngine = router
	go router.Run(":" + man.appCfg.Port)
	man.isInited = true
}

func (man webManager) IsInited() bool {
	return man.isInited
}

func (man webManager) WebEngine() *gin.Engine {
	return man.webEngine
}
