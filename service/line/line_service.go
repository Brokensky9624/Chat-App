package line

import (
	"context"
	"sync"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var (
	mu          sync.RWMutex
	LineManager *lineManager
)

type lineTextMsg struct {
	Type       linebot.MessageType      `json:"type"`
	Text       string                   `json:"text"`
	QuickReply *linebot.QuickReplyItems `json:"quickReply,omitempty"`
	Sender     *linebot.Sender          `json:"sender,omitempty"`
	Emojis     []*linebot.Emoji         `json:"emojis,omitempty"`
}

func InitLineBot(ctx context.Context) {
	LineManager = NewLineManager(ctx)
	LineManager.Run()
}

func GetLineManager() *lineManager {
	mu.RLocker().Lock()
	defer mu.RLocker().Unlock()
	return LineManager
}
