package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	TimeFormat = time.RFC3339
)

type dbModel interface {
	GetID() primitive.ObjectID
	SetID(primitive.ObjectID)
}

type UserMsgBaseModel struct {
	id      primitive.ObjectID
	User    string    `bson:"user"`
	Message string    `bson:"message"`
	Recv    time.Time `bson:"recv"`
	Reply   time.Time `bson:"reply"`
}

func NewUserMsgBaseModel(user, message string) *UserMsgBaseModel {
	return &UserMsgBaseModel{
		User:    user,
		Message: message,
		Recv:    time.Now(),
	}
}

func (model UserMsgBaseModel) GetID() primitive.ObjectID {
	return model.id
}

func (model *UserMsgBaseModel) SetID(id primitive.ObjectID) {
	model.id = id
}

type UserMsgModel struct {
	ID      primitive.ObjectID `bson:"_id"`
	User    string             `bson:"user"`
	Message string             `bson:"message"`
	Recv    time.Time          `bson:"recv"`
	Reply   time.Time          `bson:"reply"`
}

func (model UserMsgModel) GetUser() string {
	return model.User
}

func (model UserMsgModel) GetMessage() string {
	return model.Message
}

func (model UserMsgModel) GetRecv() string {
	return model.Recv.Format(TimeFormat)
}

func (model UserMsgModel) GetReply() string {
	return model.Reply.Format(TimeFormat)
}

func (model *UserMsgModel) SetReply() {
	model.Reply = time.Now()
}

func (model UserMsgModel) GetID() primitive.ObjectID {
	return model.ID
}

func (model *UserMsgModel) SetID(id primitive.ObjectID) {
	model.ID = id
}

type bookBaseModel struct {
	id     primitive.ObjectID
	Name   string
	Author string
}

func (model bookBaseModel) GetID() primitive.ObjectID {
	return model.id
}

func (model *bookBaseModel) SetID(id primitive.ObjectID) {
	model.id = id
}

type bookModel struct {
	ID primitive.ObjectID `bson:"_id"`
	bookBaseModel
}

func (model bookModel) GetID() primitive.ObjectID {
	return model.ID
}

func (model *bookModel) SetID(id primitive.ObjectID) {
	model.ID = id
}
