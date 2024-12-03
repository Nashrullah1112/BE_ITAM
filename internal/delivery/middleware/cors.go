package middleware

import (
	"strconv"

	"github.com/banggibima/be-itam/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type CORSMiddleware struct {
	Config *config.Config
}

func NewCORSMiddleware(
	config *config.Config,
) *CORSMiddleware {
	return &CORSMiddleware{
		Config: config,
	}
}

func (c *CORSMiddleware) CORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: c.GetAllowedOrigins(),
		AllowHeaders: c.GetAllowedHeaders(),
	})
}

func (c *CORSMiddleware) GetAllowedOrigins() string {
	origins := []string{}

	for i := 3000; i <= 3010; i++ {
		origins = append(origins, "http://localhost:"+strconv.Itoa(i))
	}

	origins = append(origins, "http://localhost:5173")

	result := ""

	for i, element := range origins {
		result += element
		if i < len(origins)-1 {
			result += ","
		}
	}

	return result
}

func (c *CORSMiddleware) GetAllowedHeaders() string {
	headers := []string{
		"Origin",
		"Content-Type",
		"Accept",
		"Authorization",
	}

	result := ""

	for i, element := range headers {
		result += element
		if i < len(headers)-1 {
			result += ","
		}
	}

	return result
}
