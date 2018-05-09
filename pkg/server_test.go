// Copyright 2018 The gitsync authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitsync_test

import (
	"testing"

	"github.com/seibert-media/gitsync/pkg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/playnet-public/libs/log"
)

var _ = Describe("gitsync", func() {
	var (
		l *log.Logger
		s *gitsync.Server
	)

	BeforeEach(func() {
		l = log.NewNop()
		s = &gitsync.Server{
			Log: l,
		}
	})

	Describe("Prepare", func() {
		It("does set non nil http server", func() {
			s.Prepare()
			Expect(s.Server).NotTo(BeNil())
		})
	})
})

func TestGitSync(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GitSync Test Suite")
}
