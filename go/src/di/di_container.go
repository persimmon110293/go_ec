package di

import (
	"main/handler"
)

type Container struct {
	HealthCheckHandler *handler.HealthCheckHandler
}

func NewContainer() *Container {
	// repositoryの初期化

	// serviceの初期化

	// handlerの初期化
	healthCheckHandler := handler.NewHealthCheckHandler()

	return &Container{
		HealthCheckHandler: healthCheckHandler,
	}
}
