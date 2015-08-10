package model

type Codebase interface {
	OriginURL() string
	LatestBuild() (Build, error)
	LatestSuccessfulBuild() (Build, error)
	AllBuildIDs() ([]string, error)
	Build(id string) (Build, error)
	Branch(name string) (Branch, error)
	PrimaryBranch() (Branch, error)
	DefaultUser() (User, error)
	Commit(id string) (Commit, error)
	CommitFuzzy(commitish string) (Commit, error)
}
