package main

import (
	log "github.com/am-dms/log4go"
)

type A struct {
	Name string
	Age  int
}

type B struct {
	Nest    A
	NestPtr *A
	Kind    string
	Arr     []int
	ArrStr  []C
}

type C struct {
	Val string
}

func apiLogs() {
	logger := log.NewLogger("API")

	// log := logrus.New()
	logent := logger.WithFields(log.Fields{
		"users": 12,
		"name":  "mahdi",
	})
	logent.Info("hi mahdi")
	logent.Data["users"] = 14
	logent.Debug("this is debug")
	logent.Error("system got screwed")

	a := A{
		Name: "mahdi",
		Age:  33,
	}

	b := B{
		Nest: A{
			Name: "innermahdi",
			Age:  10,
		},
		NestPtr: &A{
			Name: "pointToMahdi",
			Age:  1,
		},
		Kind:   "developer",
		Arr:    []int{2, 3, 4},
		ArrStr: make([]C, 0),
	}
	b.ArrStr = append(b.ArrStr, C{Val: "JustAValue"})

	logger.WithObject(a).Warn("flat object")
	logger.WithObject(b).Info("nested object")
}

func frontLogs() {
	logger := log.NewLogger("frontEnd")

	// log := logrus.New()
	user1 := A{Age: 20, Name: "user1"}
	logger.WithObjectAndFileds(user1, log.Fields{
		"credit": 1200.23,
		"Gender": "male",
	}).Info("we have a log here")

}

func main() {
	apiLogs()
	frontLogs()
}
