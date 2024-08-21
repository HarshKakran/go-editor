package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	filePath := os.Args[1]

	clearScreen()
	fmt.Println("Press ctrl+c to exit or :wq to save and exit.")

	// read file
	textBuffer, err := readAndPrint(filePath)
	if err != nil {
		panic(err)
	}

	// start writing
	err = writeToFile(filePath, textBuffer)
	if err != nil {
		panic(err)
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func saveToFile(filePath string, textBuffer []byte) error {
	return os.WriteFile(filePath, textBuffer, 0666)
}

func writeToFile(filePath string, textBuffer []byte) error {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(">>")

		if scanner.Scan() {
			line := scanner.Text()

			if strings.ToLower(line) == ":wq" {
				break
			}

			lines = append(lines, line)
		}
	}

	textBuffer = append(textBuffer, []byte(strings.Join(lines, "\n"))...)

	return saveToFile(filePath, textBuffer)
}

func readAndPrint(filePath string) ([]byte, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return []byte{}, err
	}

	fmt.Print(string(file), "\n")

	return file, nil
}
