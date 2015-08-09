package model

type Codebase interface {
	OriginURL() string
	LatestBuild() (*Build, error)
	LatestSuccessfulBuild() (*Build, error)
	AllBuildIds() ([]string, error)
	Build(id string) (*Build, error)
	Branch(name string) (*Branch, error)
	PrimaryBranch() (*Branch, error)
	DefaultUser() (*User, error)
	Commit(id string) (*Commit, error)
}
