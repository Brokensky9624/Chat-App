package db

import (
	"context"
	"sync"
)

var (
	mu        sync.RWMutex
	DbManager *dbManager
)

func InitDb(ctx context.Context) {
	DbManager = NewDbManager(ctx)
	DbManager.Run()
}

func GetDbManager() *dbManager {
	mu.RLocker().Lock()
	defer mu.RLocker().Unlock()
	return DbManager
}
