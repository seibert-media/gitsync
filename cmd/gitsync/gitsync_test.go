// Copyright 2018 The gitsync authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"fmt"
	"net/http"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/ghttp"
)

var pathToServerBinary string
var serverSession *gexec.Session
var server *ghttp.Server

var _ = BeforeSuite(func() {
	var err error
	pathToServerBinary, err = gexec.Build("github.com/seibert-media/gitsync/cmd/gitsync")
	Expect(err).NotTo(HaveOccurred())
})

var _ = BeforeEach(func() {
	server = ghttp.NewServer()
	server.RouteToHandler(http.MethodGet, "/", ghttp.RespondWithJSONEncoded(http.StatusOK, []string{"a.example.com", "b.example.com"}))
})

var _ = AfterEach(func() {
	serverSession.Interrupt()
	Eventually(serverSession).Should(gexec.Exit())
	server.Close()
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

type args map[string]string

func (a args) list() []string {
	var result []string
	for k, v := range a {
		if len(v) == 0 {
			result = append(result, fmt.Sprintf("-%s", k))
		} else {
			result = append(result, fmt.Sprintf("-%s=%s", k, v))
		}
	}
	return result
}

var validargs args

var _ = BeforeEach(func() {
	validargs = map[string]string{
		"debug":     "",
		"version":   "0",
		"sentryDsn": "",
	}
})

var _ = Describe("gitsync", func() {
	var err error
	Describe("when asked for version", func() {
		It("prints version string", func() {
			serverSession, err = gexec.Start(exec.Command(pathToServerBinary, "-version"), GinkgoWriter, GinkgoWriter)
			Expect(err).To(BeNil())
			serverSession.Wait(time.Second)
			Expect(serverSession.ExitCode()).To(Equal(0))
			Expect(serverSession.Out).To(gbytes.Say(`-- //S/M gitsync --
 - version: unknown
   branch: 	unknown
   revision: 	unknown
   build date: 	unknown
   build user: 	unknown
   go version: 	unknown
`))
		})
	})
	Describe("when validating parameters", func() {

	})
	Describe("when called with valid input", func() {

	})

	Context("when given parameters via environment", func() {
		Describe("when no arguments are given via command line", func() {
			BeforeEach(func() {
				validargs = nil
			})
			It("uses version environment variable", func() {
				cmd := exec.Command(pathToServerBinary, validargs.list()...)
				cmd.Env = []string{"VERSION=true"}
				serverSession, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
				Expect(err).To(BeNil())
				serverSession.Wait(time.Second)
				Expect(serverSession.ExitCode()).To(Equal(0))
			})
		})
		Describe("when version is set via command line", func() {
			BeforeEach(func() {
				validargs = map[string]string{
					"version": "true",
				}
			})
			It("uses command line argument value prioritized over environment", func() {
				cmd := exec.Command(pathToServerBinary, validargs.list()...)
				cmd.Env = []string{"VERSION=false"}
				serverSession, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
				Expect(err).To(BeNil())
				serverSession.Wait(time.Second)
				Expect(serverSession.ExitCode()).To(Equal(0))
			})
		})
	})
})

func TestSystem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "System Test Suite")
}
