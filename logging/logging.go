/*
 * Author: fasion
 * Created time: 2019-05-13 09:40:16
 * Last Modified by: fasion
 * Last Modified time: 2020-04-07 13:56:34
 */

package logging

import (
	"io"
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type CloserFunc func() error

type DynamicWriteSyncer struct {
	zapcore.WriteSyncer
	closer CloserFunc
	mutex  sync.Mutex
}

func (dws *DynamicWriteSyncer) Update(ws zapcore.WriteSyncer, closer CloserFunc) error {
	dws.mutex.Lock()
	defer dws.mutex.Unlock()

	if dws.closer != nil {
		if err := dws.closer(); err != nil {
			return err
		}
	}

	dws.WriteSyncer = ws
	dws.closer = closer

	return nil
}

var encoderConfig = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "level",
	NameKey:        "name",
	MessageKey:     "message",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

type LoggerContainer struct {
	loggerLevel zap.AtomicLevel
	logger      *zap.Logger
	writeSyncer *DynamicWriteSyncer
}

func NewLoggerContainerWithStdout() *LoggerContainer {
	// default log level
	loggerLevel := zap.NewAtomicLevelAt(zapcore.InfoLevel)

	// json format encoder
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// write syncer with stdout
	stdoutWriteSyncer := zapcore.AddSync(os.Stdout)

	// dynamic write syncer
	writeSyncer := &DynamicWriteSyncer{
		WriteSyncer: stdoutWriteSyncer,
	}

	// logging core
	core := zapcore.NewCore(jsonEncoder, writeSyncer, loggerLevel)

	// logger
	logger := zap.New(core)

	return &LoggerContainer{
		loggerLevel: loggerLevel,
		logger:      logger,
		writeSyncer: writeSyncer,
	}
}

func (container *LoggerContainer) GetLogger() *zap.Logger {
	return container.logger
}

func (container *LoggerContainer) SetLoggerLevel(level zapcore.Level) {
	container.loggerLevel.SetLevel(level)
}

func (container *LoggerContainer) UseCustomWriteSyncer(ws zapcore.WriteSyncer, closer CloserFunc) error {
	return container.writeSyncer.Update(ws, closer)
}

func (container *LoggerContainer) UseCustomWriter(w io.Writer) error {
	return container.UseCustomWriteSyncer(zapcore.AddSync(w), nil)
}

func (container *LoggerContainer) UseCustomWriteCloser(wc io.WriteCloser) error {
	return container.UseCustomWriteSyncer(zapcore.AddSync(wc), wc.Close)
}

func (container *LoggerContainer) UseCustomFileWriteSyncer(path string, maxSize, maxAge, maxBackups int, localTime, compress bool) error {
	// create parent directories
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	return container.UseCustomWriteCloser(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		LocalTime:  localTime,
		Compress:   compress,
	})
}

func (container *LoggerContainer) UseFileWriteSyncer(path string) error {
	return container.UseCustomFileWriteSyncer(
		// path
		path,
		// max size in megabytes
		20,
		// max age in days
		7,
		// max backups
		5,
		// use local time
		true,
		// compress
		true,
	)
}

var loggerContainer = NewLoggerContainerWithStdout()

var GetLogger = loggerContainer.GetLogger
var SetLoggerLevel = loggerContainer.SetLoggerLevel
