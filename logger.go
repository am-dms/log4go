package log4go

import (
	"os"

	"github.com/sirupsen/logrus"
)

//Logger is logrus logger + a category to identify and filter logs with
type Logger struct {
	*logrus.Logger
	category string
}

//Fields just to make it easy to use loger object
type Fields = logrus.Fields

//NewLogger creates new instance of Logger and configures it
func NewLogger(cat string) *Logger {
	lg := &Logger{
		Logger:   logrus.New(),
		category: cat,
	}
	if os.Getenv("LOG_FMT") == "text" {
		lg.SetTextFormatter()
	} else {
		lg.SetJSONFormatter()
	}
	lg.setlevel()
	return lg
}

//WithObject adds a sugar
func (lg *Logger) WithObject(obj interface{}) *logrus.Entry {
	return lg.WithFields(getFieldMap(obj))
}

//WithObjectAndFileds adds a sugar
func (lg *Logger) WithObjectAndFileds(obj interface{}, fl Fields) *logrus.Entry {
	ent := lg.WithFields(getFieldMap(obj))
	for k, v := range fl {
		ent.Data[k] = v
	}
	return ent

}

//SetJSONFormatter changes the output format to JSON
func (lg *Logger) SetJSONFormatter() {
	lg.Formatter = getJsonFormater(lg.category)
}

//SetTextFormatter changes the output format to Text
func (lg *Logger) SetTextFormatter() {
	lg.Formatter = getTextFormater(lg.category)
}

//SetJSONFormatter changes the output format to JSON
func (lg *Logger) setlevel() {
	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		lg.SetLevel(logrus.DebugLevel)
	case "INFO":
		lg.SetLevel(logrus.InfoLevel)
	case "WARN":
		lg.SetLevel(logrus.WarnLevel)
	case "ERROR":
		lg.SetLevel(logrus.ErrorLevel)
	default:
		lg.SetLevel(logrus.InfoLevel)

	}
}
