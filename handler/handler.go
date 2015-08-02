package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
)

type Handler interface {
	HandlePullRequest(pull *github.PullRequestEvent, c *gin.Context)
	HandlePush(push *github.PushEvent, c *gin.Context)
}

type EventHandler struct {
}
