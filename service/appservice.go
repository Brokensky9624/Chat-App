package service

import (
	. "example/homework/chatapp/utils"
	"net/http"
)

var (
	AppManager *appManager
)

func InitApp() {
	mngr := NewAppManager()
	mngr.Run()
}

type appManager struct {
	appCfg   *AppConfig
	isInited bool
}

func NewAppManager() *appManager {
	appCfg := LoadAppCfg()
	AppManager = &appManager{
		appCfg: appCfg,
	}
	defer func() {
		AppManager.isInited = true
	}()
	return AppManager
}

func GetAppManager() *appManager {
	return AppManager
}

func (man *appManager) Run() {
	go func() {
		// Setup HTTP Server for receiving requests from LINE platform
		http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
			LineManager.ParseRequest(w, req)
		})
		Logger.Printf("Listen and serve %s:%s\n", man.appCfg.Host, man.appCfg.Port)
		if err := http.ListenAndServe(":"+man.appCfg.Port, nil); err != nil {
			Logger.Panicln("Failed to read config of db by viper", err)
		}
	}()
}

func (man appManager) IsInited() bool {
	return man.isInited
}
