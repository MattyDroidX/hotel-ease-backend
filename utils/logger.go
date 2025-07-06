package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *log.Logger

func InitLogger() {
	
	_ = os.MkdirAll("logs", os.ModePerm)

	rotatingFile := &lumberjack.Logger{
		Filename:   filepath.Join("logs", "app.log"),
		MaxSize:    5,  // MB
		MaxBackups: 3,
		MaxAge:     28, // dias
		Compress:   true,
	}

	multi := io.MultiWriter(os.Stdout, rotatingFile)

	Logger = log.New(multi, "", log.LstdFlags|log.Lshortfile)
}


func Info(format string, v ...interface{})  { Logger.Printf("[INFO]  "+format, v...) }
func Warn(format string, v ...interface{})  { Logger.Printf("[WARN]  "+format, v...) }
func Error(format string, v ...interface{}) { Logger.Printf("[ERROR] "+format, v...) }
func Infof(format string, a ...any)  { Logger.Printf("[INFO]  "+format, a...) }
func Warnf(format string, a ...any)  { Logger.Printf("[WARN]  "+format, a...) }
func Errorf(format string, a ...any) { Logger.Printf("[ERROR] "+format, a...) }
