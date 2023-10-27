package main

import (
	"context"
	"example/homework/chatapp/service"
	_ "example/homework/chatapp/service"
	"example/homework/chatapp/service/db"
	. "example/homework/chatapp/utils"
)

func main() {
	Logger.Println("Chat app was started.")
	ctx, cancel := context.WithCancel(context.Background())
	db.InitDb(ctx)
	dbMngr := db.GetDbManager()
	for {
		if dbMngr.IsInited() {
			break
		}
	}
	service.InitLineBot(ctx)
	lineMngr := service.GetLineManager()
	for {
		if lineMngr.IsInited() {
			break
		}
	}
	service.InitApp()
	appMngr := service.GetAppManager()
	for {
		if appMngr.IsInited() {
			break
		}
	}
	defer func() {
		if err := recover(); err != nil {
			cancel()
			Logger.Println("Panic happended", err)
		}
		Logger.Println("Chat app was finish.")
	}()
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}
