package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mulab/TUNet-build-server/worker"
)

func PayloadHandlerFactory(h Handler, job <-chan worker.Job) (handler func(c *gin.Context)) {
	return
}
