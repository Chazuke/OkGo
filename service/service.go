package service

import (
	"okgo/logger"
	"time"

	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
)

type (
	Service interface {
		Attach(string, *fiber.App, *logger.OkGoLogger) (func(), error)
		Config() any
		Description() string
		Name() string
	}

	Project[T Service] struct {
		Fiber *fiber.App

		logger   *logger.OkGoLogger
		service  T
		shutdown func(func())
	}
)

var (
	compiledAt string
	executedAt time.Time
	gitBranch  string
	gitCommit  string
	version    string
)

func init() {
	executedAt = time.Now().UTC()
}

func new[T Service](service T, projectId string, logger *logger.OkGoLogger) (*Project[T], error) {
	s := &Project[T]{
		Fiber:   fiber.New(),
		service: service,
	}

	s.Fiber.Use(otelfiber.Middleware(WithServerName(service.Name()))))

	return s, nil
}
