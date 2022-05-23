package main

import (
	app "github.com/duxianghua/gin-standard-template/internal/api"
	_ "github.com/duxianghua/gin-standard-template/pkg/log"
)

// Version comes from build time
var Version string

func main() {
	// log.WithFields(log.Fields{"app": "golang-project-demo", "version": Version}).Info("started")

	// conf, err := internal.NewConf()
	// if err != nil {
	// 	log.Errorln("config error:", err)
	// }

	// svc := internal.NewService(conf)
	// svc.Run()

	app.Run()
}
