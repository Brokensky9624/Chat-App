package cmd

import (
	"context"
	"example/homework/chatapp/service/db"
	"example/homework/chatapp/service/line"
	"example/homework/chatapp/service/router"
	"example/homework/chatapp/service/web"
	. "example/homework/chatapp/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run app as web server",
	Long: `Run app as web server. 
				Web server works with linebot webhook`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		finish := make(chan int, 1)
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
			finish <- 1
			Logger.Println("Chat app was finish.")
		}()
		for {
			select {
			case <-finish:
				return
			}
		}
	},
}
