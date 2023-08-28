package logger

import (
	"fmt"
	"jobor/conf"
	"jobor/pkg/utils"
	"log"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
	Jobor  = logrus.New()
	Casbin = logrus.New()
	Gin    = logrus.New()
	L400X  = logrus.New()
	L500X  = logrus.New()
)
var logLevels = map[string]logrus.Level{
	"DEBUG": logrus.DebugLevel,
	"INFO":  logrus.InfoLevel,
	"WARN":  logrus.WarnLevel,
	"ERROR": logrus.ErrorLevel,
}

func Initial() {
	formatter := &Formatter{
		LogFormat: "",
		//LogFormat:       "%time% [%lvl%] %msg%",
		TimestampFormat: "2006-01-02 15:04:05",
	}
	ginFormatter := &Formatter{
		LogFormat:       "%msg%",
		TimestampFormat: "2006-01-02 15:04:05",
	}
	InitLog("jobor.log", formatter, Jobor)
	InitLog("std.log", formatter, logger)
	InitLog("400.log", formatter, L400X)
	InitLog("500.log", formatter, L500X)
	InitLog("casbin.log", formatter, Casbin)
	InitLog("gin.log", ginFormatter, Gin)

}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func InitOutToFile(logPath, logFileName string) error {
	logFilePath := path.Join(logPath, logFileName)
	logDirPath := path.Dir(logFilePath)
	if !utils.FileExists(logDirPath) {
		err := os.MkdirAll(logDirPath, os.ModePerm)
		if err != nil {
			msg := fmt.Sprintf("K8sDeployment log dir %s error: %s\n", logDirPath, err)
			fmt.Println(msg)
			return err
		}
	}
	return nil
}

func InitLog(logFileName string, formatter *Formatter, loggerObj *logrus.Logger) {

	c := conf.GetConf()
	level, ok := logLevels[strings.ToUpper(c.Server.LogLevel)]
	if !ok {
		level = logrus.InfoLevel
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	loggerObj.SetFormatter(formatter)
	loggerObj.SetOutput(os.Stdout)
	loggerObj.SetLevel(level)

	// Output to file
	logFilePath := path.Join(c.Server.LogFileName, logFileName)
	if err := InitOutToFile(c.Server.LogFileName, logFileName); err != nil {
		log.Fatal(err)
	}
	rotateFileHook, err := NewRotateFileHook(RotateFileConfig{
		Filename:   logFilePath,
		MaxSize:    50,
		MaxBackups: 7,
		MaxAge:     7,
		LocalTime:  true,
		Level:      level,
		Formatter:  formatter,
	})
	if err != nil {
		fmt.Printf("log rotate hook error: %s\n", err)
		return
	}
	loggerObj.AddHook(rotateFileHook)
}

//gin log file
//logFilePath := path.Join(config.Configs.Server.LogPath, "gin.log")
//if err:=logger.InitOutToFile(config.Configs.Server.LogPath, "gin.log");err!=nil{
//	log.Fatalln(err)
//}
//
//logFile,err:=os.OpenFile(logFilePath,os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
//if err!=nil{log.Fatalln(err)}
