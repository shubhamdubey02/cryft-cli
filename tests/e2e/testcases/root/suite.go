// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package root

import (
	"fmt"
	"os"

	ginkgo "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/shubhamdubey02/cryft-cli/tests/e2e/commands"
)

var _ = ginkgo.Describe("[Root]", func() {
	ginkgo.It("can print version", func() {
		expectedVersion, err := os.ReadFile("VERSION")
		expectedVersionStr := fmt.Sprintf("metal version %s\n", string(expectedVersion))
		gomega.Expect(err).Should(gomega.BeNil())

		version := commands.GetVersion()

		gomega.Expect(version).Should(gomega.Equal(expectedVersionStr))
		gomega.Expect(err).Should(gomega.BeNil())
	})
})
