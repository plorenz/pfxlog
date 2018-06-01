package main

import (
	"git.quigley.com/products/pfxlog"
	"github.com/sirupsen/logrus"
	"fmt"
)

func init() {
	pfxlog.Global(logrus.DebugLevel)
}

func main() {
	log := pfxlog.Logger()
	log.Info("hello world.")

	notifications := make(chan int)
	for i := 0; i < 5; i++ {
		go counter(i, notifications)
	}

	for i := 0; i < 5; i++ {
		n := <- notifications
		log.Infof("%d done.", n)
	}

	log.Info("complete.")
}

func counter(number int, notify chan int) {
	log := pfxlog.AttachedLogger(fmt.Sprintf("counter #%d", number))

	for i := 0; i < 5; i++ {
		log.Infof("visited %d.", i)
	}

	log.Info("complete")

	notify <- number
}
