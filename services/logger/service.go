package logger

import (
	"crypto-wallet/config"
	"crypto-wallet/pkg/logger"
)

type LogService interface {
	Info(message string)
	Warning(message string)
	Error(message string)
	Dev(message string)
}

type logService struct {
	c config.Config
}

func (l logService) Info(message string) {
	logger.Write("info.log", message)
}

func (l logService) Warning(message string) {
	logger.Write("warning.log", message)
}

func (l logService) Error(message string) {
	logger.Write("error.log", message)
}

func (l logService) Dev(message string) {
	logger.Write("dev.log", message)
}

func NewLogService(c config.Config) LogService {
	return &logService{c}
}
