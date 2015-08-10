package model

type Build interface {
	LatestAction() (BuildAction, error)
}

type BuildAction interface {
	Type() BuildActionType
	ResultingStatus() BuildStatus
	// State(), returning a terraform state object
	SourceCommitRef() CommitRef
	DeployConfigTreeId() (string, error)
	InfrastructureConfigTreeId() (string, error)
}

type BuildStatus string
type BuildActionType string

const (
	BuildStatusBuilt     BuildStatus = "BUILT"
	BuiltStatusErrored   BuildStatus = "ERRORED"
	BuildStatusDestroyed BuildStatus = "DESTROYED"
	BuildStatusUnknown   BuildStatus = ""
)

const (
	BuildActionInit    BuildActionType = "INIT"
	BuildActionBuild   BuildActionType = "BUILD"
	BuildActionDestroy BuildActionType = "DESTROY"
	BuildActionUnknown BuildActionType = ""
)
