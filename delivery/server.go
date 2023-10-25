package delivery

import (
	"api-payment/config"
	"api-payment/delivery/controller"
	"api-payment/delivery/middleware"
	"api-payment/manager"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine *gin.Engine
	host string
	log *logrus.Logger
}

func (a *appServer) initController(){
	a.engine.Use(middleware.LogRequestMiddleware(a.log))
	controller.NewMerchantController(a.engine,a.useCaseManager.MerchantUseCase())
	controller.NewPaymentController(a.engine,a.useCaseManager.PaymentUseCase())	
	controller.NewUserController(a.engine,a.useCaseManager.UserUseCase())
	controller.NewAuthController(a.engine,a.useCaseManager.AuthUseCase())	
}

func (a *appServer) Run(){
	a.initController()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err.Error())
	}
}

func Server()*appServer{
	engine := gin.Default()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln("Error Config : ()",err.Error())
	}
	infraManager, errConnect := manager.NewInfraManager(cfg)
	if errConnect != nil{
		log.Fatalln("Error Connection : ", errConnect.Error())
	}
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)

	host := fmt.Sprintf("%s:%s", cfg.ApiHost,cfg.ApiPort)
	return &appServer{	
		engine:engine ,
		useCaseManager: useCaseManager,
		host: host,
		log: logrus.New(),
	}
}