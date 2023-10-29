package cmd

import (
	"context"
	"example/homework/chatapp/service/db"
	"example/homework/chatapp/service/line"
	. "example/homework/chatapp/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	user string
	text string
)

func init() {
	rootCmd.AddCommand(messageCmd)
	messageCmd.Flags().StringVarP(&user, "user", "u", "", "User ID of line (required)")
	messageCmd.MarkFlagRequired("user")
	messageCmd.Flags().StringVarP(&text, "text", "t", "", "Text message content")
}

var messageCmd = &cobra.Command{
	Use:   "message <ACTION>",
	Short: "handle message event",
	Long: `Handle message event. 
				need ACTION, ex: send, query, delete`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		switch args[0] {
		case "send":
			err = actionSend(user, text)
		case "query":
			err = actionQuery(user)
		case "delete":
			err = actionDelete(user)
		default:
			err = fmt.Errorf("Failed to SendMsg no valid ACTION")
		}
		if err != nil {
			Logger.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func actionSend(inputUser, inputText string) error {
	ctx, cancel := context.WithCancel(context.Background())
	line.InitLineBot(ctx)
	lineMngr := line.GetLineManager()
	for {
		if lineMngr.IsInited() {
			break
		}
	}
	defer func() {
		if err := recover(); err != nil {
			cancel()
			Logger.Println("Panic happended", err)
			os.Exit(1)
		}
	}()
	if err := line.SendMsg(inputUser, inputText); err != nil {
		return fmt.Errorf("Failed to message actionSend %s %s %s", inputUser, inputText, err)
	}
	Logger.Println("Succeed to message actionSend", inputUser, inputText)
	return nil
}

func actionQuery(inputUser string) error {
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
	defer func() {
		if err := recover(); err != nil {
			cancel()
			Logger.Println("Panic happended", err)
			os.Exit(1)
		}
	}()
	ret, err := line.QueryUserMsg(inputUser)
	if err != nil {
		return fmt.Errorf("Failed to message actionQuery %s %s", inputUser, err)
	}
	Logger.Println(string(ret))
	return nil
}

func actionDelete(inputUser string) error {
	ctx, cancel := context.WithCancel(context.Background())
	db.InitDb(ctx)
	dbMngr := db.GetDbManager()
	for {
		if dbMngr.IsInited() {
			break
		}
	}
	defer func() {
		if err := recover(); err != nil {
			cancel()
			Logger.Println("Panic happended", err)
			os.Exit(1)
		}
	}()
	filter := bson.D{primitive.E{Key: "user", Value: inputUser}}
	_, err := db.DeleteDoc(db.UserMsgColName, filter)
	if err != nil {
		return fmt.Errorf("Failed to message actionDelete %s %s", inputUser, err)
	}
	Logger.Println("Succeed to message actionDelete", inputUser)
	return nil
}
