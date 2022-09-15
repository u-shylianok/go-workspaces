package library

import (
	"github.com/sirupsen/logrus"
)

func Greetings(name string) {
	logrus.Infof("Hi, %s!\n", name)
}

func Bye(name string) {
	logrus.Infof("Bye, %s!\n", name)
}
