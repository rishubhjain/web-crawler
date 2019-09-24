package tests

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// readFromFile reads the entire file into a string given by the path parameter
func readFromFile(path string) (string, error) {
	input := ""
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		input = input + line
	}

	return input, nil
}

// ParseHTML parses HTML file in the given path and returns a string
func ParseHTML(t *testing.T, path string) string {
	body, err := readFromFile(path)
	assert.NoError(t, err)
	return body
}
