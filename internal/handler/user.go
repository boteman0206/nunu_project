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

func (h *UserHandler) GetUserById(ctx *gin.Context) {
	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	user, err := h.userService.GetUserById(params.Id)
	h.logger.Info("GetUserByID", zap.Any("user", user))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, user)
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}

// 登陆
func (h *UserHandler) Login(ctx *gin.Context) {

	params := params.LoginParams{}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, model.CodeParamErr, "", nil)
		return
	}

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

// 注册
func (h *UserHandler) Register(ctx *gin.Context) {

	params := params.RegisterParams{}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, model.CodeParamErr, "", nil)
		return
	}

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
