package middleware

import (
	"time"

	"github.com/banggibima/be-itam/pkg/config"
	"github.com/banggibima/be-itam/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type LoggerMiddleware struct {
	Config *config.Config
	Logger *logrus.Logger
}

func NewLoggerMiddleware(
	config *config.Config,
	logger *logrus.Logger,
) *LoggerMiddleware {
	return &LoggerMiddleware{
		Config: config,
		Logger: logger,
	}
}

func (l *LoggerMiddleware) Logging() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		stop := time.Now()
		latency := stop.Sub(start)
		status := c.Response().StatusCode()

		utils.FormatLogRequest(l.Logger, string(c.Request().Header.Method()), c.OriginalURL(), status, int(latency.Nanoseconds()))

		return err
	}
}
