// +build mage

package main

import (
  "github.com/carolynvs/magex/pkg"
  "github.com/magefile/mage/mg"
  "github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Build

// Compile and lint the cli
func Build() error {
  mg.Deps(Lint)

  return sh.RunV("go", "build", "./...")
}

// nit the hell outta my code
func Lint() error {
  mg.Deps(EnsureGoLint)

  return sh.RunV("golint", "./...")
}

// Install golint if needed
func EnsureGoLint() error {
  return pkg.EnsurePackage("golang.org/x/lint/golint", "")
}

// Install Mage if necessary
func EnsureMage() error {
  return pkg.EnsureMage("v1.11.0")
}