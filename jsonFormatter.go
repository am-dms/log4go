package log4go

import "github.com/sirupsen/logrus"

type dmsJsonFormatter struct {
	*logrus.JSONFormatter
	category string
}

func getJsonFormater(cat string) *dmsJsonFormatter {
	return &dmsJsonFormatter{
		JSONFormatter: &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime: "ts",
			},
		},
		category: cat,
	}
}

func (dmf *dmsJsonFormatter) Format(ent *logrus.Entry) ([]byte, error) {
	ent.Data["category"] = dmf.category
	return dmf.JSONFormatter.Format(ent)
}
