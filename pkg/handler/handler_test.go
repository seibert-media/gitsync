// Copyright 2018 The gitsync authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/seibert-media/gitsync/pkg/handler"
	"github.com/seibert-media/gitsync/pkg/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("gitsync", func() {
	var (
		git    *mocks.Git
		hook   *mocks.Hook
		syncer *handler.Syncer
	)

	BeforeEach(func() {
		git = &mocks.Git{}
		hook = &mocks.Hook{}
		syncer = &handler.Syncer{
			Git:  git,
			Hook: hook,
		}
	})

	It("return status code 200", func() {
		recorder := httptest.NewRecorder()
		syncer.ServeHTTP(recorder, &http.Request{})
		Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
	})
	It("write empty json on success", func() {
		recorder := httptest.NewRecorder()
		syncer.ServeHTTP(recorder, &http.Request{})
		content, _ := ioutil.ReadAll(recorder.Result().Body)
		Expect(gbytes.BufferWithBytes(content)).To(gbytes.Say("{}"))
	})
})

func TestGitSync(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GitSync Test Suite")
}
