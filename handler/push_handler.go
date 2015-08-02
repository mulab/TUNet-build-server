package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
)

func (h *EventHandler) HandlePush(push *github.PushEvent, c *gin.Context) {
}
