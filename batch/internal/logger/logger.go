package logger

import (
	"encoding/json"
	"fmt"

	"github.com/htsuchinga/golang-localstack/batch/internal/datetime"
)

type Level string

const (
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
)

type fields struct {
	DateTime   string `json:"dateTime"`
	ModuleName string `json:"moduleName"`
	Level      Level  `json:"level"`
	Message    string `json:"message"`
}

type Log struct {
	f fields
}

func (l *Log) ModuleName(moduleName string) *Log {
	l.f.ModuleName = moduleName
	return l
}

func (l *Log) output(level Level, format interface{}, args ...interface{}) {
	l.f.DateTime = timeNow().Format("2006-01-02T15:04:05.000Z07:00")
	l.f.Level = level
	l.f.Message = formatMessage(format, args...)
	buf, _ := json.Marshal(&l.f)
	line := string(buf)
	_, _ = fmtPrintln(line)
}

func formatMessage(format interface{}, args ...interface{}) string {
	switch format.(type) {
	case string:
		formatString := fmt.Sprintf("%v", format)
		return fmt.Sprintf(formatString, args...)
	default:
		if len(args) == 0 {
			return fmt.Sprint(format)
		}
		return fmt.Sprintf("%v %s", format, fmt.Sprint(args...))
	}
}

var timeNow = datetime.NowInJST
var fmtPrintln = fmt.Println

func (l *Log) Info(format interface{}, args ...interface{}) {
	l.output(LevelInfo, format, args...)
}

func (l *Log) Warn(format interface{}, args ...interface{}) {
	l.output(LevelWarn, format, args...)
}

func (l *Log) Error(format interface{}, args ...interface{}) {
	l.output(LevelError, format, args...)
}

var DefaultModuleName string = ""

func newLog() *Log {
	return &Log{
		f: fields{
			ModuleName: DefaultModuleName,
		},
	}
}

func ModuleName(moduleName string) *Log {
	return newLog().ModuleName(moduleName)
}

func Info(format interface{}, args ...interface{}) {
	newLog().Info(format, args...)
}

func Warn(format interface{}, args ...interface{}) {
	newLog().Warn(format, args...)
}

func Error(format interface{}, args ...interface{}) {
	newLog().Error(format, args...)
}
