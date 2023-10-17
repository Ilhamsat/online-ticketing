package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"online-ticketing/app/config"
	"online-ticketing/app/global"
	"online-ticketing/infrastructures/api"
)

func main() {
	logrus.Infoln(fmt.Sprintf("Application Version : %s\n", global.BuildVersion))
	config.Init()

	forever := make(chan int)
	go func() {
		go api.Serve()
	}()
	<-forever
}
