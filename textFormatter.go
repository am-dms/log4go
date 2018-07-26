package log4go

import "github.com/sirupsen/logrus"

type dmsTextFormatter struct {
	*logrus.TextFormatter
	category string
}

func getTextFormater(cat string) *dmsTextFormatter {
	return &dmsTextFormatter{
		TextFormatter: &logrus.TextFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime: "ts",
			},
		},
		category: cat,
	}
}

func (dmf *dmsTextFormatter) Format(ent *logrus.Entry) ([]byte, error) {
	ent.Data["category"] = dmf.category
	return dmf.TextFormatter.Format(ent)
}
