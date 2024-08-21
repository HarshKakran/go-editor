package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	clearScreen()
	fmt.Println("Press ctrl+c to exit or :wq to save and exit.")

	// write
	{
		var lines []string
		scanner := bufio.NewScanner(os.Stdin)

		for {
			fmt.Printf(">>")

			if scanner.Scan() {
				line := scanner.Text()

				if line == ":wq" {
					break
				}

				lines = append(lines, line)
			}
		}

		var textBuffer []byte

		textBuffer = append(textBuffer, []byte(strings.Join(lines, "\n"))...)

		err := saveToFile("output.txt", textBuffer)
		if err != nil {
			fmt.Println(err)
		}
	}

	// read
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func saveToFile(fileName string, textBuffer []byte) error {
	return os.WriteFile(fileName, textBuffer, 0666)
}
