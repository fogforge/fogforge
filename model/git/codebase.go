package git

import (
	"github.com/fogforge/fogforge/model"

	git2go "github.com/libgit2/git2go"
)

type Codebase struct {
	Repository *git2go.Repository
	OriginURL  string
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
		Repository: repo,
		OriginURL:  originURL,
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
		Repository: repo,
		OriginURL:  originURL,
	}, nil
}
