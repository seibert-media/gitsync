// Copyright 2018 The gitsync authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitsync_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"testing"
	"time"

	"gopkg.in/src-d/go-billy.v4/osfs"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"

	"github.com/seibert-media/gitsync/pkg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/playnet-public/libs/log"
)

var _ = Describe("gitsync", func() {
	var (
		l          *log.Logger
		s          *gitsync.Server
		tmp1, tmp2 string
	)

	BeforeEach(func() {
		l = log.NewNop()
		s = &gitsync.Server{
			Log: l,
		}
	})

	Describe("Prepare", func() {
		BeforeEach(func() {
			tmp1, _ = ioutil.TempDir(os.TempDir(), "gitsync_test")
			tmp2, _ = ioutil.TempDir(os.TempDir(), "gitsync_test")

			fs := osfs.New(tmp1)
			st, _ := filesystem.NewStorage(fs)
			r1, _ := git.Init(st, fs)
			ioutil.WriteFile(path.Join(tmp1, "file"), []byte("foobar"), 777)
			wt, _ := r1.Worktree()
			wt.Add("file")
			_, err := wt.Commit("commit", &git.CommitOptions{All: true,
				Author: &object.Signature{
					Name:  "test",
					Email: "test@test.test",
					When:  time.Now(),
				}})
			Expect(err).To(BeNil())

			s.GitHost = tmp1
			s.Path = tmp2
		})
		AfterEach(func() {
			os.Remove(tmp1)
			os.Remove(tmp2)
		})
		It("does not return error", func() {
			Expect(s.Prepare()).To(BeNil())
		})
		It("does set non nil http server", func() {
			s.Prepare()
			Expect(s.Server).NotTo(BeNil())
		})
		It("does return error on invalid git", func() {
			s.GitHost = "."
			Expect(s.Prepare()).NotTo(BeNil())
		})
	})

	Describe("WithContext", func() {
		It("does pass context correctly", func() {
			ctx := context.Background()
			ctx, c := context.WithCancel(ctx)
			c()
			h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
				Expect(ctx.Err()).NotTo(BeNil())
			}
			gitsync.WithContext(ctx, h).ServeHTTP(nil, nil)
		})
	})
})

func TestGitSync(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GitSync Test Suite")
}
