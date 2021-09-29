package middleware

import (
	"consumerinformationmodule/pkg/controller_echo"
	"consumerinformationmodule/pkg/env"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoswagger "github.com/swaggo/echo-swagger"
	"strings"
)

func ConfigEchoNode(e *echo.Echo) string {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	default_ := e.Group("")

	defaultSecured_ := e.Group("")

	defaultSecured_.Use(getMiddleware("XIJWTAUTHNODE"))

	defaultSecured_.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			customClaimsJWT := c.Get("user").(*jwt.Token)

			jwtClaims := customClaimsJWT.Claims.(jwt.MapClaims)

			c.Set("Role", jwtClaims["Role"])

			return next(c)
		}
	})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	FN_Default(default_)
	FN_DefaultSecured(defaultSecured_)

	return env.RESTPORT
}

func getMiddleware(name string) echo.MiddlewareFunc {
	if strings.EqualFold(strings.ToUpper(name), "XIJWTAUTHNODE") {
		return middleware.JWTWithConfig(middleware.JWTConfig{
			ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
				keyFunc := func(t *jwt.Token) (interface{}, error) {
					if t.Method.Alg() != "HS256" {
						return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
					}
					return []byte(env.JWT_KEY_GETTER), nil
				}

				// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
				token, err := jwt.Parse(auth, keyFunc)
				if err != nil {
					return nil, err
				}
				if !token.Valid {
					return nil, errors.New("invalid token")
				}
				return token, nil
			},
		})

	}

	return nil

}

func FN_Default(g *echo.Group) {

	g.DELETE("/consumerinformationmodule/api/consumer", controller_echo.DeleteConsumers)
	g.PUT("/consumerinformationmodule/api/consumer", controller_echo.EditConsumer)
	g.GET("/consumerinformationmodule/api/consumer/profile", controller_echo.FilterCompanyById)
	g.GET("/consumerinformationmodule/api/consumerbyid", controller_echo.FindConsumer)
	g.GET("/consumerinformationmodule/api/consumer", controller_echo.ViewAllConsumers)
	g.GET("/consumerinformationmodule/api/swagger/*any", echoswagger.WrapHandler)
}

func FN_DefaultSecured(g *echo.Group) {

	g.POST("/consumerinformationmodule/api/consumer", controller_echo.CreateConsumer)
}
