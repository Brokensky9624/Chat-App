package line

import (
	"context"
	"example/homework/chatapp/service"
	. "example/homework/chatapp/utils"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type lineManager struct {
	lineCfg  *service.LineConfig
	ctx      context.Context
	bot      *linebot.Client
	isInited bool
}

func NewLineManager(ctx context.Context) *lineManager {
	lineCfg := service.LoadLineCfg()
	return &lineManager{
		lineCfg: lineCfg,
		ctx:     ctx,
	}
}

func (mngr *lineManager) Run() {
	go func() {
		bot, err := linebot.New(
			mngr.lineCfg.LineSecret,
			mngr.lineCfg.LineToken,
		)
		if err != nil {
			Logger.Panicln("Failed to connect to line bot", err)
		}
		mngr.bot = bot
		mngr.isInited = true
		for {
			select {
			case <-mngr.ctx.Done():
				return
			}
		}
	}()
}

func (mngr lineManager) IsInited() bool {
	return mngr.isInited
}

func (mngr lineManager) Bot() *linebot.Client {
	return mngr.bot
}
