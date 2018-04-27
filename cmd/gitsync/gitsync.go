// Copyright 2018 The gitsync authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"runtime"

	gitsync "github.com/seibert-media/gitsync/pkg"

	flag "github.com/bborbe/flagenv"
	"github.com/golang/glog"
	"github.com/kolide/kit/version"
	"github.com/playnet-public/libs/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	appName = "gitsync"
	appKey  = "gitsync"
)

var (
	debug       = flag.Bool("debug", false, "show debug info")
	showVersion = flag.Bool("version", false, "show version info")
	sentryDsn   = flag.String("sentryDsn", "", "sentry dsn key")
	port        = flag.Int("port", 8080, "port for the webhook to listen on")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	if *showVersion {
		v := version.Version()
		fmt.Printf("-- //S/M %s --\n", appName)
		fmt.Printf(" - version: %s\n", v.Version)
		fmt.Printf("   branch: \t%s\n", v.Branch)
		fmt.Printf("   revision: \t%s\n", v.Revision)
		fmt.Printf("   build date: \t%s\n", v.BuildDate)
		fmt.Printf("   build user: \t%s\n", v.BuildUser)
		fmt.Printf("   go version: \t%s\n", v.GoVersion)
		os.Exit(0)
	}

	var zapFields []zapcore.Field
	if !*debug {
		zapFields = []zapcore.Field{
			zap.String("app", appKey),
			zap.String("version", version.Version().Version),
		}
	}

	log := log.New(appKey, *sentryDsn, *debug).WithFields(zapFields...)
	defer log.Sync()
	log.Info("preparing")

	if err := do(log); err != nil {
		log.Fatal("failed", zap.Error(err))
	}
}

func do(log *log.Logger) error {
	srv := gitsync.Server{
		Port: *port,
	}
	return srv.PrepareAndServe()
}
