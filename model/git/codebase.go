package git

import (
	"strings"

	"github.com/fogforge/fogforge/model"

	git2go "github.com/libgit2/git2go"
)

var BuildRefPrefix = "refs/builds/"

type Codebase struct {
	repo      *git2go.Repository
	originURL string
}

func CreateCodebase(localPath string, originURL string) (model.Codebase, error) {
	repo, err := git2go.InitRepository(localPath, true)
	if err != nil {
		return nil, err
	}

	if originURL != "" {
		config, err := repo.Config()
		if err != nil {
			return nil, err
		}

		err = config.SetString("remote.origin.url", originURL)
		if err != nil {
			return nil, err
		}
	}

	return &Codebase{
		repo:      repo,
		originURL: originURL,
	}, nil
}

func OpenCodebase(localPath string) (model.Codebase, error) {
	repo, err := git2go.OpenRepository(localPath)
	if err != nil {
		return nil, err
	}

	config, err := repo.Config()
	if err != nil {
		return nil, err
	}

	originURL, err := config.LookupString("remote.origin.url")
	if err == nil {
		originURL = ""
	}

	return &Codebase{
		repo:      repo,
		originURL: originURL,
	}, nil
}

func (c *Codebase) OriginURL() string {
	return c.originURL
}

func (c *Codebase) LatestBuild() (model.Build, error) {
	return nil, nil
}

func (c *Codebase) LatestSuccessfulBuild() (model.Build, error) {
	return nil, nil
}

func (c *Codebase) AllBuildIDs() ([]string, error) {
	iter, err := c.repo.NewReferenceNameIterator()
	if err != nil {
		return nil, err
	}

	ret := []string{}
	name, err := iter.Next()
	for err == nil {
		if strings.HasPrefix(name, BuildRefPrefix) {
			ret = append(ret, name[len(BuildRefPrefix):])
		}
		name, err = iter.Next()
	}

	if gitErr, ok := err.(*git2go.GitError); ok {
		if gitErr.Code == git2go.ErrIterOver {
			err = nil
		}
	}

	return ret, err
}

func (c *Codebase) Build(id string) (model.Build, error) {
	refName := BuildRefPrefix + id
	ref, err := c.repo.References.Lookup(refName)
	if err != nil {
		return nil, err
	}

	return c.BuildForGitReference(ref)
}

func (c *Codebase) PrimaryBranch() (model.Branch, error) {
	return nil, nil
}

func (c *Codebase) Branch(name string) (model.Branch, error) {
	return nil, nil
}

func (c *Codebase) DefaultUser() (model.User, error) {
	sig, err := c.repo.DefaultSignature()
	if err != nil {
		return nil, err
	}

	return UserForGitSignature(sig), nil
}

func (c *Codebase) Commit(id string) (model.Commit, error) {
	return nil, nil
}

func (c *Codebase) CommitFuzzy(commitish string) (model.Commit, error) {
	return nil, nil
}
