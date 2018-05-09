package gitsync

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/playnet-public/libs/log"
	"go.uber.org/zap"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	githttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"

	"github.com/seibert-media/gitsync/pkg/git"
	"github.com/seibert-media/gitsync/pkg/handler"
)

// Server creates all required components and starts the http server
type Server struct {
	*http.Server

	Log  *log.Logger
	Port int

	GitHost    string
	GitRef     string
	Username   string
	Password   string
	PrivateKey string

	Path string
}

// Prepare the server
func (s *Server) Prepare() error {
	fs := osfs.New(s.Path)
	storer, err := filesystem.NewStorage(fs)
	if err != nil {
		return errors.Wrap(err, "create filesystem")
	}

	var auth transport.AuthMethod
	if s.PrivateKey != "" {
		auth, err = ssh.NewPublicKeysFromFile(s.Username, s.PrivateKey, s.Password)
		if err != nil {
			return errors.Wrap(err, "create public key")
		}
	} else {
		auth = &githttp.BasicAuth{Username: s.Username, Password: s.Password}
	}

	repository, err := git.New(storer, fs, s.GitHost, s.GitRef, auth)
	if err != nil {
		return errors.Wrap(err, "initialize git")
	}

	syncer := &handler.Syncer{
		Git:  repository,
		Hook: nil,
	}

	ctx := context.Background()
	m := mux.NewRouter()
	m.PathPrefix("/").HandlerFunc(WithContext(ctx, syncer.ServeHTTP))

	s.Server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: m,
	}

	return nil
}

// WithContext wraps contextual handleFuncs into default ones for use with mux
func WithContext(ctx context.Context, handleFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleFunc(ctx, w, r)
	}
}

// Serve via http
func (s *Server) Serve() error {
	go func() {
		s.Log.Info("listening", zap.Int("port", s.Port))
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			s.Log.Fatal("fatal server error", zap.Error(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	s.Log.Info("graceful shutdown", zap.Duration("timeout", 1*time.Second))

	if err := s.Shutdown(ctx); err != nil {
		s.Log.Error("graceful shutdown", zap.Error(err))
		return err
	}

	s.Log.Info("stopped listening")
	return nil
}
