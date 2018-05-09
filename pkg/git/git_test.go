// Copyright 2018 The gitsync authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package git_test

import (
	"context"
	"testing"

	"github.com/seibert-media/gitsync/pkg/git"

	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/storage/memory"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("git", func() {
	var (
		fs     billy.Filesystem
		storer *memory.Storage
		s      *git.Repository
		ctx    context.Context
		auth   transport.AuthMethod
	)

	BeforeEach(func() {
		fs = memfs.New()
		storer = memory.NewStorage()
		auth = &http.BasicAuth{Username: "", Password: ""}
		s, _ = git.New(storer, fs, "", "HEAD", auth)
		ctx = context.Background()
	})

	Describe("New", func() {
		It("does not return nil", func() {
			s, _ := git.New(storer, fs, "", "HEAD", auth)
			Expect(s).NotTo(BeNil())
		})
		It("does return error on failed clone", func() {
			_, err := git.New(storer, fs, "", "HEAD", auth)
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("Sync", func() {
		It("does return error on failed pull", func() {
			ctx = nil
			Expect(s.Sync(ctx)).NotTo(BeNil())
		})
	})
})

func TestGitSync(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GitSync Test Suite")
}
