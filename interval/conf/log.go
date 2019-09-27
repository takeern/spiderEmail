package conf

import (
	log "github.com/sirupsen/logrus"
	
)

type UTCFormatter struct {
    log.Formatter
}

func (u UTCFormatter) Format(e *log.Entry) ([]byte, error) {
    e.Time = e.Time.UTC()
    return u.Formatter.Format(e)
}

func init() {
	log.SetFormatter(UTCFormatter{&log.JSONFormatter{}})
	log.SetLevel(log.DebugLevel)
	log.Info("something")
}