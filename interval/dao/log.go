package dao

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
	"fmt"
)


var Log = logrus.New()

func init() {
	lasttime := time.Now().Format("2006-01-02")
	logPath := "/" + lasttime + "-spider.log"
	fmt.Println(logPath)
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY, 0666)
	Log.SetLevel(logrus.DebugLevel)
	Log.SetFormatter(&logrus.JSONFormatter{})
	if err == nil {
		Log.Out = file
	} else {
		Log.Info("Failed to log to file, using default stderr")
	}

	Log.Info("something")
}