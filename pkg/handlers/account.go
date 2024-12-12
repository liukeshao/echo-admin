package handlers

import (
	"github.com/liukeshao/echo-admin/pkg/log"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/liukeshao/echo-admin/ent"
	"github.com/liukeshao/echo-admin/pkg/echox"
	"github.com/liukeshao/echo-admin/pkg/services"
	"github.com/liukeshao/echo-admin/types"
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
func (h *AccountHandler) Login(c echo.Context) error {
	var input types.LoginInput
	err := echox.Bind(c, types.LoginInputSchema, &input)
	if err != nil {
		return err
	}

	// Attempt to load the user
	u, err := h.user.FindByEmail(c.Request().Context(), input.Email)
	switch {
	// If the entity does not meet a specific condition,
	// the operation will return an "ent.NotFoundError".
	case ent.IsNotFound(err):
		log.Ctx(c).Warn("user not found", slog.String("email", input.Email))
		return echox.Error401Unauthorized("user not found")
		// Any other error.
	case err != nil:
		log.Ctx(c).Error(
			"error querying user during login",
			slog.Any("input", input),
			slog.Any("error", err),
		)
		return echox.Error500InternalServerError("error querying user during login")
	}

	// Check if the password is correct
	err = h.auth.CheckPassword(input.Password, u.Password)
	if err != nil {
		return echox.Error401Unauthorized("invalid password", err)
	}

	// Log the user in
	token, err := h.auth.GenerateToken(u.ID)
	if err != nil {
		return echox.Error500InternalServerError("generate token error", err)
	}

	return c.JSON(http.StatusOK, types.LoginOutput{
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
	err := echox.Bind(c, types.RegisterInputSchema, &input)
	if err != nil {
		return err
	}
	password, err := h.auth.HashPassword(input.Password)
	if err != nil {
		return echox.Error500InternalServerError("password hash error", err)
	}
	user, err := h.user.Create(c.Request().Context(), input.Email, password)
	if err != nil {
		return echox.Error500InternalServerError("error creating user", err)
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
		return echox.Error400BadRequest("invalid user id", err)
	}
	user, err := h.user.FindById(ctx.Request().Context(), id)
	switch err.(type) {
	case *ent.NotFoundError:
		return echox.Error401Unauthorized("user not found")
	case nil:
	default:
		return echox.Error500InternalServerError("error querying user during profile", err)
	}

	return ctx.JSON(http.StatusOK, user)
}

func (h *AccountHandler) Logout(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "success")
}
