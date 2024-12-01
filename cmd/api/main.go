package main

import (
	"github.com/banggibima/be-itam/internal/container"
	"github.com/banggibima/be-itam/internal/persistence/postgres"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/banggibima/be-itam/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger, err := logger.Initialize(config)
	if err != nil {
		panic(err)
	}

	db, err := postgres.Client(config, logger)
	if err != nil {
		panic(err)
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	deps := container.NewContainer(config, app, logger, db, nil)

	if err := deps.Setup(); err != nil {
		panic(err)
	}

	if err := deps.Start(); err != nil {
		panic(err)
	}
}
