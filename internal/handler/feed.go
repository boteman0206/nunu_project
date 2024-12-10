package handler

import (
	"encoding/json"
	"net/http"
	"projectName/internal/model"
	"projectName/internal/model/params"
	"projectName/internal/service"
	"projectName/pkg/helper/resp"
	"strings"
	"time"

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
	params := params.GetFeeInfoParams{}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, model.CodeParamErr, "", nil)
		return
	}

	h.logger.Info("Login", zap.Any("params", params))
	data, err := h.feedService.GetFeed(ctx, params.FeedID)
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, model.CodeNetError, "", nil)
		return
	}

	resp.HandleSuccess(ctx, data)
}

func (h *FeedHandler) CreateFeed(ctx *gin.Context) {

	params := params.CreateFeedParams{}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, model.CodeParamErr, "", nil)
		return
	}

	var (
		tags                  = make([]string, 0)
		imgaeList             = make([]string, 0)
		tagsStr, imageListStr []byte
		feedType              = 0 // 1为普通投稿,2图片投稿3为视频投稿
	)

	if len(params.Tags) > 0 {
		tags = strings.Split(params.Tags, ",")
		tagsStr, _ = json.Marshal(tags)
	}
	if len(params.ImageList) > 0 {
		imgaeList = strings.Split(params.ImageList, ",")
		imageListStr, _ = json.Marshal(imgaeList)
		feedType = 1
	}

	t := time.Now().Unix()
	input := &model.Feed{
		Title:       params.Title,
		Tag:         string(tagsStr),
		Description: params.Description,
		ImageList:   string(imageListStr),
		DbTime:      t,
		UpdateTime:  t,
		Type:        int8(feedType),
	}

	h.logger.Info("CreateFeed", zap.Any("params", params))
	affetcRows, err := h.feedService.CreateFeed(ctx, input)
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, model.CodeNetError, "", nil)
		return
	}
	h.logger.Info("CreateFeed", zap.Int64("affetcRows", affetcRows))
	resp.HandleSuccess(ctx, affetcRows)
}
