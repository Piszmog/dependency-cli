package util

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"time"
)

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
