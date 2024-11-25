package delivery

import (
	"database/sql"
	"fmt"
	"test-mnc/config"
	"test-mnc/delivery/controller"
	"test-mnc/delivery/middleware"
	"test-mnc/logger"
	"test-mnc/repository"
	"test-mnc/shared/service"
	"test-mnc/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	authUc      usecase.AuthUsecase
	merchantUc  usecase.MerchantUsecase
	productUc   usecase.ProductUsecase
	topupUc     usecase.TopupUsecase
	blacklistUc usecase.BlacklistUsecase
	jwtService  service.JwtService
	engine      *gin.Engine
	host        string
}

var log = logger.NewLogger()

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)
	authMiddleware := middleware.NewAuthMiddleware(s.blacklistUc)

	controller.NewCustomerController(s.authUc, rg, &log).Route()
	controller.NewMerchantController(s.merchantUc, rg, &log).Route()
	controller.NewProductController(s.productUc, rg, &log).Route()
	controller.NewTopupController(s.topupUc, s.productUc, authMiddleware, rg, &log).Route()
	controller.NewBlacklistController(s.blacklistUc, rg, &log).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running on host %s, becauce error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic(err)
	}

	JwtService := service.NewJwtService(cfg.TokenConfig)
	customerRepo := repository.NewCustomerRepository(db)
	customerUc := usecase.NewCustomerUsecase(customerRepo)
	authUc := usecase.NewAuthUsecase(customerUc, JwtService)

	merchantRepo := repository.NewMerchantRepository(db)
	merchantUc := usecase.NewMerchantUsecase(merchantRepo)

	productRepo := repository.NewProductRepository(db)
	productUc := usecase.NewProductUsecase(productRepo)

	topupRepo := repository.NewTopupRepository(db)
	topupUc := usecase.NewTopupUsecase(topupRepo)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	blacklistRepo := repository.NewBlacklistRepository(redisClient)
	blacklistUc := usecase.NewBlacklistUsecase(blacklistRepo, JwtService)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		authUc:      authUc,
		merchantUc:  merchantUc,
		productUc:   productUc,
		topupUc:     topupUc,
		blacklistUc: blacklistUc,
		jwtService:  JwtService,
		engine:      engine,
		host:        host,
	}
}
