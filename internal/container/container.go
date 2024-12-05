package container

import (
	"database/sql"
	"strconv"

	"github.com/banggibima/be-itam/internal/delivery/handlers"
	"github.com/banggibima/be-itam/internal/delivery/middleware"
	"github.com/banggibima/be-itam/internal/delivery/routes"
	applicationcommand "github.com/banggibima/be-itam/modules/applications/application/command"
	applicationquery "github.com/banggibima/be-itam/modules/applications/application/query"
	applicationdomain "github.com/banggibima/be-itam/modules/applications/domain"
	applicationrepository "github.com/banggibima/be-itam/modules/applications/infrastructure/repository"
	assetcommand "github.com/banggibima/be-itam/modules/assets/application/command"
	assetquery "github.com/banggibima/be-itam/modules/assets/application/query"
	assetdomain "github.com/banggibima/be-itam/modules/assets/domain"
	assetrepository "github.com/banggibima/be-itam/modules/assets/infrastructure/repository"
	authcommand "github.com/banggibima/be-itam/modules/auth/application/command"
	devicecommand "github.com/banggibima/be-itam/modules/devices/application/command"
	devicequery "github.com/banggibima/be-itam/modules/devices/application/query"
	devicedomain "github.com/banggibima/be-itam/modules/devices/domain"
	devicerepository "github.com/banggibima/be-itam/modules/devices/infrastructure/repository"
	divisioncommand "github.com/banggibima/be-itam/modules/divisions/application/command"
	divisionquery "github.com/banggibima/be-itam/modules/divisions/application/query"
	divisiondomain "github.com/banggibima/be-itam/modules/divisions/domain"
	divisionrepository "github.com/banggibima/be-itam/modules/divisions/infrastructure/repository"
	hardwarecommand "github.com/banggibima/be-itam/modules/hardwares/application/command"
	hardwarequery "github.com/banggibima/be-itam/modules/hardwares/application/query"
	hardwaredomain "github.com/banggibima/be-itam/modules/hardwares/domain"
	hardwarerepository "github.com/banggibima/be-itam/modules/hardwares/infrastructure/repository"
	licensecommand "github.com/banggibima/be-itam/modules/licenses/application/command"
	licensequery "github.com/banggibima/be-itam/modules/licenses/application/query"
	licensedomain "github.com/banggibima/be-itam/modules/licenses/domain"
	licenserepository "github.com/banggibima/be-itam/modules/licenses/infrastructure/repository"
	positioncommand "github.com/banggibima/be-itam/modules/positions/application/command"
	positionquery "github.com/banggibima/be-itam/modules/positions/application/query"
	positiondomain "github.com/banggibima/be-itam/modules/positions/domain"
	positionrepository "github.com/banggibima/be-itam/modules/positions/infrastructure/repository"
	usercommand "github.com/banggibima/be-itam/modules/users/application/command"
	userquery "github.com/banggibima/be-itam/modules/users/application/query"
	userdomain "github.com/banggibima/be-itam/modules/users/domain"
	userrepository "github.com/banggibima/be-itam/modules/users/infrastructure/repository"
	vendorcommand "github.com/banggibima/be-itam/modules/vendors/application/command"
	vendorquery "github.com/banggibima/be-itam/modules/vendors/application/query"
	vendordomain "github.com/banggibima/be-itam/modules/vendors/domain"
	vendorrepository "github.com/banggibima/be-itam/modules/vendors/infrastructure/repository"
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
	divisionPostgresRepository := divisionrepository.NewPostgresDivisionRepository(c.Config, c.DB, c.Logger)
	divisionService := divisiondomain.NewDivisionService(c.Config, divisionPostgresRepository)
	divisionCommandUsecase := divisioncommand.NewDivisionCommandUsecase(c.Config, divisionService)
	divisionQueryUsecase := divisionquery.NewDivisionQueryUsecase(c.Config, divisionService)
	divisionHandler := handlers.NewDivisionHandler(c.Config, divisionCommandUsecase, divisionQueryUsecase)

	positionPostgresRepository := positionrepository.NewPostgresPositionRepository(c.Config, c.DB, c.Logger)
	positionService := positiondomain.NewPositionService(c.Config, positionPostgresRepository)
	positionCommandUsecase := positioncommand.NewPositionCommandUsecase(c.Config, positionService)
	positionQueryUsecase := positionquery.NewPositionQueryUsecase(c.Config, positionService)
	positionHandler := handlers.NewPositionHandler(c.Config, positionCommandUsecase, positionQueryUsecase)

	vendorPostgresRepository := vendorrepository.NewPostgresVendorRepository(c.Config, c.DB, c.Logger)
	vendorService := vendordomain.NewVendorService(c.Config, vendorPostgresRepository)
	vendorCommandUsecase := vendorcommand.NewVendorCommandUsecase(c.Config, vendorService)
	vendorQueryUsecase := vendorquery.NewVendorQueryUsecase(c.Config, vendorService)
	vendorHandler := handlers.NewVendorHandler(c.Config, vendorCommandUsecase, vendorQueryUsecase)

	assetPostgresRepository := assetrepository.NewPostgresAssetRepository(c.Config, c.DB, c.Logger)
	assetService := assetdomain.NewAssetService(c.Config, assetPostgresRepository)
	assetCommandUsecase := assetcommand.NewAssetCommandUsecase(c.Config, assetService)
	assetQueryUsecase := assetquery.NewAssetQueryUsecase(c.Config, assetService)
	assetHandler := handlers.NewAssetHandler(c.Config, assetCommandUsecase, assetQueryUsecase)

	applicationPostgresRepository := applicationrepository.NewPostgresApplicationRepository(c.Config, c.DB, c.Logger)
	applicationService := applicationdomain.NewApplicationService(c.Config, applicationPostgresRepository)
	applicationCommandUsecase := applicationcommand.NewApplicationCommandUsecase(c.Config, applicationService)
	applicationQueryUsecase := applicationquery.NewApplicationQueryUsecase(c.Config, applicationService)
	applicationHandler := handlers.NewApplicationHandler(c.Config, applicationCommandUsecase, applicationQueryUsecase)

	userPostgresRepository := userrepository.NewPostgresUserRepository(c.Config, c.DB, c.Logger)
	userService := userdomain.NewUserService(c.Config, userPostgresRepository)
	userCommandUsecase := usercommand.NewUserCommandUsecase(c.Config, userService)
	userQueryUsecase := userquery.NewUserQueryUsecase(c.Config, userService)
	userHandler := handlers.NewUserHandler(c.Config, userCommandUsecase, userQueryUsecase)

	devicePostgresRepository := devicerepository.NewPostgresDeviceRepository(c.Config, c.DB, c.Logger)
	deviceService := devicedomain.NewDeviceService(c.Config, devicePostgresRepository)
	deviceCommandUsecase := devicecommand.NewDeviceCommandUsecase(c.Config, deviceService)
	deviceQueryUsecase := devicequery.NewDeviceQueryUsecase(c.Config, deviceService)
	deviceHandler := handlers.NewDeviceHandler(c.Config, deviceCommandUsecase, deviceQueryUsecase)

	hardwarePostgresRepository := hardwarerepository.NewPostgresHardwareRepository(c.Config, c.DB, c.Logger)
	hardwareService := hardwaredomain.NewHardwareService(c.Config, hardwarePostgresRepository)
	hardwareCommandUsecase := hardwarecommand.NewHardwareCommandUsecase(c.Config, hardwareService)
	hardwareQueryUsecase := hardwarequery.NewHardwareQueryUsecase(c.Config, hardwareService)
	hardwareHandler := handlers.NewHardwareHandler(c.Config, hardwareCommandUsecase, hardwareQueryUsecase)

	licensePostgresRepository := licenserepository.NewPostgresLicenseRepository(c.Config, c.DB, c.Logger)
	licenseService := licensedomain.NewLicenseService(c.Config, licensePostgresRepository)
	licenseCommandUsecase := licensecommand.NewLicenseCommandUsecase(c.Config, licenseService)
	licenseQueryUsecase := licensequery.NewLicenseQueryUsecase(c.Config, licenseService)
	licenseHandler := handlers.NewLicenseHandler(c.Config, licenseCommandUsecase, licenseQueryUsecase)

	authCommandUsecase := authcommand.NewAuthCommandUsecase(c.Config, userService)
	authHandler := handlers.NewAuthHandler(c.Config, authCommandUsecase)

	corsMiddleware := middleware.NewCORSMiddleware(c.Config)
	jwtMiddleware := middleware.NewJWTMiddleware(c.Config)
	loggerMiddleware := middleware.NewLoggerMiddleware(c.Config, c.Logger)

	c.App.Use(corsMiddleware.CORS())
	c.App.Use(loggerMiddleware.Logging())

	routes := &routes.Routes{
		App:                c.App,
		DivisionHandler:    divisionHandler,
		PositionHandler:    positionHandler,
		VendorHandler:      vendorHandler,
		AssetHandler:       assetHandler,
		ApplicationHandler: applicationHandler,
		UserHandler:        userHandler,
		DeviceHandler:      deviceHandler,
		HardwareHandler:    hardwareHandler,
		LicenseHandler:     licenseHandler,
		AuthHandler:        authHandler,
		JWTMiddleware:      jwtMiddleware,
	}

	routes.Resource()

	return nil
}

func (c *Container) Start() error {
	c.Logger.Info("server is running on port " + strconv.Itoa(c.Config.App.Port))
	c.App.Listen(":" + strconv.Itoa(c.Config.App.Port))

	return nil
}
