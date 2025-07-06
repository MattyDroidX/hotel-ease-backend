package utils

import (
	"fmt"
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *log.Logger

func InitLogger() {
	Logger = log.New(&lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    5, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   // dias
		Compress:   true, // gzip
	}, "", log.LstdFlags|log.Lshortfile)
}

func Info(format string, args ...interface{}) {
	Logger.Println("[INFO] " + fmt.Sprintf(format, args...))
}

func Warn(format string, args ...interface{}) {
	Logger.Println("[WARN] " + fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}) {
	Logger.Println("[ERROR] " + fmt.Sprintf(format, args...))
}
