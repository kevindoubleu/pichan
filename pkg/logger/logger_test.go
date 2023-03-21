package logger

import (
	"errors"
	"fmt"
	"testing"

	"github.com/kevindoubleu/pichan/configs"
	"github.com/stretchr/testify/assert"
)

var (
	testLabel    = "main label"
	testSubLabel = "sub label"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger(testLabel).(testLogger)

	assert.Equal(t, testLabel, logger.getLabel())
	assert.NotNil(t, logger.GetZapLogger())
}

func TestSetSublabel(t *testing.T) {
	testLabel := "main label"
	logger := NewLogger(testLabel).(testLogger)
	logger.SetSubLabel(testSubLabel)

	assert.Equal(t, testSubLabel, logger.getSubLabel())
}

func TestInfow(t *testing.T) {
	logger := NewLogger(testLabel).(testLogger)
	logger.SetSubLabel(testSubLabel)
	capturer := newLogOutputCapturer(logger)

	logger.Infow("info message", "key", "val")

	captured := capturer.get()
	expectedLabel := testLabel + "." + testSubLabel + ": "
	expectedMsg := "info message"
	expectedArgs := `{"key": "val"}`
	assert.Contains(t, captured, "INFO")
	assert.Contains(t, captured, expectedLabel)
	assert.Contains(t, captured, expectedMsg)
	assert.Contains(t, captured, expectedArgs)
}

func TestError(t *testing.T) {
	logger := NewLogger(testLabel).(testLogger)
	logger.SetSubLabel(testSubLabel)
	capturer := newLogOutputCapturer(logger)

	logger.Error("error message", errors.New("example error"))

	captured := capturer.get()
	expectedLabel := testLabel + "." + testSubLabel + ": "
	expectedMsg := "error message"
	expectedArgs := `{"err": "example error"}`
	assert.Contains(t, captured, "ERROR")
	assert.Contains(t, captured, expectedLabel)
	assert.Contains(t, captured, expectedMsg)
	assert.Contains(t, captured, expectedArgs)
}

func TestErrorMsg(t *testing.T) {
	logger := NewLogger(testLabel).(testLogger)
	logger.SetSubLabel(testSubLabel)
	capturer := newLogOutputCapturer(logger)

	errMsgs := []string{"error message", "msg2", "msg3"}
	logger.ErrorMsg(errMsgs)

	captured := capturer.get()
	expectedLabel := testLabel + "." + testSubLabel + ": "
	assert.Contains(t, captured, "ERROR")
	assert.Contains(t, captured, expectedLabel)
	assert.Contains(t, captured, fmt.Sprint(errMsgs))
}

func TestErrorw(t *testing.T) {
	logger := NewLogger(testLabel).(testLogger)
	logger.SetSubLabel(testSubLabel)
	capturer := newLogOutputCapturer(logger)

	logger.Errorw("error message", "key", "val", "key2", "val2")

	captured := capturer.get()
	expectedLabel := testLabel + "." + testSubLabel + ": "
	expectedMsg := "error message"
	expectedArgs := `{"key": "val", "key2": "val2"}`
	assert.Contains(t, captured, "ERROR")
	assert.Contains(t, captured, expectedLabel)
	assert.Contains(t, captured, expectedMsg)
	assert.Contains(t, captured, expectedArgs)
}

func TestFatalError(t *testing.T) {
	configs.LOG_SKIP_FATAL = true
	logger := NewLogger(testLabel).(testLogger)
	logger.SetSubLabel(testSubLabel)
	capturer := newLogOutputCapturer(logger)

	logger.FatalError("fatal message", errors.New("example fatal error"))

	captured := capturer.get()
	expectedLabel := testLabel + "." + testSubLabel + ": "
	expectedMsg := "fatal message"
	expectedArgs := `{"err": "example fatal error"}`
	assert.Contains(t, captured, "skipped fatal")
	assert.Contains(t, captured, expectedLabel)
	assert.Contains(t, captured, expectedMsg)
	assert.Contains(t, captured, expectedArgs)
}

func TestFatalw(t *testing.T) {
	configs.LOG_SKIP_FATAL = true
	logger := NewLogger(testLabel).(testLogger)
	logger.SetSubLabel(testSubLabel)
	capturer := newLogOutputCapturer(logger)

	logger.Fatalw("fatal message", "key", "val", "key2", "val2")

	captured := capturer.get()
	expectedLabel := testLabel + "." + testSubLabel + ": "
	expectedMsg := "fatal message"
	expectedArgs := `{"key": "val", "key2": "val2"}`
	assert.Contains(t, captured, "skipped fatal")
	assert.Contains(t, captured, expectedLabel)
	assert.Contains(t, captured, expectedMsg)
	assert.Contains(t, captured, expectedArgs)
}
