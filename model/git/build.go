package git

import (
	"github.com/fogforge/fogforge/model"

	git2go "github.com/libgit2/git2go"
)

type Build struct {
	codebase *Codebase
	ref      *git2go.Reference
}

type BuildAction struct {
}

func (c *Codebase) BuildForGitReference(ref *git2go.Reference) (model.Build, error) {
	return &Build{
		codebase: c,
		ref:      ref,
	}, nil
}

func (b *Build) LatestAction() (model.BuildAction, error) {
	return nil, nil
}
