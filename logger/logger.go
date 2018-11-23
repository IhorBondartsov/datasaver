package logger

import (
	"os"
	"time"

	"github.com/IhorBondartsov/datasaver/cfg"
	"github.com/sirupsen/logrus"
)

func CreateLogger() *logrus.Logger {
	log := logrus.New()
	log.Level = cfg.LogLevel
	file, err := os.OpenFile("logs/"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Println("Failed to log to file, using default stderr")
	}
	return log
}

func CreateLoggerWithRecreater() *logrus.Logger {
	log := CreateLogger()
	go recreateLogger(log)
	return log
}

// recreateLogsEveryDay - change log output every 24 hours
func recreateLogsEveryDay(log *logrus.Logger) {
	tick := time.NewTicker(24 * time.Hour)
	for t := range tick.C {
		setOutputToLog(log)
		log.Info("Recreate log %v", t)
	}
}

// setOutputToLog create new output for logger
func setOutputToLog(log *logrus.Logger) {
	file, err := os.OpenFile("logs/"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Println("Failed to log to file, using default stderr")
	}
}

// Recreate logger every midnight by UNIX
func recreateLogger(log *logrus.Logger) {
	var (
		timeInDay          int64 = 86400
		spentAfterMidNight       = time.Now().Unix() % timeInDay
		leftTime                 = time.Duration(timeInDay - spentAfterMidNight)
	)
	time.Sleep(leftTime * time.Second)
	setOutputToLog(log)
	recreateLogsEveryDay(log)
}
