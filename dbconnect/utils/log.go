package utils

import (
	"bufio"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/onrik/logrus/filename"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func InitLogger(logLevel string, logPath string, showSource bool) error {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)

	formatter := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.0000Z07:00",
	}

	log.SetFormatter(formatter)

	if showSource {
		filenameHook := filename.NewHook()
		filenameHook.Field = "source"
		log.AddHook(filenameHook)
	}

	if logPath == "" {
		log.SetOutput(os.Stdout)
		return nil
	}

	baseLogPath := logPath
	debugWriter, err := rotatelogs.New(
		baseLogPath+"debug.log.%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath+"debug.log"),
		rotatelogs.WithMaxAge(3*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		log.Errorf("config local file system logger error. %v", errors.WithStack(err))
		return err
	}

	writer, err := rotatelogs.New(
		baseLogPath+"info.log.%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath+"info.log"),
		rotatelogs.WithMaxAge(3*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		log.Errorf("config local file system logger error. %v", errors.WithStack(err))
		return err
	}

	switch level.String() {
	case "debug":
		//log.SetOutput(os.Stdout)
		setNull()
	case "info":
		setNull()
	case "warn":
		setNull()
	case "error":
		setNull()
	default:
		setNull()
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: debugWriter, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.0000Z07:00",
	})
	log.AddHook(lfHook)

	return nil
}

func setNull() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	writer := bufio.NewWriter(src)
	log.SetOutput(writer)
}
