package cmd

import (
	. "example/homework/chatapp/utils"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chatapp",
	Short: "App for chat",
	Long:  `Simple application with linebot and mongoDB`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		Logger.Println(err)
		os.Exit(1)
	}
}
