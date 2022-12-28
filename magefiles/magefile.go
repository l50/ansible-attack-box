//go:build mage

package main

import (
	"fmt"
	"os"

	"github.com/bitfield/script"
	"github.com/fatih/color"
	goutils "github.com/l50/goutils"

	// mage utility functions
	"github.com/magefile/mage/mg"
)

func init() {
	os.Setenv("GO111MODULE", "on")
}

// installDeps Installs project dependencies
func installDeps() error {
	fmt.Println(color.YellowString("Installing dependencies."))
	goutils.Cd("magefiles")

	if err := goutils.Tidy(); err != nil {
		return fmt.Errorf(color.RedString(
			"failed to install dependencies: %v", err))
	}

	return nil
}

// InstallPreCommitHooks Installs pre-commit hooks locally
func InstallPreCommitHooks() error {
	mg.Deps(installDeps)

	fmt.Println(color.YellowString("Installing pre-commit hooks."))
	if err := goutils.InstallPCHooks(); err != nil {
		return err
	}

	return nil
}

// RunPreCommit runs all pre-commit hooks locally
func RunPreCommit() error {
	// mg.Deps(installDeps)

	fmt.Println(color.YellowString("Updating pre-commit hooks."))
	if err := goutils.UpdatePCHooks(); err != nil {
		return err
	}

	fmt.Println(color.YellowString(
		"Clearing the pre-commit cache to ensure we have a fresh start."))
	if err := goutils.ClearPCCache(); err != nil {
		return err
	}

	fmt.Println(color.YellowString("Running all pre-commit hooks locally."))
	if err := goutils.RunPCHooks(); err != nil {
		return err
	}

	return nil
}

func runCmds(cmds []string) error {
	for _, cmd := range cmds {
		if _, err := script.Exec(cmd).Stdout(); err != nil {
			return err
		}
	}

	return nil

}

// LintAnsible runs ansible-lint.
func LintAnsible() error {
	cmds := []string{
		"ansible-lint --force-color -c .hooks/linters/.ansible-lint",
	}

	fmt.Println(color.YellowString("Running ansible-lint."))
	if err := runCmds(cmds); err != nil {
		return fmt.Errorf(color.RedString("failed to run ansible-lint: %v", err))
	}

	return nil
}

// RunMoleculeTests runs the molecule tests.
func RunMoleculeTests() error {
	cmds := []string{
		"molecule create",
		"molecule converge",
		"molecule idempotence",
		"molecule destroy",
	}

	fmt.Println(color.YellowString("Running molecule tests."))
	if err := runCmds(cmds); err != nil {
		return fmt.Errorf(color.RedString("failed to run molecule tests: %v", err))
	}

	return nil
}
