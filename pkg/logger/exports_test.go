package logger

import (
	"bufio"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type testLogger interface {
	Logger

	getLabel() string
	getSubLabel() string
	setZapSugaredLogger(*zap.SugaredLogger)
}

func (l *ZapLogger) getLabel() string {
	return l.label
}

func (l *ZapLogger) getSubLabel() string {
	return l.subLabel
}

func (l *ZapLogger) setZapSugaredLogger(zsl *zap.SugaredLogger) {
	l.zapSugaredLogger = zsl
}

type logOutputCapturer struct {
	buf    strings.Builder
	writer *bufio.Writer
}

func newLogOutputCapturer(logger testLogger) *logOutputCapturer {
	c := logOutputCapturer{}
	c.buf = strings.Builder{}
	c.writer = bufio.NewWriter(&c.buf)

	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	newZapSugaredLogger := zap.New(
		zapcore.NewCore(encoder, zapcore.AddSync(c.writer), zapcore.DebugLevel)).
		Sugar()

	logger.setZapSugaredLogger(newZapSugaredLogger)

	return &c
}

func (c *logOutputCapturer) get() string {
	c.writer.Flush()
	return c.buf.String()
}
