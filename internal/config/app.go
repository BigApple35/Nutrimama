package config

import (
	"nutrimama/internal/delivery/http"
	"nutrimama/internal/delivery/http/route"
	"nutrimama/internal/repository"
	"nutrimama/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB *gorm.DB
	App *fiber.App
}

func Bootstrap(config *BootstrapConfig) {
	
	userRepo := repository.NewUserRepository()
	motherRepo := repository.NewMotherRepository()
	consultantRepo := repository.NewConsultantRepository()
	
	userUseCase := usecase.NewUserUseCase(config.DB, userRepo, motherRepo, consultantRepo)
	userController := http.NewUserController(userUseCase)

	eduToolsRepo := repository.NewEduToolsRepository()
	eduToolsUseCase := usecase.NewEduToolsUseCase(config.DB, eduToolsRepo)
	eduToolsController := http.NewEduToolsController(eduToolsUseCase)

	trackingLogRepo := repository.NewTrackingLogRepository()
	trackingQuestionRepo := repository.NewTrackingQuestionRepository()
	trackingUseCase := usecase.NewTrackingUseCase(config.DB, trackingLogRepo, trackingQuestionRepo, motherRepo)
	trackingController := http.NewTrackingController(trackingUseCase)

	routeConfig := &route.RouteConfig{
		App:                config.App,
		UserController:     userController,
		EduToolsController: eduToolsController,
		TrackingController: trackingController,
	}
	routeConfig.Setup()
}