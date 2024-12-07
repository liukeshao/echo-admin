package handlers

import "C"
import (
	"fmt"
	"github.com/Oudwins/zog/zhttp"
	"github.com/labstack/echo/v4"
	"github.com/liukeshao/echo-admin/ent"
	"github.com/liukeshao/echo-admin/pkg/services"
	"github.com/liukeshao/echo-admin/types"
	"net/http"
)

type AccountHandler struct {
	auth *services.AuthService
	user *services.UserService
}

func NewAccountHandler(c *services.Container) *AccountHandler {
	return &AccountHandler{c.Auth, c.User}
}

// Login godoc
// @Summary 账号登录
// @Schemes
// @Description
// @Tags account
// @Accept json
// @Produce json
// @Param request body types.LoginInput true "params"
// @Success 200 {object} types.LoginOutput
// @Router /account/login [post]
func (h *AccountHandler) Login(ctx echo.Context) error {
	var input types.LoginInput
	errs := types.LoginInputSchema.Parse(zhttp.Request(ctx.Request()), &input)
	if errs != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errs)
	}

	authFailed := func() error {
		return echo.NewHTTPError(http.StatusForbidden, "Invalid credentials. Please try again.")
	}

	// Attempt to load the user
	u, err := h.user.FindByEmail(ctx.Request().Context(), input.Email)

	switch err.(type) {
	case *ent.NotFoundError:
		return Fail(fmt.Errorf("invalid credentials. Please try again"), "")
	case nil:
	default:
		return Fail(err, "error querying user during login")
	}

	// Check if the password is correct
	err = h.auth.CheckPassword(input.Password, u.Password)
	if err != nil {
		return authFailed()
	}

	// Log the user in
	token, err := h.auth.GenerateToken(u.ID)
	if err != nil {
		return Fail(err, "unable to log in user")
	}

	return ctx.JSON(http.StatusOK, types.LoginOutput{
		Token: token,
	})
}

// Register godoc
// @Summary 账号注册
// @Schemes
// @Description
// @Tags account
// @Accept json
// @Produce json
// @Param request body types.RegisterInput true "params"
// @Success 200 {object} ent.User
// @Router /account/register [post]
func (h *AccountHandler) Register(c echo.Context) error {
	var input types.RegisterInput
	errs := types.RegisterInputSchema.Parse(zhttp.Request(c.Request()), &input)
	if errs != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errs)
	}
	password, err := h.auth.HashPassword(input.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := h.user.Create(c.Request().Context(), input.Email, password)
	if err != nil {
		return Fail(err, "error creating user")
	}
	return c.JSON(http.StatusOK, user)
}

// Profile godoc
// @Summary 账号信息
// @Schemes
// @Description
// @Tags account
// @Accept json
// @Produce json
// @Success 200 {object} ent.User
// @Router /account/profile [get]
func (h *AccountHandler) Profile(ctx echo.Context) error {
	id, err := currentUserId(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := h.user.FindById(ctx.Request().Context(), id)
	if err != nil {
		return Fail(err, "error creating user")
	}
	return ctx.JSON(http.StatusOK, user)
}

func (h *AccountHandler) Logout(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "success")
}
