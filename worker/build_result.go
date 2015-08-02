package worker

type BuildResult struct {
	Owner       *string
	Repo        *string
	State       *string
	TargetURL   *string
	Description *string
}
