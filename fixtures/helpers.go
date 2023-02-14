package fixtures

import (
	"bufio"
	"fmt"
	"os"
)

func ContainsText(text string, fixturePath string) bool {
	// Open the file
	file, err := os.Open(fixturePath)
	if err != nil {
		return false
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text()
		if line == text {
			fmt.Println("Match found:", line)
			return true
		}
	}

	return false
}
