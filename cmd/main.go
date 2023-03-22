package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the file path: ")

	filePath, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("error while reading filepath: %w", err)
		return
	}
	fmt.Printf("The file path is %s\n", filePath)
}
