package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

var (
	Logger *logrus.Logger

	LevelList = map[string]logrus.Level{
		"PANIC": logrus.PanicLevel,
		"FATAL": logrus.FatalLevel,
		"ERROR": logrus.ErrorLevel,
		"WARN":  logrus.WarnLevel,
		"INFO":  logrus.InfoLevel,
		"DEBUG": logrus.DebugLevel,
		"TRACE": logrus.TraceLevel,
	}
)

// ***********************************************************************
func LogToFile(FileName string) (error, io.Writer) {
	var (
		f   io.Writer
		err error = nil
	)
	f, err = os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Error opening log file: " + err.Error())
		return err, f
	}
	return err, f
}

// ***********************************************************************
func LogToStdout() io.Writer {
	return os.Stdout
}

// ***********************************************************************
func LogLevel(DebugLevel string) {
	Logger.SetLevel(str2lvl(DebugLevel))
	Logger.Infoln("Logging level: " + DebugLevel)
}

// ***********************************************************************
func InitLogger(DebugLevel string, FileName string) {
	var (
		LogStream io.Writer
		err       error
	)

	Logger = logrus.New()
	if strings.ToLower(FileName) == "stdout" {
		Logger.SetOutput(os.Stdout)
	} else {
		err, LogStream = LogToFile(FileName)
		if err == nil {
			Logger.SetOutput(LogStream)
		} else {
			Logger.SetOutput(os.Stdout)
		}
	}

	Logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	LogLevel(DebugLevel)
	Logger.Infoln("Logging initialized")
}

// ***********************************************************************
func str2lvl(StrLvl string) logrus.Level {
	if _, ok := LevelList[strings.ToUpper(StrLvl)]; !ok {
		fmt.Println("ERROR: Unknown logging level (" + StrLvl + "). Set default (DEBUG)")
		return logrus.DebugLevel
	}
	return LevelList[strings.ToUpper(StrLvl)]
}
