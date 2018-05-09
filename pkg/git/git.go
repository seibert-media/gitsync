package git

import (
	"context"

	"gopkg.in/src-d/go-git.v4/plumbing"

	"gopkg.in/src-d/go-git.v4/plumbing/transport"

	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

// Syncer is responsible for fetching a git repository using go-git and then keeping it in sync when triggered
type Syncer struct {
	Repository *git.Repository
}

// New Syncer
func New(
	storer *memory.Storage, fs billy.Filesystem,
	url string, ref string, auth transport.AuthMethod,
) (*Syncer, error) {
	repo, err := git.Clone(storer, fs, &git.CloneOptions{
		URL:           url,
		Auth:          auth,
		ReferenceName: plumbing.ReferenceName(ref),
	})
	return &Syncer{
		Repository: repo,
	}, err
}

// Run sync to pull all recent changes for repository
func (s *Syncer) Run(ctx context.Context) error {
	tree, _ := s.Repository.Worktree()
	err := tree.PullContext(ctx, &git.PullOptions{})
	return err
}
