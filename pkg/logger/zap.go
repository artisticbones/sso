package logger

import (
	"go.uber.org/zap"
	"sync"
)

type Logger struct {
	*zap.Logger
	sync.Once
}

var _lg *Logger

// Default function initialize a default instance
func Default() {
}

// GetLogger obtain the global logger
func GetLogger() *Logger { return _lg }
