package main

import (
	"github.com/sirupsen/logrus"
	"indev-engine/window"
	"os"
	"os/signal"
	"time"
)

var logger = logrus.StandardLogger()

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	win := window.NewWindow(300, 300)
	err := win.Create()
	if err != nil {
		logger.Error(err)
		return
	}

	<-win.Finished
	time.Sleep(time.Second)
	close(stop)

	<-stop
}
