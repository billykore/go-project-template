package log

import (
	"fmt"
	"os"
	"sync"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

var (
	_once sync.Once
	_log  *Logger
)

// Logger provides logging functions.
type Logger struct {
	logger  *zap.Logger
	usecase string
}

// New return the singleton instance of Logger.
func New() *Logger {
	_once.Do(func() {
		logger := newZap()
		_log = &Logger{logger: logger}
	})
	return _log
}

func newZap() *zap.Logger {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.InfoLevel)
	logger := zap.New(core, zap.AddStacktrace(zap.PanicLevel))
	return logger
}

// Logger returns the underlying zap.Logger instance for the Logger.
func (l *Logger) Logger() *zap.Logger {
	return l.logger
}

func (l *Logger) Usecase(usecase string) *Logger {
	l.usecase = usecase
	return l
}

// Errorf log.
func (l *Logger) Errorf(format string, args ...any) {
	l.logger.Error(
		fmt.Sprintf(format, args...),
		zap.String("usecase", l.usecase),
	)
}

// Fatalf logs an error, then shutdown the domain.
func (l *Logger) Fatalf(format string, args ...any) {
	l.logger.Fatal(
		fmt.Sprintf(format, args...),
		zap.String("usecase", l.usecase),
	)
}

// Infof log.
func (l *Logger) Infof(format string, args ...any) {
	l.logger.Info(
		fmt.Sprintf(format, args...),
		zap.String("usecase", l.usecase),
	)
}
