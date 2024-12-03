package routes

import (
	"github.com/banggibima/be-itam/internal/delivery/handlers"
	"github.com/banggibima/be-itam/internal/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	App           *fiber.App
	UserHandler   *handlers.UserHandler
	AuthHandler   *handlers.AuthHandler
	JWTMiddleware *middleware.JWTMiddleware
}

func NewRoutes() *Routes {
	return &Routes{}
}

func (r *Routes) Resource() {
	prefix := r.App.Group("/api")

	auth := prefix.Group("/auth")
	auth.Post("/register", r.AuthHandler.Register)
	auth.Post("/login", r.AuthHandler.Login)

	users := prefix.Group("/users")
	users.Get("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor", "user"), r.UserHandler.FindAll)
	users.Get("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor", "user"), r.UserHandler.FindByID)
	users.Get("/email/:email", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor", "user"), r.UserHandler.FindByEmail)
	users.Post("", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor", "user"), r.UserHandler.Create)
	users.Put("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor", "user"), r.UserHandler.Update)
	users.Patch("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor", "user"), r.UserHandler.UpdatePartial)
	users.Delete("/:id", r.JWTMiddleware.Authentication, r.JWTMiddleware.Authorization("admin", "supervisor", "user"), r.UserHandler.Delete)
}
