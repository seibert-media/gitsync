package gitsync

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

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

	syncer := &handler.Syncer{
		Git:  repository,
		Hook: nil,
	}

	ctx := context.Background()
	m := mux.NewRouter()
	m.PathPrefix("/").HandlerFunc(withContext(ctx, syncer.ServeHTTP))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: m,
	}
	s.Log.Info("listening", zap.Int("port", s.Port))
	return server.ListenAndServe()
}

func withContext(ctx context.Context, handleFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleFunc(ctx, w, r)
	}
}
