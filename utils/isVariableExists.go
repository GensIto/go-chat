package utils

import (
	"bufio"
	"os"
	"strings"
)

func IsVariableExists(filename string, varName string) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "var "+varName+" ") {
			return true, nil
		}
	}
	return false, scanner.Err()
}
