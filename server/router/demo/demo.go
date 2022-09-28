package demo

import "github.com/sirupsen/logrus"

var log *logrus.Logger

func Log(l *logrus.Logger) {
	log = l
}
