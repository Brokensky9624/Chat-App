package service

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _  = runtime.Caller(0)
	projectPath = filepath.Dir(filepath.Dir(b))
)

func getProjectPath() string {
	return projectPath
}

func getConfigPath() string {
	return filepath.Join(projectPath, "config")
}

func getCertPath() string {
	return filepath.Join(projectPath, "cert")
}
