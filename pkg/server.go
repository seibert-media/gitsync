package gitsync

import (
	"fmt"
	"net/http"

	"github.com/seibert-media/gitsync/pkg/git"
	"github.com/seibert-media/gitsync/pkg/handler"

	"github.com/playnet-public/libs/log"
	"go.uber.org/zap"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	githttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
)

// Server creates all required components and starts the http server
type Server struct {
	Log  *log.Logger
	Port int

	GitHost    string
	GitRef     string
	Username   string
	Password   string
	PrivateKey string

	Path string
}

// PrepareAndServe the handler
func (s *Server) PrepareAndServe() error {
	fs := osfs.New(s.Path)
	storer, _ := filesystem.NewStorage(fs)

	var auth transport.AuthMethod
	if s.PrivateKey != "" {
		auth, _ = ssh.NewPublicKeysFromFile(s.Username, s.PrivateKey, s.Password)
	} else {
		auth = &githttp.BasicAuth{Username: s.Username, Password: s.Password}
	}

	repository, _ := git.New(
		storer, fs, s.GitHost, s.GitRef, auth,
	)

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", s.Port),
		Handler: &handler.Syncer{
			Git:  repository,
			Hook: nil,
		},
	}
	s.Log.Info("listening", zap.Int("port", s.Port))
	return server.ListenAndServe()
}
