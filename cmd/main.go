package main

import (
	"github.com/sirupsen/logrus"
	"indev-engine/window"
	"os"
	"os/signal"
)

var logger = logrus.StandardLogger()

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func() {
		err := window.CreateWindow(300, 300)
		if err != nil {
			logger.Error(err)
			return
		}
		stop <- nil
	}()

	<-stop
}
