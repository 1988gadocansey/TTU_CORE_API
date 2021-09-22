package middlewares

import (
	"TTU_CORE_PORTAL_GO/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
)

func RequestAuthencation(ctx *fiber.Ctx) error {
	ctx.Set("X-XSS-Protection", "1; mode=block")
	ctx.Set("Strict-Transport-Security", "max-age=5184000")
	ctx.Set("X-DNS-Prefetch-Control", "off")

	url_whitelist := []string{
		"/api/v1/user/login",
		"/api/v1/user/create",
		"/api/v1/levels",
		"/api/v1/students",
	}

	// Check if url is in whitelist
	if helpers.InArray(ctx.OriginalURL(), url_whitelist) || strings.Contains(ctx.OriginalURL(), "cdn") {
		return ctx.Next()
	}

	authorization_token := string(ctx.Request().Header.Peek("Authorization"))

	// Check if user passed authorization header.
	if authorization_token == "" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	// Check if token is valid
	_, err := jwt.ParseWithClaims(authorization_token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Next()
}
