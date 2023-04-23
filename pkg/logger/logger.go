package logger

import (
	"log"

	"github.com/kevindoubleu/pichan/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	SetSubLabel(subLabel string)
	SetConfig(config configs.Log)

	Infow(msg string, keysAndValues ...interface{})

	Error(msg string, err error)
	ErrorMsg(msg ...interface{})
	Errorw(msg string, keysAndValues ...interface{})

	FatalError(msg string, err error)
	Fatalw(msg string, keysAndValues ...interface{})

	GetZapLogger() *zap.Logger
}

type ZapLogger struct {
	zapSugaredLogger *zap.SugaredLogger

	// identifies general label such as
	// appname
	// appname.packagename
	// appname.packagename.structname
	label string

	// identifies specific label such as
	// functionname
	subLabel string

	config configs.Log
}

func NewLogger(label string) Logger {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatal("Logger.NewLogger: cannot build logger", err)
	}
	defer logger.Sync()

	return &ZapLogger{
		zapSugaredLogger: logger.Sugar(),
		label:            label,
		subLabel:         "",
	}
}

func (l *ZapLogger) GetZapLogger() *zap.Logger {
	return l.zapSugaredLogger.Desugar()
}

func (l *ZapLogger) SetSubLabel(subLabel string) {
	l.subLabel = subLabel
}

func (l *ZapLogger) SetConfig(config configs.Log) {
	l.config = config
}

func (l *ZapLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.zapSugaredLogger.Infow(l.prefix(msg),
		keysAndValues...)
}

func (l *ZapLogger) ErrorMsg(msg ...interface{}) {
	completeMsg := make([]interface{}, 0)
	completeMsg = append(completeMsg, l.getPrefix()+": ")
	completeMsg = append(completeMsg, msg...)
	l.zapSugaredLogger.Error(completeMsg...)
}

func (l *ZapLogger) Error(msg string, err error) {
	l.zapSugaredLogger.Errorw(l.prefix(msg),
		"err", err,
	)
}

func (l *ZapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.zapSugaredLogger.Errorw(l.prefix(msg), keysAndValues...)
}

func (l *ZapLogger) FatalError(msg string, err error) {
	if l.config.SkipFatal {
		l.Error(l.prefix("skipped fatal: "+msg), err)
		return
	}

	l.zapSugaredLogger.Fatalw(l.prefix(msg),
		"err", err,
	)
}

func (l *ZapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	if l.config.SkipFatal {
		l.Errorw(l.prefix("skipped fatal: "+msg), keysAndValues...)
		return
	}

	l.zapSugaredLogger.Fatalw(l.prefix(msg), keysAndValues...)
}

func (l *ZapLogger) prefix(msg string) string {
	return l.label + "." + l.subLabel + ": " + msg
}

func (l *ZapLogger) getPrefix() string {
	return l.label + "." + l.subLabel
}
