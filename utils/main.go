package utils

import (
	"fmt"
	"io/ioutil"
)

// ReadFile reads in contents of a file
func ReadFile(filename string) (string, error) {
	read, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Error reading file: %s", err)
	}
	return string(read), nil
}
