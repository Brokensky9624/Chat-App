package service

import (
	"context"
	. "example/homework/chatapp/utils"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var (
	LineManager *lineManager
)

type lineManager struct {
	lineCfg  *LineConfig
	ctx      context.Context
	bot      *linebot.Client
	isInited bool
}

func InitLineBot(ctx context.Context) {
	mngr := NewLineManager(ctx)
	mngr.Run()
}

func NewLineManager(ctx context.Context) *lineManager {
	lineCfg := LoadLineCfg()
	LineManager = &lineManager{
		lineCfg: lineCfg,
		ctx:     ctx,
	}
	return LineManager
}

func GetLineManager() *lineManager {
	return LineManager
}

func (man *lineManager) Run() {
	go func() {
		bot, err := linebot.New(
			man.lineCfg.LineSecret,
			man.lineCfg.LineToken,
		)
		if err != nil {
			Logger.Panicln("Failed to connect to line bot", err)
		}
		man.bot = bot
		man.isInited = true
		for {
			select {
			case <-man.ctx.Done():
				return
			}
		}
	}()
}

func (man lineManager) IsInited() bool {
	return man.isInited
}

func (man *lineManager) ParseRequest(w http.ResponseWriter, req *http.Request) {
	events, err := LineManager.bot.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = LineManager.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					Logger.Println("Failed to reply message by line bot", err)
				}
			}
		}
	}
}
