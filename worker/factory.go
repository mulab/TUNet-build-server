package worker

import (
	"github.com/mulab/TUNet-build-server/helper"
)

func BuildWorkerFactory(git *helper.GitHelper, in <-chan Job, out chan<- string, result chan<- BuildResult) (worker func()) {
	return
}

func StatusWorkerFactory(token *string, result <-chan BuildResult) (worker func()) {
	return
}

func LogWorkerFactory(path *string, out chan<- string) (worker func()) {
	return
}
