package git

import (
	"github.com/pkg/errors"
	"os/exec"
	"strings"
)

const (
	gitCommand    = "git"
	pushCommand   = "push"
	commitCommand = "commit"
	statusCommand = "status"
)

//git config --get remote.origin.url - get url for pushing

func IsConfigured(projectRoot string) (bool, error) {
	statusCommand := createCommand(projectRoot, statusCommand)
	_, err := statusCommand.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			exitMessage := string(ee.Stderr)
			if strings.Contains(exitMessage, "Not a git repository") {
				return false, nil
			}
		}
		return false, errors.Wrapf(err, "failed to determine if %s is configured with GIT", projectRoot)
	}
	return true, nil
}

func CommitAllChanges(projectRoot, commitMessage string) error {
	command := createCommand(projectRoot, commitCommand, "-a", "-m", commitMessage)
	_, err := command.Output()
	if err != nil {
		return errors.Wrapf(err, "failed to commit changes to %s with message %s", projectRoot, commitMessage)
	}
	return nil
}

func PushChanges(projectRoot string) error {
	command := createCommand(projectRoot, pushCommand)
	_, err := command.Output()
	if err != nil {
		return errors.Wrapf(err, "failed to push changes for %s", projectRoot)
	}
	return nil
}

func getPushURL() {

}

func createCommand(projectRoot, command string, additionalArgs ...string) *exec.Cmd {
	gitOperation := make([]string, len(additionalArgs)+1)
	gitOperation[0] = command
	for index, arg := range additionalArgs {
		gitOperation[index+1] = arg
	}
	cmd := exec.Command(gitCommand, gitOperation...)
	cmd.Dir = projectRoot
	return cmd
}
