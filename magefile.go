//+build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"os"
)

// Runs dep ensure and then installs the binary.
func Build() error {
	if err := sh.Run("go", "vet", "./..."); err != nil {
		return err
	}
	return sh.Run("go", "build", "-o", "hello-go")
}

func Deploy() error {
	for key, value := range os.Environ() {
		fmt.Printf("deploy: %s => %s\n", key, value)
	}
	if err := sh.Run("docker", "push", "amarantin/hello-go:latest"); err != nil {
		return err
	}
	return nil
}
