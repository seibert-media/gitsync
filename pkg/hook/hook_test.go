// Copyright 2018 The gitsync authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hook_test

import (
	"context"
	"testing"

	"github.com/seibert-media/gitsync/pkg/hook"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("hook", func() {
	var (
		h   *hook.Hook
		ctx context.Context
	)

	BeforeEach(func() {
		h = &hook.Hook{}
		ctx = context.Background()
	})

	Describe("New", func() {
		It("does not return nil", func() {
			Expect(hook.New("")).NotTo(BeNil())
		})
	})

	Describe("Call", func() {
		It("does not return error", func() {
			Expect(h.Call(ctx)).To(BeNil())
		})
	})
})

func TestGitSync(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hook Test Suite")
}
