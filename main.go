package main

import (
	"encoding/json"
	"flag"
	"github.com/Piszmog/dependency-cli/maven"
	"github.com/Piszmog/dependency-cli/util"
	"github.com/pkg/errors"
	"io/ioutil"
	"path"
	"sync"
	"time"
)

type ConfigurationFile struct {
	MavenProjects []MavenProject `json:"mavenProjects"`
}

type MavenProject struct {
	BaseDirectory string   `json:"baseDirectory"`
	Projects      []string `json:"projects"`
}

func main() {
	defer util.Runtime(time.Now())
	d := flag.String("f", "", "the file containing list of projects to update")
	flag.Parse()
	if len(*d) == 0 {
		panic(errors.New("requires a file to run"))
	}
	configFile, err := readConfigFile(*d)
	if err != nil {
		panic(err)
	}
	handleConfigFile(configFile)
}

func readConfigFile(configFile string) (*ConfigurationFile, error) {
	file, err := util.OpenFile(configFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open file %s", configFile)
	}
	defer util.CloseFile(file)
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file %s", configFile)
	}
	var configurationFile *ConfigurationFile
	err = json.Unmarshal(bytes, &configurationFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal file %s", configFile)
	}
	return configurationFile, err
}

func handleConfigFile(configFile *ConfigurationFile) {
	var waitGroup sync.WaitGroup
	for _, mavenProject := range configFile.MavenProjects {
		baseDir := mavenProject.BaseDirectory
		for _, project := range mavenProject.Projects {
			waitGroup.Add(1)
			go updateMavenProject(baseDir, project, &waitGroup)
		}
	}
	waitGroup.Wait()
}

func updateMavenProject(baseDir string, project string, waitGroup *sync.WaitGroup) {
	pathToProject := path.Join(baseDir, project)
	err := maven.UpdateProject(pathToProject)
	if err != nil {
		panic(err)
	}
	waitGroup.Done()
}
