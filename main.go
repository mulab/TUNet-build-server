package main

import (
    "github.com/gin-gonic/gin"
    "github.com/mulab/TUNet-build-server/handler"
    "github.com/mulab/TUNet-build-server/worker"
    "github.com/mulab/TUNet-build-server/helper"
)

var token string
var logPath string

func main() {
    eventHandler := &handler.EventHandler{}
    git := &helper.GitHelper{}
    job := make(chan worker.Job)
    output := make(chan string)
    result := make(chan worker.BuildResult)
    buildWorkder := worker.BuildWorkerFactory(git, job, output, result)
    statusWorker := worker.StatusWorkerFactory(&token, result)
    logWorker := worker.LogWorkerFactory(&logPath, output)

    go buildWorkder()
    go statusWorker()
    go logWorker()

    router := gin.Default()

    router.POST("/payload", handler.PayloadHandlerFactory(eventHandler, job))

    router.Run(":8080")
}
