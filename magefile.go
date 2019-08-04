//+build mage

package main

import (
	"fmt"
	"os"

	_ "github.com/kastenpotential/mage-utils/gen/travis-ci"
	"github.com/magefile/mage/sh"
)

func getGitVersion() (string, error) {
	return sh.Output("git", "describe", "--abbrev=0", "--tags")
}

// Runs dep ensure and then installs the binary.
func BuildDockerContainer() error {
	version, _ := getGitVersion()
	if err := sh.Run("docker", "build", "-t", "amarantin/hello-go:latest", "."); err != nil {
		return err
	}
	if err := sh.Run("docker", "build", "-t", "amarantin/hello-go:arm32v7", "-f", "Dockerfile.arm32v7", "."); err != nil {
		return err
	}
	if version != "" {
		if err := sh.Run("docker", "tag", "amarantin/hello-go:latest", fmt.Sprintf("amarantin/hello-go:%s", version)); err != nil {
			return err
		}
		if err := sh.Run("docker", "tag", "amarantin/hello-go:arm32v7", fmt.Sprintf("amarantin/hello-go:%s-arm32v7", version)); err != nil {
			return err
		}
	}
	// if err := sh.Run("docker", "build", "-t", fmt.Sprintf("amarantin/hello-go:%s", version)); err != nil {
	// 	return err
	// }
	return nil
}

func DeployLatest() error {
	for key, value := range os.Environ() {
		fmt.Printf("deploy: %s => %s\n", key, value)
	}
	if err := sh.Run("docker", "push", "amarantin/hello-go:latest"); err != nil {
		return err
	}
	if err := sh.Run("docker", "push", "amarantin/hello-go:arm32v7"); err != nil {
		return err
	}
	return nil
}

func DeployVersion() error {
	for key, value := range os.Environ() {
		fmt.Printf("deploy: %s => %s\n", key, value)
	}
	if err := sh.Run("docker", "push", "amarantin/hello-go"); err != nil {
		return err
	}
	return nil
}
