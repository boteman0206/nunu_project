package handler

import (
	"net/http"
	"projectName/internal/model"
	"projectName/internal/model/params"
	"projectName/internal/service"
	"projectName/pkg/helper/resp"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewUserHandler(handler *Handler,
	userService service.UserService,
) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

type UserHandler struct {
	*Handler
	userService service.UserService
}

// 登陆
func (h *UserHandler) Login(ctx *gin.Context) {

	params := params.LoginParams{}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, model.CodeParamErr, "", nil)
		return
	}
	h.logger.Info("Login", zap.Any("params", params))
	res, code, err := h.userService.Login(&params)
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, model.CodeNetError, "", nil)
		return
	}
	if code > 0 {
		resp.HandleError(ctx, http.StatusBadRequest, code, "", nil)
		return
	}

	resp.HandleSuccess(ctx, res)
}

func (h *UserHandler) LoginOut(ctx *gin.Context) {

	params := params.LoginOutParams{}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, model.CodeParamErr, "", nil)
		return
	}
	h.logger.Info("LoginOut", zap.Any("params", params))
	err := h.userService.LoginOut(&params)
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, model.CodeNetError, "", nil)
		return
	}
	resp.HandleSuccess(ctx, "")
}

// 注册
func (h *UserHandler) Register(ctx *gin.Context) {

	params := params.RegisterParams{}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, model.CodeParamErr, "", nil)
		return
	}

	h.logger.Info("Register", zap.Any("params", params))
	code, err := h.userService.Register(&params)
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, model.CodeNetError, "", nil)
		return
	}
	if code > 0 {
		resp.HandleError(ctx, http.StatusBadRequest, code, "", nil)
		return
	}
	resp.HandleSuccess(ctx, nil)
}

func (h *UserHandler) ChangePassword(ctx *gin.Context) {

	params := params.ChangeParams{}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, model.CodeParamErr, "", nil)
		return
	}

	h.logger.Info("ChangePassword", zap.Any("params", params))
	code, err := h.userService.ChangePassword(&params)
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, model.CodeNetError, "", nil)
		return
	}
	if code > 0 {
		resp.HandleError(ctx, http.StatusBadRequest, code, "", nil)
		return
	}
	resp.HandleSuccess(ctx, nil)
}

// 个人中心
func (h *UserHandler) UserInfoCenter(ctx *gin.Context) {

	params := params.CommonParam{}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, model.CodeParamErr, "", nil)
		return
	}

	// h.userService.UserInfoCenter(params)

	resp.HandleSuccess(ctx, nil)
}

// 更新用户信息
func (h *UserHandler) UpdateUserInfo(ctx *gin.Context) {

	params := params.CommonParam{}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, model.CodeParamErr, "", nil)
		return
	}

	resp.HandleSuccess(ctx, nil)
}
