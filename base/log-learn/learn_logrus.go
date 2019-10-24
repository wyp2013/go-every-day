package main

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

type LogConfig struct {
	Level      string `yaml:"level"`      // panic, fatal, error, warn, info, debug
	ShowSource bool   `yaml:"showSource"` // show function call depth in logging or not
	Path       string `yaml:"path"`
}

func initLogger(cfg LogConfig) {
	level, err := log.ParseLevel(cfg.Level)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(level)

	formatter := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.0000Z07:00",
	}

	log.SetFormatter(formatter)

	if cfg.ShowSource {
		filenameHook := filename.NewHook()
		filenameHook.Field = "source"
		log.AddHook(filenameHook)
	}

	if cfg.Path == "" {
		log.SetOutput(os.Stdout)
		return
	}

	// init logger with log path
	baseLogPath := cfg.Path
	debugriter, err := rotatelogs.New(
		baseLogPath+"debug.log.%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath+"debug.log"),
		rotatelogs.WithMaxAge(1 * time.Minute),
		rotatelogs.WithRotationTime(1 * time.Minute),
	)

	writer, err := rotatelogs.New(
		baseLogPath+"info.log.%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath+"info.log"),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	if err != nil {
		log.Errorf("config local file system logger error. %v", errors.WithStack(err))
	}

	// log level settings, if debug level, output to console.
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
		log.DebugLevel: debugriter, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.0000Z07:00",
	})

	log.AddHook(lfHook)
}

func setNull() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	writer := bufio.NewWriter(src)
	log.SetOutput(writer)
}

func main() {
	cfg := LogConfig{
		Level: "debug",
		ShowSource: true,
		Path:"/Users/bitmain/tmp/",
	}

	initLogger(cfg)

	for i := 0; i < 10000; i++ {
		log.Debug("test ", i)
	}

	for i := 0; i < 10000; i++ {
		log.Info("test ", i)
	}

	for i := 0; i < 10000; i++ {
		log.Error("test ", i)
	}

	fmt.Println("xxx")

}
