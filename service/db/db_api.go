package db

import (
	. "example/homework/chatapp/utils"
	"slices"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	colName string
)

const (
	DefaultColName colName = "test"
	UserMsgColName colName = "user_msg"
	BookColName    colName = "book"
)

func (name colName) str() string {
	return string(name)
}

func Db() *mongo.Database {
	return DbManager.dbClient.Database(DbManager.dbCfg.DBName)
}

func Col(name colName) *mongo.Collection {
	return Db().Collection(name.str())
}

func InitDefaultCollection() error {
	names, err := ListCols()
	if err != nil {
		return err
	}
	if !slices.Contains(names, DefaultColName.str()) {
		if err = CreateCol(DefaultColName); err != nil {
			return err
		}
	}
	return nil
}

func ListCols() ([]string, error) {
	names, err := Db().ListCollectionNames(DbManager.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	return names, nil
}

func CreateCol(name colName) error {
	err := Db().CreateCollection(DbManager.ctx, name.str())
	if err != nil {
		return err
	}
	return nil
}

func InsertDoc(name colName, doc dbModel) error {
	cl := Col(name)
	result, err := cl.InsertOne(DbManager.ctx, doc)
	if err != nil {
		return err
	}
	doc.SetID(result.InsertedID.(primitive.ObjectID))
	return nil
}

func QueryDoc(name colName, filter bson.D) ([]UserMsgModel, error) {
	cl := Col(name)
	cursor, err := cl.Find(DbManager.ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []UserMsgModel
	if err = cursor.All(DbManager.ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func QueryUserMsg(name colName, filter bson.D) ([]UserMsgModel, error) {
	cl := Col(name)
	cursor, err := cl.Find(DbManager.ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []UserMsgModel
	if err = cursor.All(DbManager.ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func DeleteDoc(name colName, filter bson.D) (*mongo.DeleteResult, error) {
	cl := Col(name)
	result, err := cl.DeleteMany(DbManager.ctx, filter, nil)
	if err != nil {
		return nil, err
	}
	Logger.Println("Number of documents deleted", result.DeletedCount, "in collection", name)
	return result, nil
}
