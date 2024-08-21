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
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(string(file), "\n")

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

	// var file []byte

	file = append(file, []byte(strings.Join(lines, "\n"))...)

	err = saveToFile(filePath, file)
	if err != nil {
		fmt.Println(err)
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func saveToFile(fileName string, textBuffer []byte) error {
	return os.WriteFile(fileName, textBuffer, 0666)
}
