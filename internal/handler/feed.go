package handler

import (
	"github.com/gin-gonic/gin"
	"projectName/internal/service"
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
		Handler:      handler,
		feedService: feedService,
	}
}

func (h *FeedHandler) GetFeed(ctx *gin.Context) {

}
