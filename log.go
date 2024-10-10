package simplelog

import (
	"fmt"
	"io"
	"os"
	"time"
)

type LogLevel int

var (
	LogCache     []string //nolint:gochecknoglobals // wontfix
	LogCacheSize = 256    //nolint:gochecknoglobals // wontfix
)

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	GoRoutineErrorLevel
)

type Logger struct {
	MinLevel LogLevel
}

func NewLogger(minLevel LogLevel) *Logger {
	return &Logger{MinLevel: minLevel}
}

func (l *Logger) log(level LogLevel, message string, writer io.Writer) {
	if Stdout == nil {
		fmt.Fprintln(os.Stdout, "Logger STDOUT is nil")
		return
	}

	color := "\033[0m"
	errColor := "\033[0;31m"
	levelName := "DEBUG"

	switch level {
	case DebugLevel:
		levelName = "DEBUG"
		color = "\033[0;34m"
	case InfoLevel:
		levelName = "INFO"
		color = "\033[0;32m"
	case WarnLevel:
		levelName = "WARN"
		color = "\033[0;33m"
	case ErrorLevel:
		levelName = "ERROR"
		color = errColor
	case FatalLevel:
		levelName = "FATAL"
		color = errColor
	case GoRoutineErrorLevel:
		levelName = "GOROUTINE_ERROR"
		color = errColor
	}

	if level >= l.MinLevel {
		oColor := fmt.Sprintf("[%s] %s[%s]%s %s\n", time.Now().Format("2006-01-02 15:04:05"), color, levelName, "\033[0m", message)
		oRaw := fmt.Sprintf("[%s] [%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), levelName, message)

		if len(LogCache) > LogCacheSize {
			ClearCache()
		}

		LogCache = append(LogCache, oRaw)

		fmt.Fprint(writer, oColor)
	}
}

func (l *Logger) Debug(format string) {
	l.log(DebugLevel, format, Stdout)
}

func (l *Logger) Debugf(format string, a ...any) {
	l.log(DebugLevel, fmt.Sprintf(format, a...), Stdout)
}

func (l *Logger) Info(format string) {
	l.log(InfoLevel, format, Stdout)
}

func (l *Logger) Infof(format string, a ...any) {
	l.log(InfoLevel, fmt.Sprintf(format, a...), Stdout)
}

func (l *Logger) Warn(format string) {
	l.log(WarnLevel, format, Stdout)
}

func (l *Logger) Warnf(format string, a ...any) {
	l.log(WarnLevel, fmt.Sprintf(format, a...), Stdout)
}

func (l *Logger) Error(format string) {
	l.log(ErrorLevel, format, Stdout)
}

func (l *Logger) Errorf(format string, a ...any) {
	l.log(ErrorLevel, fmt.Sprintf(format, a...), Stdout)
}

func (l *Logger) Fatal(format string) {
	l.log(FatalLevel, format, Stdout)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, a ...any) {
	l.log(FatalLevel, fmt.Sprintf(format, a...), Stdout)
	os.Exit(1)
}

func (l *Logger) NewLine() {
	fmt.Fprint(Stdout, "\n")
}

func ClearCache() {
	LogCache = []string{}
}

var (
	Stdout       io.Writer = os.Stdout            //nolint:gochecknoglobals // wontfix
	SharedLogger           = NewLogger(InfoLevel) //nolint:gochecknoglobals // wontfix
)
