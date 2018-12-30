package maven

import (
	"fmt"
	"github.com/pkg/errors"
	"os/exec"
)

const (
	mvnCommand                = "mvn"
	updateDependenciesCommand = "versions:use-latest-releases"
	updateParentCommand       = "versions:update-parent"
)

func UpdateProject(projectRoot string) error {
	err := updateDependencies(projectRoot)
	if err != nil {
		return errors.Wrapf(err, "failed to update project %s", projectRoot)
	}
	err = updateParent(projectRoot)
	if err != nil {
		return errors.Wrapf(err, "failed to update project %s", projectRoot)
	}
	return nil
}

func updateDependencies(projectRoot string) error {
	updateDependenciesCommand := createMavenCommand(projectRoot, updateDependenciesCommand)
	fmt.Printf("Updating dependency versions for project %s...\n", projectRoot)
	_, err := updateDependenciesCommand.Output()
	if err != nil {
		return errors.Wrapf(err, "failed to update dependencies for project %s", projectRoot)
	}
	fmt.Printf("Successfully updated dependency versions for project %s\n", projectRoot)
	return nil
}

func updateParent(projectRoot string) error {
	updateParentCommand := createMavenCommand(projectRoot, updateParentCommand)
	fmt.Printf("Updating parent version for project %s...\n", projectRoot)
	_, err := updateParentCommand.Output()
	if err != nil {
		return errors.Wrapf(err, "failed to update parent for project %s", projectRoot)
	}
	fmt.Printf("Successfully updated parent version for project %s\n", projectRoot)
	return nil
}

func createMavenCommand(projectRoot, mavenCommand string) *exec.Cmd {
	command := exec.Command(mvnCommand, mavenCommand)
	command.Dir = projectRoot
	return command
}
