package cmd

import (
	"context"
	"example/homework/chatapp/service/line"
	. "example/homework/chatapp/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
				need ACTION, ex: send, query`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		switch args[0] {
		case "send":
			err = actionSend(user, text)
		default:
			err = fmt.Errorf("Failed to SendMsg no valid ACTION")
			Logger.Println(err)
		}
		if err != nil {
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
		return fmt.Errorf("Failed to SendMsg %s", err)
	}
	return nil
}
