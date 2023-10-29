package db

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertDoc(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	InitDb(ctx)
	for {
		if DbManager.isInited {
			break
		}
	}
	type args struct {
		name colName
		doc  dbModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "save user info to mongodb",
			args: args{
				name: UserMsgColName,
				doc:  NewUserMsgBaseModel("Jason", "Hello MongoDB"),
			},
			wantErr: false,
		},
		{
			name: "save book info to mongodb",
			args: args{
				name: BookColName,
				doc: &bookBaseModel{
					Name:   "Golang doc",
					Author: "someone",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertDoc(tt.args.name, tt.args.doc); (err != nil) != tt.wantErr {
				t.Errorf("InsertDoc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFindDoc(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	InitDb(ctx)
	for {
		if DbManager.isInited {
			break
		}
	}
	type args struct {
		name   colName
		filter bson.D
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "query user info from mongodb",
			args: args{
				name:   UserMsgColName,
				filter: bson.D{primitive.E{Key: "user", Value: "Jason"}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := QueryUserMsg(tt.args.name, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryUserMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
