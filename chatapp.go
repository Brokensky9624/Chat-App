package main

import (
	"context"
	_ "example/homework/chatapp/service"
	"example/homework/chatapp/service/db"
	"example/homework/chatapp/service/line"
	"example/homework/chatapp/service/router"
	"example/homework/chatapp/service/web"
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
	line.InitLineBot(ctx)
	lineMngr := line.GetLineManager()
	for {
		if lineMngr.IsInited() {
			break
		}
	}
	web.InitWeb(ctx)
	webMngr := web.GetWebManager()
	for {
		if webMngr.IsInited() {
			break
		}
	}
	router.InitAllRouter()

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
