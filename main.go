package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	mvnCommand                = "mvn"
	updateDependenciesCommand = "versions:use-latest-releases"
	updateParentCommand       = "versions:update-parent"
)

func main() {
	defer Runtime(time.Now())
	d := flag.String("f", "", "the file containing list of projects to update")
	flag.Parse()
	if len(*d) == 0 {
		panic(errors.New("requires a file to run"))
	}
	file, err := OpenFile(*d)
	if err != nil {
		panic(err)
	}
	defer CloseFile(file)
	lineChannel := make(chan string, 100)
	done := make(chan bool)
	go read(lineChannel, done)
	ReadTXTFile(file, lineChannel)
	<-done
}

func read(lineChannel chan string, done chan bool) {
	for line := range lineChannel {
		err := UpdateProject(line)
		if err != nil {
			panic(err)
		}
	}
	done <- true
}

func Runtime(now time.Time) {
	fmt.Printf("\nApplication took %f seconds to complete\n", time.Since(now).Seconds())
}

// Opens the specified file
func OpenFile(filename string) (*os.File, error) {
	pathToFile, err := filepath.Abs(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get absolute path of %s", filename)
	}
	file, err := os.Open(pathToFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open file %s", filename)
	}
	return file, nil
}

// Closes the file and panics if an error occurs
func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(errors.Wrapf(err, "failed to close %s", file.Name()))
	}
}

func ReadTXTFile(file *os.File, lines chan string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines <- scanner.Text()
	}
	close(lines)
}

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
