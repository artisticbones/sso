package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func init() {
	Default()
}

const (
	defaultLog = "sso.log"
	maxSize    = 500 //megabytes
	maxBackups = 3
	maxAge     = 7
)

type Level = zapcore.Level

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in production.
	DebugLevel Level = zapcore.DebugLevel
	// InfoLevel is the default logging priority.
	InfoLevel Level = zapcore.InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual human review.
	WarnLevel Level = zapcore.WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly, it shouldn't generate any error-level logs.
	ErrorLevel Level = zapcore.ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the logger panics after writing the message.
	DPanicLevel Level = zapcore.DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel Level = zapcore.PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1)
	FatalLevel Level = zapcore.FatalLevel
)

type Field = zapcore.Field
type Option = zap.Option

type Logger struct {
	log        *zap.Logger
	filename   string
	maxSize    int
	maxAge     int
	maxBackups int
	compress   bool
}

var _lg *Logger

// GetLogger obtain the global logger
func GetLogger() *Logger { return _lg }

// Default function initialize a default instance
func Default() {
	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   defaultLog,
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge, // days
	})
	std := zapcore.AddSync(os.Stdout)

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), file, InfoLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()), std, InfoLevel),
	)

	_lg = &Logger{
		log:        zap.New(core, nil),
		filename:   defaultLog,
		maxSize:    maxSize,
		maxAge:     maxAge,
		maxBackups: maxBackups,
		compress:   true,
	}
}

func New(filename string, level Level, opts ...Option) *Logger {
	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge, // days
	})
	std := zapcore.AddSync(os.Stdout)

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), file, level),
		zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()), std, level),
	)

	return &Logger{
		log:        zap.New(core, opts...),
		filename:   filename,
		maxSize:    maxSize,
		maxAge:     maxAge,
		maxBackups: maxBackups,
		compress:   true,
	}
}

func (lg *Logger) Sync() error {
	return lg.log.Sync()
}

func (lg *Logger) Debug(msg string, fields ...Field) {
	lg.log.Debug(msg, fields...)
}

func (lg *Logger) Debugf(template string, args ...interface{}) {
	lg.log.Sugar().Debugf(template, args...)
}

func (lg *Logger) Info(msg string, fields ...Field) {
	lg.log.Info(msg, fields...)
}

func (lg *Logger) Infof(template string, args ...interface{}) {
	lg.log.Sugar().Debugf(template, args...)
}

func (lg *Logger) Error(msg string, fields ...Field) {
	lg.log.Error(msg, fields...)
}

func (lg *Logger) Errorf(template string, args ...interface{}) {
	lg.log.Sugar().Errorf(template, args...)
}

func (lg *Logger) DPanic(msg string, fields ...Field) {
	lg.log.DPanic(msg, fields...)
}

func (lg *Logger) DPanicf(template string, args ...interface{}) {
	lg.log.Sugar().DPanicf(template, args...)
}

func (lg *Logger) Panic(msg string, fields ...Field) {
	lg.log.Panic(msg, fields...)
}

func (lg *Logger) Panicf(template string, args ...interface{}) {
	lg.log.Sugar().Panicf(template, args...)
}

func (lg *Logger) Fatal(msg string, fields ...Field) {
	lg.log.Fatal(msg, fields...)
}

func (lg *Logger) Fatalf(template string, args ...interface{}) {
	lg.log.Sugar().Fatalf(template, args...)
}

func Info(msg string, fields ...Field) {
	_lg.Info(msg, fields...)
}

func Infof(template string, args ...interface{}) {
	_lg.Infof(template, args...)
}

func Debug(msg string, fields ...Field) {
	_lg.Debug(msg, fields...)
}

func Debugf(template string, args ...interface{}) {
	_lg.Debugf(template, args...)
}

func Error(msg string, fields ...Field) {
	_lg.Error(msg, fields...)
}

func Errorf(template string, args ...interface{}) {
	_lg.Errorf(template, args...)
}

func Fatal(msg string, fields ...Field) {
	_lg.Fatal(msg, fields...)
}

func Fatalf(template string, args ...interface{}) {
	_lg.Fatalf(template, args...)
}
