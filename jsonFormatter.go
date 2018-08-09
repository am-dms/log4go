package log4go

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type dmsJsonFormatter struct {
	*logrus.JSONFormatter
	category string
}

func getJsonFormater(cat string) *dmsJsonFormatter {
	return &dmsJsonFormatter{
		JSONFormatter: &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "ts",
				logrus.FieldKeyLevel: "_level",
			},
		},
		category: cat,
	}
}

func (dmf *dmsJsonFormatter) Format(ent *logrus.Entry) ([]byte, error) {
	ent.Data["category"] = dmf.category
	ent.Data["level"] = strings.ToUpper(ent.Level.String())
	return dmf.JSONFormatter.Format(ent)
}
