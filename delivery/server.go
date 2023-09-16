package delivery

import (
	"fmt"
	"polen/config"
	"polen/delivery/controller/api"

	"polen/delivery/middleware"
	"polen/manager"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *gin.Engine
	host      string
	log       *logrus.Logger
}

func (s *Server) Run() {
	s.initMiddlewares()
	s.initControllers()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initMiddlewares() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
}

func (s *Server) initControllers() {
	rg := s.engine.Group("/api/v1")
	api.NewAuthController(s.ucManager.UserUseCase(), s.ucManager.AuthUseCase(), rg).Route()
	api.NewBiodataController(s.ucManager.BiodataUserUseCase(), rg).Route()
	api.NewTopUpController(s.ucManager.TopUpUsecase(), s.ucManager.BiodataUserUseCase(), rg).Route()
	api.NewDepositeInterestController(s.ucManager.DepositerInterestUseCase(), rg).Route()
	api.NewLoanInterestController(s.ucManager.LoanInterestUseCase(), rg).Route()
	api.NewSaldoController(s.ucManager.SaldoUsecase(), rg).Route()
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		fmt.Println(err)
	}
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUsecaseManager(repoManager)

	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	log := logrus.New()

	engine := gin.Default()
	return &Server{
		ucManager: useCaseManager,
		engine:    engine,
		host:      host,
		log:       log,
	}
}
