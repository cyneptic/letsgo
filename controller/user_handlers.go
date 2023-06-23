package controller

import (
	"net/http"

	// repositories "github.com/cyneptic/letsgo/infrastructure/repository"
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/cyneptic/letsgo/internal/core/ports"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"github.com/cyneptic/letsgo/controller/middleware"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	echo *echo.Echo
	svc  ports.UserServiceContract
}

func NewHandler(svc ports.UserServiceContract, e *echo.Echo) *Handler {
	return &Handler{
		echo: e,
		svc:  svc,
	}
}

func (h *Handler) SetupRoutes() {
	h.echo.POST("/login", h.login)
	h.echo.POST("/register", h.register)
	h.echo.GET("/passengers/:userId", h.giveAllPassenger, middleware.AuthMiddleware)
	h.echo.POST("/passengers", h.addPassengersToUser, middleware.AuthMiddleware)
}

func (h *Handler) login(c echo.Context) error {
	user := new(entities.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	token, err := h.svc.LoginHandler(*user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Error": "Invalid Email Or Password",
		})
	}

	return c.JSON(200, token)
}

func (h *Handler) register(c echo.Context) error {

	newUser := new(entities.User)
	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	err := h.svc.AddUser(*newUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"newUser": newUser,
	})
}
func (h *Handler) logout(c echo.Context) error {

	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
	}
	err := h.svc.Logout(authHeader)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "unable to log out"})
	}

	return c.JSON(http.StatusOK, "logout successful")
}

func (h *Handler) giveAllPassenger(c echo.Context) error {
	userId := c.Param("userId")

	passengers, err := h.svc.GetAllUserPassengers(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusAccepted, passengers)
}

func (h *Handler) addPassengersToUser(c echo.Context) error {

	userId := c.Get("id").(string)
	passenger := new(entities.Passenger)

	if err := c.Bind(&passenger); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	passenger.UserID = uuid.MustParse(userId)

	err := h.svc.AddPassengersToUser(userId, *passenger)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "passenger added successfully")

}
