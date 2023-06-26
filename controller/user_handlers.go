package controller

import (
	"net/http"
	"time"

	// repositories "github.com/cyneptic/letsgo/infrastructure/repository"
	"github.com/cyneptic/letsgo/controller/middleware"
	"github.com/cyneptic/letsgo/controller/validators"
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/cyneptic/letsgo/internal/core/ports"
	"github.com/cyneptic/letsgo/internal/core/service"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	echo *echo.Echo
	svc  ports.UserServiceContract
}

func NewUserHandler() *UserHandler {
	svc := service.NewUserService()
	return &UserHandler{
		svc: svc,
	}
}

func AddUserRoutes(e *echo.Echo) {

	h := NewUserHandler()

	h.echo.POST("/login", h.login)
	h.echo.POST("/logout", h.logout)
	h.echo.POST("/register", h.register)
	h.echo.GET("/passengers", h.giveAllPassenger, middleware.AuthMiddleware)

	h.echo.POST("/passengers", h.addPassengersToUser, middleware.AuthMiddleware)
}

// validation done
func (h *UserHandler) login(c echo.Context) error {
	user := new(entities.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	err := validators.ValidateUserLogin(*user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	token, err := h.svc.LoginHandler(*user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Error": "Invalid Email Or Password",
		})
	}

	return c.JSON(200, token)
}

func (h *UserHandler) register(c echo.Context) error {

	newUser := new(entities.User)

	newUser.ID = uuid.New()
	newUser.CreatedAt = time.Now()
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	err := validators.ValidateUserRegister(*newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.svc.AddUser(*newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"newUser": newUser,
	})
}

// validation done
func (h *UserHandler) logout(c echo.Context) error {

	authHeader := c.Request().Header.Get("Authorization")
	err := validators.ValidateAuthToken(authHeader)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = h.svc.Logout(authHeader)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, "logout successful")
}

// validation done
func (h *UserHandler) giveAllPassenger(c echo.Context) error {

	userId := c.Get("id").(string)

	err := validators.ValidateUserID(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	passengers, err := h.svc.GetAllUserPassengers(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusAccepted, passengers)
}
// validation done
func (h *UserHandler) addPassengersToUser(c echo.Context) error {
	userId := c.Get("id").(string)
	err := validators.ValidateUserID(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	passenger := new(entities.Passenger)

	passenger.ID = uuid.New()
	passenger.UserID = uuid.MustParse(userId)
	passenger.CreatedAt = time.Now()
	if err := c.Bind(&passenger); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = validators.ValidatePassenger(*passenger)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = h.svc.AddPassengersToUser(userId, *passenger)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "add passenger successfully")
}
