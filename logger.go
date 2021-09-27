package main

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

func getLogger() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   "C:/Users/81905/Desktop/gin-starter/_log/go.log",
		MaxSize:    1, // megabytes
		MaxBackups: 60,
		MaxAge:     60, //days
		Compress:   false,
	}
}
