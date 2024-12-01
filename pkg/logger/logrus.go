package logger

import (
	"os"

	"github.com/banggibima/be-itam/pkg/config"
	"github.com/sirupsen/logrus"
)

func Initialize(config *config.Config) (*logrus.Logger, error) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetOutput(os.Stdout)

	return logger, nil
}
