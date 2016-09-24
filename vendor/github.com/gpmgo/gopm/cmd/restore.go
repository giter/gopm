// Copyright 2014 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package cmd

import (
	"github.com/gpmgo/gopm/modules/base"
	"github.com/gpmgo/gopm/modules/cli"
	"github.com/gpmgo/gopm/modules/doc"
	"github.com/gpmgo/gopm/modules/errors"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/gpmgo/gopm/modules/setting"
)

var CmdRestore = cli.Command{
	Name:   "restore",
	Usage:  "restore remote package(s) to $GOPATH",
	Action: runRestore,
}

var (
	restoreCount     int
	failRestoreConut int
)

func restorePackages(target string, ctx *cli.Context, nodes []*doc.Node) error {

	var err error

	for _, n := range nodes {

		// Check if it is a valid remote path or C.
		if n.ImportPath == "C" {
			continue
		} else if !base.IsValidRemotePath(n.ImportPath) {
			// Invalid import path.
			if setting.LibraryMode {
				errors.AppendError(errors.NewErrInvalidPackage(n.VerString()))
			}
			log.Error("Skipped invalid package: " + n.VerString())
			failConut++
			continue
		}

		// Valid import path.
		if isSubpackage(n.RootPath, target) {
			continue
		}

		// Indicates whether need to download package or update.
		if n.IsFixed() && n.IsExist() {
			n.IsGetDepsOnly = true
		}

		if n.IsExist() && n.ValString() != "<UTD>" {

			err = n.ForceCopyToGopath()
			restoreCount++

			log.Info("Copied %s:%s...", n.ImportPath, n.ValString())

			if err != nil {
				return err
			}
		} else {
			log.Info("Skipped uncached package: %s", n.VerString())
			failRestoreConut++
			continue
		}
	}

	log.Info("%d package(s) downloaded, %d failed", restoreCount, failRestoreConut)

	return nil
}

func restoreByGopmfile(ctx *cli.Context) error {

	// Make sure gopmfile exists and up-to-date.

	gf, target, err := parseGopmfile(setting.GOPMFILE)

	if err != nil {
		return err
	}

	imports, err := getDepList(ctx, target, setting.WorkDir, setting.DefaultVendor)

	if err != nil {
		return err
	}

	// Check if dependency has version.
	nodes := make([]*doc.Node, 0, len(imports))

	for _, name := range imports {
		name = doc.GetRootPath(name)
		n := doc.NewNode(name, doc.BRANCH, "", !ctx.Bool("download"))

		// Check if user specified the version.
		if v := gf.MustValue("deps", name); len(v) > 0 {

			n.Type, n.Value, err = validPkgInfo(v)
			n = doc.NewNode(name, n.Type, n.Value, !ctx.Bool("download"))
		}

		nodes = append(nodes, n)
	}

	return restorePackages(target, ctx, nodes)
}

func runRestore(ctx *cli.Context) {

	var err error

	if err = setup(ctx); err != nil {
		errors.SetError(err)
		return
	}

	err = restoreByGopmfile(ctx)

	if err != nil {
		errors.SetError(err)
		return
	}
}
