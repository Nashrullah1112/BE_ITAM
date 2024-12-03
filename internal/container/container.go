package container

import (
	"database/sql"
	"strconv"

	"github.com/banggibima/be-itam/internal/delivery/handlers"
	"github.com/banggibima/be-itam/internal/delivery/middleware"
	"github.com/banggibima/be-itam/internal/delivery/routes"
	authcommand "github.com/banggibima/be-itam/modules/auth/application/command"
	usercommand "github.com/banggibima/be-itam/modules/user/application/command"
	userquery "github.com/banggibima/be-itam/modules/user/application/query"
	userdomain "github.com/banggibima/be-itam/modules/user/domain"
	userrepository "github.com/banggibima/be-itam/modules/user/infrastructure/repository"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
)

type Container struct {
	Config  *config.Config
	App     *fiber.App
	Logger  *logrus.Logger
	DB      *sql.DB
	Storage *minio.Client
}

func NewContainer(
	config *config.Config,
	app *fiber.App,
	logger *logrus.Logger,
	db *sql.DB,
	storage *minio.Client,
) *Container {
	return &Container{
		Config:  config,
		App:     app,
		Logger:  logger,
		DB:      db,
		Storage: storage,
	}
}

func (c *Container) Setup() error {
	postgresUserRepository := userrepository.NewPostgresUserRepository(c.Config, c.DB, c.Logger)
	userService := userdomain.NewUserService(c.Config, postgresUserRepository)
	userCommandUsecase := usercommand.NewUserCommandUsecase(c.Config, userService)
	userQueryUsecase := userquery.NewUserQueryUsecase(c.Config, userService)
	userHandler := handlers.NewUserHandler(c.Config, userCommandUsecase, userQueryUsecase)

	authCommandUsecase := authcommand.NewAuthCommandUsecase(c.Config, userService)
	authHandler := handlers.NewAuthHandler(c.Config, authCommandUsecase)

	corsMiddleware := middleware.NewCORSMiddleware(c.Config)
	jwtMiddleware := middleware.NewJWTMiddleware(c.Config)
	loggerMiddleware := middleware.NewLoggerMiddleware(c.Config, c.Logger)

	c.App.Use(corsMiddleware.CORS())
	c.App.Use(loggerMiddleware.Logging())

	routes := &routes.Routes{
		App:           c.App,
		UserHandler:   userHandler,
		AuthHandler:   authHandler,
		JWTMiddleware: jwtMiddleware,
	}

	routes.Resource()

	return nil
}

func (c *Container) Start() error {
	c.Logger.Info("server is running on port " + strconv.Itoa(c.Config.App.Port))
	c.App.Listen(":" + strconv.Itoa(c.Config.App.Port))

	return nil
}
