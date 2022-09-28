package demo

import "github.com/sirupsen/logrus"

var log *logrus.Logger

func UseLog(l *logrus.Logger) {
	log = l
}
