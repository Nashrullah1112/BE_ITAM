package routes

import (
	"github.com/banggibima/be-itam/internal/delivery/handlers"
	"github.com/banggibima/be-itam/internal/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	App                *fiber.App
	DivisionHandler    *handlers.DivisionHandler
	PositionHandler    *handlers.PositionHandler
	VendorHandler      *handlers.VendorHandler
	AssetHandler       *handlers.AssetHandler
	ApplicationHandler *handlers.ApplicationHandler
	UserHandler        *handlers.UserHandler
	DeviceHandler      *handlers.DeviceHandler
	HardwareHandler    *handlers.HardwareHandler
	LicenseHandler     *handlers.LicenseHandler
	AuthHandler        *handlers.AuthHandler
	JWTMiddleware      *middleware.JWTMiddleware
}

func NewRoutes() *Routes {
	return &Routes{}
}

func (r *Routes) Resource() {
	prefix := r.App.Group("/api")

	auth := prefix.Group("/auth")
	auth.Post("/register", r.AuthHandler.Register)
	auth.Post("/login", r.AuthHandler.Login)

	divisions := prefix.Group("/divisions")
	divisions.Get("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor"), r.DivisionHandler.FindAll)
	divisions.Get("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor"), r.DivisionHandler.FindByID)
	divisions.Post("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.DivisionHandler.Create)
	divisions.Put("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.DivisionHandler.Update)
	divisions.Patch("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.DivisionHandler.UpdatePartial)
	divisions.Delete("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.DivisionHandler.Delete)

	positions := prefix.Group("/positions")
	positions.Get("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor"), r.PositionHandler.FindAll)
	positions.Get("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor"), r.PositionHandler.FindByID)
	positions.Post("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.PositionHandler.Create)
	positions.Put("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.PositionHandler.Update)
	positions.Patch("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.PositionHandler.UpdatePartial)
	positions.Delete("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.PositionHandler.Delete)

	vendors := prefix.Group("/vendors")
	vendors.Get("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor", "user"), r.VendorHandler.FindAll)
	vendors.Get("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor", "user"), r.VendorHandler.FindByID)
	vendors.Post("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.VendorHandler.Create)
	vendors.Put("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.VendorHandler.Update)
	vendors.Patch("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.VendorHandler.UpdatePartial)
	vendors.Delete("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.VendorHandler.Delete)

	assets := prefix.Group("/assets")
	assets.Get("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor"), r.AssetHandler.FindAll)
	assets.Get("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor"), r.AssetHandler.FindByID)
	assets.Post("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.AssetHandler.Create)
	assets.Put("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.AssetHandler.Update)
	assets.Patch("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.AssetHandler.UpdatePartial)
	assets.Delete("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.AssetHandler.Delete)

	applications := prefix.Group("/applications")
	applications.Get("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor"), r.ApplicationHandler.FindAll)
	applications.Get("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor"), r.ApplicationHandler.FindByID)
	applications.Post("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.ApplicationHandler.Create)
	applications.Put("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.ApplicationHandler.Update)
	applications.Patch("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.ApplicationHandler.UpdatePartial)
	applications.Delete("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.ApplicationHandler.Delete)

	users := prefix.Group("/users")
	users.Get("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor", "user"), r.UserHandler.FindAll)
	users.Get("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor", "user"), r.UserHandler.FindByID)
	users.Get("/email/:email", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor", "user"), r.UserHandler.FindByEmail)
	users.Post("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.UserHandler.Create)
	users.Put("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.UserHandler.Update)
	users.Patch("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.UserHandler.UpdatePartial)
	users.Delete("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.UserHandler.Delete)

	devices := prefix.Group("/devices")
	devices.Get("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor"), r.DeviceHandler.FindAll)
	devices.Get("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor"), r.DeviceHandler.FindByID)
	devices.Post("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.DeviceHandler.Create)
	devices.Put("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.DeviceHandler.Update)
	devices.Patch("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.DeviceHandler.UpdatePartial)
	devices.Delete("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.DeviceHandler.Delete)

	hardwares := prefix.Group("/hardwares")
	hardwares.Get("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor"), r.HardwareHandler.FindAll)
	hardwares.Get("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor"), r.HardwareHandler.FindByID)
	hardwares.Post("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.HardwareHandler.Create)
	hardwares.Put("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.HardwareHandler.Update)
	hardwares.Patch("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.HardwareHandler.UpdatePartial)
	hardwares.Delete("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.HardwareHandler.Delete)

	licenses := prefix.Group("/licenses")
	licenses.Get("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor"), r.LicenseHandler.FindAll)
	licenses.Get("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor"), r.LicenseHandler.FindByID)
	licenses.Post("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.LicenseHandler.Create)
	licenses.Put("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.LicenseHandler.Update)
	licenses.Patch("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.LicenseHandler.UpdatePartial)
	licenses.Delete("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin"), r.LicenseHandler.Delete)
}
