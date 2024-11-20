// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package commands

import (
	"os/exec"

	"github.com/onsi/gomega"
	"github.com/shubhamdubey02/Cryft-cli/pkg/constants"
)

func GetVersion() string {
	/* #nosec G204 */
	cmd := exec.Command(
		CLIBinary,
		"--version",
		"--"+constants.SkipUpdateFlag,
	)
	output, err := cmd.Output()
	gomega.Expect(err).Should(gomega.BeNil())
	return string(output)
}
