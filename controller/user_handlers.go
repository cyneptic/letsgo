package controller

import (
	"net/http"
	"time"

	// repositories "github.com/cyneptic/letsgo/infrastructure/repository"
	"github.com/cyneptic/letsgo/controller/middleware"
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/cyneptic/letsgo/internal/core/ports"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
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
	h.echo.POST("/logout", h.logout)
	h.echo.POST("/register", h.register)
	h.echo.GET("/passengers", h.giveAllPassenger, middleware.AuthMiddleware)
	h.echo.GET("/test", h.test, middleware.AuthMiddleware)
	h.echo.POST("/passengers", h.addPassengersToUser , middleware.AuthMiddleware)
}
func (h *Handler) test(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "test",
	})
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

	newUser.ID = uuid.New()
	newUser.CreatedAt = time.Now()
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	err := h.svc.AddUser(*newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
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
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "logout successful")
}

func (h *Handler) giveAllPassenger(c echo.Context) error {
	userId := c.Get("id").(string)

	passengers, err := h.svc.GetAllUserPassengers(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusAccepted, passengers)
}

func (h *Handler) addPassengersToUser(c echo.Context) error {

	
	passenger := new(entities.Passenger)
	userId := c.Get("id").(string)
	passenger.ID = uuid.New()
	passenger.UserID = uuid.MustParse(userId)
	passenger.CreatedAt = time.Now()

	if err := c.Bind(&passenger); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := h.svc.AddPassengersToUser(userId, *passenger)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "add passenger successfully")

}
