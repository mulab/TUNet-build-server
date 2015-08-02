package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
)

func (h *EventHandler) HandlePullRequest(pull *github.PullRequestEvent, c *gin.Context) {
}
