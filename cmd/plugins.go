// Copyright 2018 The OPA Authors.  All rights reserved.
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

// Specifies additional cmd commands that available to systems that can load plugins
// +build linux,cgo darwin,cgo

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/open-policy-agent/opa/runtime"
)

func init() {
	var builtinDir string

	// flag is persistent (can be loaded on all children commands)
	RootCommand.PersistentFlags().StringVarP(&builtinDir, "builtin-dir", "b", "", `set path to directory from which to load builtins`)

	// Runs before *all* children commands
	RootCommand.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if builtinDir != "" {
			err := runtime.RegisterBuiltinsFromDir(builtinDir)
			if err != nil {
				return err
			}
		}
		return nil
	}
}
