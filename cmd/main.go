package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/palak92/clickhouse/pkg/processor"
)

func main() {

	var filePath string

	fmt.Print("Enter the file path: ")
	fmt.Scanf("%s", &filePath)
	fmt.Printf("The file path is %q\n", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("error while opening file:%v", err)
		return
	}
	defer file.Close()
	// create min heap.
	processor := processor.New()
	scanner := bufio.NewScanner(file)
	// process line by line as file can be large
	for scanner.Scan() {
		processor.Add(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Urls associated with 10 largest value are: %q", *processor.TopUrls())

}
