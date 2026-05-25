// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

package main

import (
	"log"

	plugin "github.com/SemRels/updater-terraform/internal/plugin"
)

func main() {
	generator := plugin.NewManifestGenerator(plugin.ManifestConfig{})
	log.Printf("updater-terraform plugin ready: generates Terraform provider manifests (%T)", generator)
}
