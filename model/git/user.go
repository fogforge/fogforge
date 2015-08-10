package git

import (
	"github.com/fogforge/fogforge/model"

	git2go "github.com/libgit2/git2go"
)

type User struct {
	sig *git2go.Signature
}

func UserForGitSignature(sig *git2go.Signature) model.User {
	return &User{
		sig: sig,
	}
}

func (u *User) Name() string {
	return u.sig.Name
}

func (u *User) Email() string {
	return u.sig.Email
}
