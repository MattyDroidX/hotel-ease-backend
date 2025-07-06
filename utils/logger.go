package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	console *log.Logger
	fileLog *log.Logger
	once    sync.Once
)

const (
	clReset = "\033[0m"
	clRed   = "\033[31m"
	clYel   = "\033[33m"
	clCya   = "\033[36m"
)

func InitLogger() {
	once.Do(initLoggers)
}

func initLoggers() {
	_ = os.MkdirAll("logs", os.ModePerm)

	fileLog = log.New(&lumberjack.Logger{
		Filename:   filepath.Join("logs", "app.log"),
		MaxSize:    5,  // MB
		MaxBackups: 3,
		MaxAge:     28, // dias
		Compress:   true,
	}, "", log.LstdFlags|log.Lshortfile)

	console = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}


func Info(format string, v ...any)  { write("INFO", clCya, format, v...) }
func Warn(format string, v ...any)  { write("WARN", clYel, format, v...) }
func Error(format string, v ...any) { write("ERROR", clRed, format, v...) }


func write(level, color, format string, v ...any) {
	InitLogger()

	msg := fmt.Sprintf(format, v...)

	if console != nil {
		console.Printf("[%s%s%s] %s", color, level, clReset, msg)
	}

	if fileLog != nil {
		fileLog.Printf("[%s] %s", level, msg)
	}
}
