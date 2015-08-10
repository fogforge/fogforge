package model

import (
	"time"
)

type CommitRef interface {
	CodebaseURL() string
	CommitID() string
}

type Commit interface {
	Author() User
	AuthorTime() time.Time
	Committer() User
	CommitTime() time.Time
	Message() string
	MessageSubject() string
	MessageBody() string
	Tree() (Tree, error)
}
