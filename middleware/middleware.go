package middleware

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/basicauth"

	"afeilulu.com/example/config"
)

// AuthReq middleware
func AuthReq() func(*fiber.Ctx) error {

	cfg := basicauth.Config{
		Users: map[string]string{
			config.Config("USERNAME"): config.Config("PASSWORD"),
		},
	}

	err := basicauth.New(cfg)

	return err

}
