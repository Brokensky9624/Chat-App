package line

import (
	"encoding/json"
	"example/homework/chatapp/service/db"
	. "example/homework/chatapp/utils"
	"fmt"
	"net/http"
	"regexp"

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
				err := HandleText(event, message)
				if err != nil {
					Logger.Println("Failed to HandleText in line_api", err)
				}
			}
		}
	}
	return nil
}

func HandleText(event *linebot.Event, textMsg *linebot.TextMessage) error {
	user := event.Source.UserID
	msg := textMsg.Text
	var err error
	var ret []byte
	switch {
	case regexp.MustCompile(`^search$`).FindString(msg) != "":
		ret, err = QueryUserMsg(user)
	case regexp.MustCompile("^delete$").FindString(msg) != "":
		ret, err = deleteUserMsg(user)
	default:
		ret, err = insertUserMsg(user, msg)
	}
	if err != nil {
		return err
	}
	if string(ret) != "" {
		if err = ReplyMessage(event.ReplyToken, linebot.NewTextMessage(string(ret))); err != nil {
			return err
		}
	}
	return nil
}

func QueryUserMsg(user string) ([]byte, error) {
	userMsgModels, err := db.QueryUserMsg(db.UserMsgColName, bson.D{primitive.E{Key: "user", Value: user}})
	if err != nil {
		return nil, err
	}
	b, err := json.MarshalIndent(userMsgModels, "", "\t")
	if err != nil {
		return nil, err
	}
	Logger.Println(string(b))
	return b, nil
}

func insertUserMsg(user, text string) ([]byte, error) {
	userDoc := db.NewUserMsgBaseModel(user, text)
	if err := db.InsertDoc(db.UserMsgColName, userDoc); err != nil {
		return nil, err
	}
	ret := fmt.Sprintf(`{col:%s, insertedId: %s}`, db.UserMsgColName, userDoc.GetID().Hex())
	Logger.Println(ret)
	return []byte(ret), nil
}

func deleteUserMsg(user string) ([]byte, error) {
	filter := bson.D{primitive.E{Key: "user", Value: user}}
	result, err := db.DeleteDoc(db.UserMsgColName, filter)
	if err != nil {
		return nil, err
	}
	ret := fmt.Sprintf(`{col:%s, deleteCount: %d}`, db.UserMsgColName, result.DeletedCount)
	return []byte(ret), nil
}

func ReplyMessage(replyToken string, message linebot.SendingMessage) error {
	mngr := LineManager
	bot := mngr.Bot()
	_, err := bot.ReplyMessage(replyToken, message).Do()
	if err != nil {
		return err
	}
	return nil
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
