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

type userBaseModel struct {
	id      primitive.ObjectID
	User    string    `bson:"user"`
	Message string    `bson:"message"`
	Recv    time.Time `bson:"recv"`
	Reply   time.Time `bson:"reply"`
}

func NewUserBaseModel(user, message string) *userBaseModel {
	return &userBaseModel{
		User:    user,
		Message: message,
		Recv:    time.Now(),
	}
}

func (model userBaseModel) GetID() primitive.ObjectID {
	return model.id
}

func (model *userBaseModel) SetID(id primitive.ObjectID) {
	model.id = id
}

type userModel struct {
	ID      primitive.ObjectID `bson:"_id"`
	User    string             `bson:"user"`
	Message string             `bson:"message"`
	Recv    time.Time          `bson:"recv"`
	Reply   time.Time          `bson:"reply"`
}

func (model userModel) GetUser() string {
	return model.User
}

func (model userModel) GetMessage() string {
	return model.Message
}

func (model userModel) GetRecv() string {
	return model.Recv.Format(TimeFormat)
}

func (model userModel) GetReply() string {
	return model.Reply.Format(TimeFormat)
}

func (model *userModel) SetReply() {
	model.Reply = time.Now()
}

func (model userModel) GetID() primitive.ObjectID {
	return model.ID
}

func (model *userModel) SetID(id primitive.ObjectID) {
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
