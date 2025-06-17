package middlwares

import (
	"api/internal/routes"
	"context"
	"strings"

	"github.com/Azat201003/eduflow_service_api/gen/go/user"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func AuthMiddlware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth_header := c.Request().Header["Authorization"]
		token := ""
		if len(auth_header) != 0 {
			token = strings.Split(auth_header[0], "Bearer ")[1]
		}
		var opts []grpc.CallOption
		authed_user, err := routes.UserClient.GetUserByToken(context.Background(), &user.Token{Token: token}, opts...)
		c.Set("is_user_setted", true)
		if err != nil {
			c.Set("is_user_setted", false)
		}
		c.Set("user", authed_user)

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
