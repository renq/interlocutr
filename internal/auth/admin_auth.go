package adminAuth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

func NewAuthHandler(e *echo.Echo) {
	e.POST("/oauth/token", authHandler)
}

type loginRequest struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type JwtResponse struct {
	Token string `json:"token"`
}

// GenerateJWTToken   godoc
// @Summary      Generate JWT token
// @Description  Generate JWT token for valid user credentials
// @Accept       json
// @Accept       application/x-www-form-urlencoded
// @Param        username  formData  string               false  "login"
// @Param        password  formData  string               false  "password"
// @Param        request   body      loginRequest         false "Login credentials"
// @Tags         auth
// @Produce      json
// @Success      200       {object}  JwtResponse
// @Failure      400       {object}  infrastructure.ErrorResponse
// @Router       /oauth/token [post]
func authHandler(c *echo.Context) error {
	formData := loginRequest{}
	bindError := c.Bind(&formData)
	if bindError != nil {
		return c.JSON(http.StatusBadRequest, bindError)
	}

	if formData.Username != "admin" || formData.Password != "secret" {
		return echo.ErrUnauthorized
	}

	claims := &JwtCustomClaims{
		"Jon Snow",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	jwtResponse := JwtResponse{
		Token: t,
	}

	return c.JSON(http.StatusOK, jwtResponse)
}
