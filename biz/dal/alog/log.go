package alog

import (
	"fmt"
	"io"
	"jobor/pkg/utils"
	"log"
	"os"
	"path"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzZerolog "github.com/hertz-contrib/logger/zerolog"

	//hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	traceIDKey    = "trace_id"
	spanIDKey     = "span_id"
	traceFlagsKey = "trace_flags"
	logEventKey   = "log"
)

var Dlog hlog.FullLogger

func InitHzLog(LogFileName string, level hlog.Level) {
	lumberjackLogger := getLumberjackLogger(LogFileName)
	iw := io.MultiWriter(lumberjackLogger, os.Stdout) // os.Stdout, logger.Gin.Writer()
	//hlog.DefaultLogger().SetOutput(iw)
	//hlog.DefaultLogger().SetLevel(level)
	//hlog.SetLogger(hertzlogrus.NewLogger())

	//var entry *logrus.Entry
	//if entry.Context == nil {
	//	//return nil
	//}
	//
	//span := trace.SpanFromContext(entry.Context)
	//
	//// attach span context to log entry data fields
	//entry.Data[traceIDKey] = span.SpanContext().TraceID()
	//entry.Data[spanIDKey] = span.SpanContext().SpanID()
	//entry.Data[traceFlagsKey] = span.SpanContext().TraceFlags()
	hlog.SetOutput(iw)
	hlog.SetLevel(level)
}

func InitZeroLog(logFilePath string, level hlog.Level) {
	// Customizable output directory.
	//var logFilePath string
	//dir := "./hlog"
	//logFilePath = dir + "/logs/"
	logPath := path.Dir(logFilePath)
	if err := os.MkdirAll(logPath, 0o777); err != nil {
		log.Println(err.Error())
		return
	}

	// Set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logPath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return
		}
	}

	// Provides compression and deletion
	lumberjackLogger := getLumberjackLogger(fileName)

	iw := io.MultiWriter(lumberjackLogger, os.Stdout) // os.Stdout, logger.Gin.Writer()

	// For logrus detailed settings, please refer to https://github.com/hertz-contrib/logger/tree/main/logrus and https://github.com/sirupsen/logrus
	//hertzZerolog.MultiLevelWriter(os.Stdout)
	//logger := hertzZerolog.New()
	//logger1, _ := hertzZerolog.GetLogger()
	logger := hertzZerolog.New(
		//hertzZerolog.WithOutput(iw),   // allows to specify output
		//hertzZerolog.WithLevel(level), // option with log level
		hertzZerolog.WithTimestamp(), // option with timestamp
		//hertzZerolog.WithCaller(),              // option with caller
		// ...
	)
	logger.SetOutput(iw)
	logger.SetLevel(level)

	hlog.SetLogger(logger)
}

func getLumberjackLogger(fileName string) *lumberjack.Logger {
	if err := InitOutToFile(fileName); err != nil {
		panic(err)
	}
	// Provides compression and deletion
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}
	return lumberjackLogger
}

func GetFileIO(LogFileName string) *os.File {
	_ = InitOutToFile(LogFileName)
	f, err := os.OpenFile(LogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	//defer f.Close()
	return f
}

func InitOutToFile(logFilePath string) error {
	//logFilePath := path.Join(logPath, logFileName)
	logDirPath := path.Dir(logFilePath)
	if !utils.FileExists(logDirPath) {
		err := os.MkdirAll(logDirPath, os.ModePerm)
		if err != nil {
			msg := fmt.Sprintf("mkdir log dir %s error: %s\n", logDirPath, err)
			log.Println(msg)
			return err
		}
	}
	if _, err := os.Stat(logFilePath); err != nil {
		if _, err := os.Create(logFilePath); err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return nil
}
