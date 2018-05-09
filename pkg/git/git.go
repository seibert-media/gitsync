package git

import (
	"context"

	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/storage"
)

// Repository is responsible for fetching a git repository using go-git and then keeping it in sync when triggered
type Repository struct {
	*git.Repository
}

// New Repository
func New(
	storer storage.Storer, fs billy.Filesystem,
	url string, ref string, auth transport.AuthMethod,
) (*Repository, error) {
	repo, err := git.Clone(storer, fs, &git.CloneOptions{
		URL:           url,
		Auth:          auth,
		ReferenceName: plumbing.ReferenceName(ref),
	})
	return &Repository{
		Repository: repo,
	}, err
}

// Sync pulls all recent changes from the remote repository
func (s *Repository) Sync(ctx context.Context) error {
	tree, _ := s.Repository.Worktree()
	err := tree.PullContext(ctx, &git.PullOptions{})
	return err
}
