package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// ZapLog is global logger
	ZapLog *zap.Logger
)

type zapLogger struct {
	zapLogger *zap.Logger
}

func getZapLevel(level Level) zapcore.Level {
	switch level {
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case DebugLevel:
		return zapcore.DebugLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func getEncoder(isJSON bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func newZapLogger(conf Configuration) Logger {
	var cores []zapcore.Core

	if conf.EnableConsole {
		level := getZapLevel(conf.ConsoleLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(conf.ConsoleJSONFormat), writer, level)
		cores = append(cores, core)
	}

	if conf.EnableFile {
		level := getZapLevel(conf.FileLevel)
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename: conf.FileLocation,
			MaxSize:  100,
			Compress: true,
			MaxAge:   28,
		})

		core := zapcore.NewCore(
			getEncoder(conf.FileJSONFormat),
			writer,
			level)

		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zap.go
	ZapLog = zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	)

	zap.RedirectStdLog(ZapLog)
	return &zapLogger{zapLogger: ZapLog}
}

func (l *zapLogger) Debug(format string, args ...interface{}) {
	if len(args) > 0 {
		l.zapLogger.Debug(fmt.Sprintf(format, args...))
	} else {
		l.zapLogger.Debug(format)
	}
}

func (l *zapLogger) Info(format string, args ...interface{}) {
	if len(args) > 0 {
		l.zapLogger.Info(fmt.Sprintf(format, args...))
	} else {
		l.zapLogger.Info(format)
	}
}

func (l *zapLogger) Warn(format string, args ...interface{}) {
	if len(args) > 0 {
		l.zapLogger.Warn(fmt.Sprintf(format, args...))
	} else {
		l.zapLogger.Warn(format)
	}
}

func (l *zapLogger) Error(format string, args ...interface{}) {
	if len(args) > 0 {
		l.zapLogger.Error(fmt.Sprintf(format, args...))
	} else {
		l.zapLogger.Error(format)
	}
}

func (l *zapLogger) Fatal(format string, args ...interface{}) {
	if len(args) > 0 {
		l.zapLogger.Fatal(fmt.Sprintf(format, args...))
	} else {
		l.zapLogger.Fatal(format)
	}
}

func (l *zapLogger) Panic(format string, args ...interface{}) {
	if len(args) > 0 {
		l.zapLogger.Panic(fmt.Sprintf(format, args...))
	} else {
		l.zapLogger.Panic(format)
	}
}
