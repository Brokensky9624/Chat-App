package line

import (
	"encoding/json"
	"example/homework/chatapp/service/db"
	. "example/homework/chatapp/utils"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParseRequest(w http.ResponseWriter, req *http.Request) error {
	mngr := LineManager
	bot := mngr.Bot()
	events, err := bot.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return err
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				retList, err := HandleText(event, message)
				retStr := strings.Join(retList, "\n")
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(retStr)).Do()
				if err != nil {
					Logger.Println("Failed to reply message by line bot", err)
				}
			}
		}
	}
	return nil
}

func HandleText(event *linebot.Event, textMsg *linebot.TextMessage) ([]string, error) {
	user := event.Source.UserID
	msg := textMsg.Text
	var retList []string
	var err error
	switch {
	case regexp.MustCompile(`^search`).FindString(msg) != "":
		if retList, err = queryUser(user); err != nil {
			return nil, err
		}
	default:
		if retList, err = insertUserMsg(user, msg); err != nil {
			return nil, err
		}
	}
	return retList, nil
}

func queryUser(user string) ([]string, error) {
	var retList []string
	userMsgModels, err := db.FindUserDoc(db.UserMsgColName, bson.D{primitive.E{Key: "user", Value: user}})
	if err != nil {
		return nil, err
	}
	for _, userMsgModel := range userMsgModels {
		b, err := json.Marshal(userMsgModel)
		if err != nil {
			return nil, err
		}
		ret := fmt.Sprintf("Success query doc %s from mongoDB col %s", string(b), db.UserMsgColName)
		retList = append(retList, ret)
		Logger.Println(ret)

	}
	return retList, nil
}

func insertUserMsg(user, text string) ([]string, error) {
	userDoc := db.NewUserMsgBaseModel(user, text)
	if err := db.InsertDoc(db.UserMsgColName, userDoc); err != nil {
		return nil, err
	}
	ret := fmt.Sprintf("Success insert doc %s to mongoDB col %s", userDoc.GetID().Hex(), db.UserMsgColName)
	Logger.Println(ret)
	return []string{ret}, nil
}

func SendMsg(inputUser, inputText string) error {
	mngr := LineManager
	bot := mngr.Bot()
	if inputText == "" {
		inputText = "N\\A"
	}
	if _, err := bot.PushMessage(inputUser, linebot.NewTextMessage(inputText)).Do(); err != nil {
		return err
	}
	return nil
}
