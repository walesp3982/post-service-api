package pkg

import (
	"log/slog"
	"os"
	"sync"
)

var singleLogger *slog.Logger
var lock = &sync.Mutex{}

func GetLogger() *slog.Logger {
	if singleLogger == nil {
		lock.Lock()
		defer lock.Unlock()
		singleLogger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	return singleLogger
}
