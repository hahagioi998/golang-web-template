package logger

import (
	"time"

	"go.uber.org/zap"

	"base/logger/zaplog"
)

var myLogger *ZapLogger
var levelMapping = levelMap()

// 初始化日志框架
func InitLogger(projectName, logPath string, maxAge, rotationTime time.Duration, level string) {
	zaplog.InitZap(projectName, logPath, maxAge, rotationTime)
	// 设置打印堆栈深度，设置日志级别
	myLogger = &ZapLogger{
		slog:  zap.L().WithOptions(zap.AddCallerSkip(2)).Sugar(),
		level: levelMapping[level],
	}
}

// 字符串和级别映射
func levelMap() map[string]LogLevel {
	return map[string]LogLevel{
		"debug": logDebug,
		"info":  logInfo,
		"warn":  logWarn,
		"err":   logErr,
		"off":   logOff,
	}
}

// 外部直接调用方法就可以，简化调用方式
func Debug(v ...interface{}) {
	myLogger.Debug(v...)
}
func Debugf(format string, v ...interface{}) {
	myLogger.Debugf(format, v...)
}
func Info(v ...interface{}) {
	myLogger.Info(v...)
}
func Infof(format string, v ...interface{}) {
	myLogger.Infof(format, v...)
}
func Warn(v ...interface{}) {
	myLogger.Warn(v...)
}
func Warnf(format string, v ...interface{}) {
	myLogger.Warnf(format, v...)
}
func Error(v ...interface{}) {
	myLogger.Error(v...)
}
func Errorf(format string, v ...interface{}) {
	myLogger.Errorf(format, v...)
}

// 具体对内实现
type ZapLogger struct {
	slog  *zap.SugaredLogger
	level LogLevel
}

func (z *ZapLogger) Debug(v ...interface{}) {
	if z.Level() <= logDebug {
		z.slog.Debug(v...)
	}
}

func (z *ZapLogger) Debugf(format string, v ...interface{}) {
	if z.Level() <= logDebug {
		z.slog.Debugf(format, v...)
	}
}

func (z *ZapLogger) Info(v ...interface{}) {
	if z.Level() <= logInfo {
		z.slog.Info(v...)
	}
}

func (z *ZapLogger) Infof(format string, v ...interface{}) {
	if z.Level() <= logInfo {
		z.slog.Infof(format, v...)
	}
}

func (z *ZapLogger) Warn(v ...interface{}) {
	if z.Level() <= logWarn {
		z.slog.Warn(v...)
	}
}

func (z *ZapLogger) Warnf(format string, v ...interface{}) {
	if z.Level() <= logWarn {
		z.slog.Warnf(format, v...)
	}
}

func (z *ZapLogger) Error(v ...interface{}) {
	if z.Level() <= logErr {
		z.slog.Error(v...)
	}
}

func (z *ZapLogger) Errorf(format string, v ...interface{}) {
	if z.Level() <= logErr {
		z.slog.Errorf(format, v)
	}
}

func (z *ZapLogger) Level() LogLevel {
	return z.level
}

func (z *ZapLogger) SetLevel(l LogLevel) {
	z.level = l
}