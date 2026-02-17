package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pranavp37/go_todo_ptsql/internal/config"
	"github.com/pranavp37/go_todo_ptsql/internal/utiles"
)

func RegisterMiddleware(c *echo.Echo) {
	c.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:   true,
		LogURI:      true,
		LogStatus:   true,
		LogLatency:  true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				log.Printf(
					"❌ %s %s → %d (%v) ERROR: %v",
					v.Method,
					v.URI,
					v.Status,
					v.Latency,
					v.Error,
				)
			} else {
				log.Printf(
					"✅ %s %s → %d (%v)",
					v.Method,
					v.URI,
					v.Status,
					v.Latency,
				)
			}
			return nil
		},
	}))
	// c.Use(middleware.RequestLogger())
	c.Use(middleware.Recover())
}

type JwtCustomeClims struct {
	users_id string `Json:"users_id"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(user_id string, jwt_secret string, expireAt time.Time) (string, error) {

	claims := &JwtCustomeClims{
		users_id: user_id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		log.Printf("Failed to sign token: %v", err)
		return "", err
	}
	return tokenString, nil
}

type jwtTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateAccessandRefershTokens(users_id string) (*jwtTokens, error) {

	config, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
		return nil, err
	}
	expireAtAccessToken := time.Now().Add(time.Hour * 24)

	access, err := GenerateJwtToken(users_id, config.JWT_SECRET_KEY, expireAtAccessToken)
	if err != nil {
		log.Fatal("Failed to generate access token:", err)
		return nil, err
	}
	expireAtRefreshToken := time.Now().Add(time.Hour * 24 * 7)
	refresh, err := GenerateJwtToken(users_id, config.JWT_SECRET_KEY, expireAtRefreshToken)
	if err != nil {
		log.Fatal("Failed to generate refresh token:", err)
		return nil, err
	}
	return &jwtTokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func AuthJwtMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header.Get("Authorization")
			if header == "" {
				return c.JSON(http.StatusUnauthorized, utiles.Response{
					Success: false,
					Message: "Authorization header required",
				})
			}
			headerList := strings.Split(header, " ")
			if len(headerList) != 2 || headerList[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, utiles.Response{
					Success: false,
					Message: "Invalid authorization header format",
				})

			}
			tokenString := headerList[1]
			config, err := config.Load()
			if err != nil {
				log.Fatal("Failed to load config:", err)
				return c.JSON(http.StatusInternalServerError, utiles.Response{
					Success: false,
					Message: "Internal server error",
				})
			}
			keyFunc := func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(config.JWT_SECRET_KEY), nil
			}
			token, err := jwt.ParseWithClaims(tokenString, &JwtCustomeClims{}, keyFunc)
			if err != nil {
				log.Printf("Failed to parse token: %v", err)
				return c.JSON(http.StatusUnauthorized, utiles.Response{
					Success: false,
					Message: "Invalid token",
				})
			}
			if claims, ok := token.Claims.(*JwtCustomeClims); ok && token.Valid {
				c.Set("users_id", claims.users_id)
				return next(c)
			} else {
				return c.JSON(http.StatusUnauthorized, utiles.Response{
					Success: false,
					Message: "Invalid token claims",
				})
			}
		}
	}
}
