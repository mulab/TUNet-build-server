package worker

type Job interface {
}

type Repo struct {
	Fullname *string
	Owner    *string
	Repo     *string
	GitURL   *string
}

type PushEventJob struct {
	Head *string
	*Repo
}

type PullRequestEventJob struct {
	Number *string
	*Repo
}
