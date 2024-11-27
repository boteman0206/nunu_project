package handler

import (
	"net/http"
	"projectName/internal/service"
	"projectName/pkg/helper/resp"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"projectName/internal/model/params"
	// "projectName/internal/model/response"
)

type LoginHandler struct {
	*Handler
	loginService service.LoginService
	userService  service.UserService
}

func NewLoginHandler(
	handler *Handler,
	loginService service.LoginService,
	userService service.UserService,
) *LoginHandler {
	return &LoginHandler{
		Handler:      handler,
		loginService: loginService,
		userService:  userService,
	}
}

func (h *LoginHandler) Login(ctx *gin.Context) {

	var params = params.LoginParams{}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	user, err := h.userService.GetUserById(params.ID)
	h.logger.Info("GetUserByID", zap.Any("user", user))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	if user.Password != "123456" {
		resp.HandleError(ctx, http.StatusOK, 1, "pawword is wrong", nil)
		return
	}

	h.loginService.Login(ctx, params)

	resp.HandleSuccess(ctx, user)
}

func (h *LoginHandler) LoginOut(ctx *gin.Context) {
	resp.HandleSuccess(ctx, "logout")

}
