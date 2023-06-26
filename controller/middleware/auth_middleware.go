package middleware

import (
	"fmt"
	"net/http"
	repositories "github.com/cyneptic/letsgo/infrastructure/repository"
	"github.com/labstack/echo/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)
var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
    SigningKey: []byte("secret"),
    
})
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		redis := repositories.RedisInit()

		
		authHeader := c.Request().Header.Get("Authorization")

		

		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
		}
		val , err := redis.TokenReceiver(authHeader)

		if val == "0" || val == "false" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token "})
		}

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token "})
		}

		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			
			return []byte("secret"), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token "})
		}
		
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			
			id, ok := claims["user_id"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token "})
			}

			
			c.Set("id", id)

			
			return next(c)
		}

		
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
	}
}
/*
func handler(c echo.Context) error {
	id := c.Get("id").(string)
	## Use the "id" in your business logic
}
*/