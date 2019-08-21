package main

import (
	"applets/models"
	"applets/routes"
	"applets/system"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	defer totalDefer()
	serverConf := system.ReadServerIni()
	gin.SetMode(serverConf.RunMode)
	server := &http.Server{
		Addr:           serverConf.HttpPort,
		Handler:        routes.CreateRoute(),
		ReadTimeout:    serverConf.ReadTimeout,
		WriteTimeout:   serverConf.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	system.Save.InitInfo(fmt.Sprintf("start http server listening %s ", serverConf.HttpPort))
	if err := server.ListenAndServe(); err != nil {
		system.Save.ServerPanic("the http service has made some errors : ", err)
	}
}

func totalDefer() {
	system.Save.DeferFile()
	models.DeferDb()
}
