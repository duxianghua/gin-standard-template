package main

import (
	//"log"

	"fmt"

	"github.com/duxianghua/gin-standard-template/internal/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Version comes from build time
var Version string

func main() {
	log.WithFields(log.Fields{"app": "golang-project-demo", "version": Version}).Info("started")
	viper.SetDefault("Host", "0.0.0.0")
	viper.SetDefault("Port", 8080)
	viper.BindEnv("Host")
	viper.BindEnv("Port")
	// conf, err := internal.NewConf()
	// if err != nil {
	// 	log.Errorln("config error:", err)
	// }

	// svc := internal.NewService(conf)
	// svc.Run()
	log.Println("starting server...")

	api.Service().Run(fmt.Sprintf("%s:%d", viper.GetString("Host"), viper.GetInt("Port")))
}
