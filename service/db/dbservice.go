package db

import (
	"context"
)

var (
	DbManager *dbManager
)

func InitDb(ctx context.Context) {
	DbManager = NewDbManager(ctx)
	DbManager.Run()
}

func GetDbManager() *dbManager {
	return DbManager
}
