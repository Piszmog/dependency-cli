package main

import (
	"flag"
	"github.com/Piszmog/dependency-cli/maven"
	"github.com/Piszmog/dependency-cli/util"
	"github.com/pkg/errors"
	"time"
)

func main() {
	defer util.Runtime(time.Now())
	d := flag.String("f", "", "the file containing list of projects to update")
	flag.Parse()
	if len(*d) == 0 {
		panic(errors.New("requires a file to run"))
	}
	file, err := util.OpenFile(*d)
	if err != nil {
		panic(err)
	}
	defer util.CloseFile(file)
	lineChannel := make(chan string, 100)
	done := make(chan bool)
	go read(lineChannel, done)
	util.ReadTXTFile(file, lineChannel)
	<-done
}

func read(lineChannel chan string, done chan bool) {
	for line := range lineChannel {
		err := maven.UpdateProject(line)
		if err != nil {
			panic(err)
		}
	}
	done <- true
}
