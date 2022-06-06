package api

import (
	"course/internal/sms_center"
	"fmt"
	"github.com/gin-gonic/gin"
)

type API struct {
	SMSCenter *sms_center.SMSCenter
}

func (api *API) Run() {
	server := gin.Default()
	api.initRoutes(server)
	err := server.Run()
	if err != nil {
		fmt.Println("server exited with error: " + err.Error())
	}
}

func (api *API) initRoutes(s *gin.Engine) {
	//s.Use(CORSMiddleware())
	g := s.Group("/api")
	g.GET("/ping", handlerPing)
	g.GET("/sms-center/", api.handlerGetSMSCenter)
	g.POST("/towers/", api.handlerCreateTower)
	g.POST("/towers/:towerNumber/phones/", api.handlerCreatePhone)
	g.GET("/towers/:towerNumber/phones/:phoneNumber/sms", api.handlerGetSMSList)
	g.POST("/sms/", api.handlerSendSMS)
}
