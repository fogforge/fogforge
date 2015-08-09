package model

type Build interface {
	LatestStatus() BuildStatus
	SourceCommitRef() *CommitRef
	LatestAction() (*BuildAction, error)
	Action(id string) (*BuildAction, error)
}

type BuildAction interface {
	Type() BuildActionType
	ResultingStatus() BuildStatus
	// State(), returning a terraform state object
	SourceCommitRef() *CommitRef
	DeployConfigTree() (*Tree, error)
	InfrastructureConfigTree() (*Tree, error)
	PreviousAction() (*BuildAction, error)
}

type BuildStatus int
type BuildActionType int
