package main

import (
	"strings"
	"path/filepath"
	"os"
)

const rawUploadCommand string = "buildkite-signed-pipeline upload"

type UnsignedCommandValidator struct {}

func (u UnsignedCommandValidator) Allowed(command string) (bool, error) {
	if command == rawUploadCommand {
		return true, nil
	}

	uploadPrefix := rawUploadCommand + " "
	// If there's no additional arguments, bail early
	if !strings.HasPrefix(command, uploadPrefix) {
		return false, nil
	}

	fileArgument := strings.TrimPrefix(command, uploadPrefix)
	isLocal, err := isWorkingDirectoryFile(fileArgument)

	if err != nil {
		return false, err
	}

	return isLocal, nil
}

func isWorkingDirectoryFile(fileName string) (bool, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return false, err
	}
	pathToFile, err := filepath.Abs(filepath.Join(workingDirectory, fileName))
	return err == nil && fileExists(pathToFile), nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}