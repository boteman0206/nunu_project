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

type FeedHandler struct {
	*Handler
	feedService service.FeedService
}

func NewFeedHandler(
	handler *Handler,
	feedService service.FeedService,
) *FeedHandler {
	return &FeedHandler{
		Handler:     handler,
		feedService: feedService,
	}
}

func (h *FeedHandler) GetFeed(ctx *gin.Context) {

}

func (h *FeedHandler) CreateFeed(ctx *gin.Context) {

	params := params.LoginParams{}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, model.CodeParamErr, "", nil)
		return
	}

	h.logger.Info("Login", zap.Any("params", params))
	code, err := h.feedService.CreateFeed(ctx, &model.Feed{})
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
