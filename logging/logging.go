/*
 * Author: fasion
 * Created time: 2019-05-13 09:40:16
 * Last Modified by: fasion
 * Last Modified time: 2019-12-11 14:52:24
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

type KeeperLoggerContext struct {
	loggerLevel zap.AtomicLevel
	logger      *zap.Logger
	writeSyncer *DynamicWriteSyncer
}

func CreateKeeperLoggerContextWithStdout() *KeeperLoggerContext {
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

	return &KeeperLoggerContext{
		loggerLevel: loggerLevel,
		logger:      logger,
		writeSyncer: writeSyncer,
	}
}

func (klc *KeeperLoggerContext) GetLogger() *zap.Logger {
	return klc.logger
}

func (klc *KeeperLoggerContext) SetLoggerLevel(level zapcore.Level) {
	klc.loggerLevel.SetLevel(level)
}

func (klc *KeeperLoggerContext) UseCustomWriteSyncer(ws zapcore.WriteSyncer, closer CloserFunc) error {
	return klc.writeSyncer.Update(ws, closer)
}

func (klc *KeeperLoggerContext) UseCustomWriter(w io.Writer) error {
	return klc.UseCustomWriteSyncer(zapcore.AddSync(w), nil)
}

func (klc *KeeperLoggerContext) UseCustomWriteCloser(wc io.WriteCloser) error {
	return klc.UseCustomWriteSyncer(zapcore.AddSync(wc), wc.Close)
}

func (klc *KeeperLoggerContext) UseCustomFileWriteSyncer(path string, maxSize, maxAge, maxBackups int, localTime, compress bool) error {
	// create parent directories
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	return klc.UseCustomWriteCloser(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		LocalTime:  localTime,
		Compress:   compress,
	})
}

func (klc *KeeperLoggerContext) UseFileWriteSyncer(path string) error {
	return klc.UseCustomFileWriteSyncer(
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

var keeperLoggerContext = CreateKeeperLoggerContextWithStdout()

var GetLogger = keeperLoggerContext.GetLogger
var SetLoggerLevel = keeperLoggerContext.SetLoggerLevel
