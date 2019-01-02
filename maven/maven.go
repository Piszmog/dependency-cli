package maven

import (
	"fmt"
	"github.com/pkg/errors"
	"os/exec"
	"strings"
)

const (
	mvnCommand                = "mvn"
	updateDependenciesCommand = "versions:use-latest-releases"
	updateParentCommand       = "versions:update-parent"
	Includes                  = "-Dincludes"
	Excludes                  = "-Dexcludes"
)

func UpdateProject(projectRoot string, includes, excludes []string, updateDependencies, updateParent bool) error {
	includesList, excludesList := createAdditionalCommands(includes, excludes)
	var err error
	if updateDependencies {
		err = updateProjectDependencies(projectRoot, includesList, excludesList)
		if err != nil {
			return errors.Wrapf(err, "failed to update project %s", projectRoot)
		}
	}
	if updateParent {
		err = updateProjectParent(projectRoot, includesList, excludesList)
		if err != nil {
			return errors.Wrapf(err, "failed to update project %s", projectRoot)
		}
	}
	return nil
}

func createAdditionalCommands(includes, excludes []string) (string, string) {
	includesList := ""
	excludesList := ""
	if len(includes) > 0 {
		includesList = Includes + "=" + strings.Join(includes, ",")
	}
	if len(excludes) > 0 {
		excludesList = Excludes + "=" + strings.Join(excludes, ",")
	}
	return includesList, excludesList
}

func updateProjectDependencies(projectRoot, includesList, excludesList string) error {
	updateDependenciesCommand := createCommand(projectRoot, updateDependenciesCommand, includesList, excludesList)
	fmt.Printf("Updating dependency versions for project %s...\n", projectRoot)
	_, err := updateDependenciesCommand.Output()
	if err != nil {
		return errors.Wrapf(err, "failed to update dependencies for project %s", projectRoot)
	}
	fmt.Printf("Successfully updated dependency versions for project %s\n", projectRoot)
	return nil
}

func updateProjectParent(projectRoot, includesList, excludesList string) error {
	updateParentCommand := createCommand(projectRoot, updateParentCommand, includesList, excludesList)
	fmt.Printf("Updating parent version for project %s...\n", projectRoot)
	_, err := updateParentCommand.Output()
	if err != nil {
		return errors.Wrapf(err, "failed to update parent for project %s", projectRoot)
	}
	fmt.Printf("Successfully updated parent version for project %s\n", projectRoot)
	return nil
}

func createCommand(projectRoot, mavenCommand, includesList, excludesList string) *exec.Cmd {
	var mavenOperation []string
	mavenOperation = append(mavenOperation, mavenCommand)
	if len(includesList) != 0 {
		mavenOperation = append(mavenOperation, includesList)
	}
	if len(excludesList) != 0 {
		mavenOperation = append(mavenOperation, excludesList)
	}
	command := exec.Command(mvnCommand, mavenOperation...)
	command.Dir = projectRoot
	return command
}
