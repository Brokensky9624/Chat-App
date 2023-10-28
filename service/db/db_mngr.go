package db

import (
	"context"
	"example/homework/chatapp/service"
	. "example/homework/chatapp/utils"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uriTemplate = "mongodb://<host>:<port>"
)

type dbManager struct {
	dbCfg    *service.MongoDBConfig
	dbClient *mongo.Client
	dbOpts   *options.ClientOptions
	ctx      context.Context
	isInited bool
}

func NewDbManager(ctx context.Context) *dbManager {
	dbCfg := service.LoadDbCfg()
	uri := strings.Replace(uriTemplate, "<host>", dbCfg.DBHost, 1)
	uri = strings.Replace(uri, "<port>", dbCfg.DBPort, 1)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	dbOpts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	dbOpts.Auth = &options.Credential{
		Username: dbCfg.DBUsername,
		Password: dbCfg.DBPassword,
	}
	return &dbManager{
		dbCfg:  dbCfg,
		dbOpts: dbOpts,
		ctx:    ctx,
	}
}

func (man *dbManager) Run() {
	go func() {
		man.dbOpts.MaxPoolSize = man.dbCfg.DBMaxPoolSize
		ctx1, cancel := context.WithTimeout(man.ctx, 10*time.Second)
		dbClient, err := mongo.Connect(ctx1, man.dbOpts)
		if err != nil {
			Logger.Panicln("Failed to connect to mongoDB", err)
		}
		man.dbClient = dbClient
		defer cancel()
		defer func() {
			if err = man.dbClient.Disconnect(context.TODO()); err != nil {
				Logger.Panicln("Failed to disconnect from mongoDB", err)
			}
		}()
		var result bson.M
		ctx2, cancel := context.WithTimeout(man.ctx, 10*time.Second)
		if err := man.dbClient.Database(man.dbCfg.DBName).RunCommand(ctx2, bson.D{{"ping", 1}}).Decode(&result); err != nil {
			Logger.Panicln("Failed to send command from mongoDB", err)
		}
		defer cancel()
		err = InitDefaultCollection()
		if err != nil {
			Logger.Panicf("Failed to init default collection %s in mongoDB %s\n", DefaultColName, err)
		}
		man.isInited = true
		Logger.Println("Succeed to connect to mongoDB")
		for {
			select {
			case <-man.ctx.Done():
				return
			}
		}
	}()
}

func (man dbManager) IsInited() bool {
	return man.isInited
}
