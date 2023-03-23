package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/palak92/clickhouse/pkg/urlheap"
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
	pq := make(urlheap.PriorityQueue, 0)
	heap.Init(&pq)
	scanner := bufio.NewScanner(file)
	// process line by line as file can be large
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		p, err := strconv.ParseInt(s[1], 10, 64)
		if err != nil {
			log.Fatalf("error while parsing int :%v", err)
			return
		}
		item := &urlheap.Item{
			Url:      s[0],
			Priority: p,
		}
		heap.Push(&pq, item)
		if pq.Len() > 10 {
			i := heap.Pop(&pq).(*urlheap.Item)
			log.Printf("popped the url %q  with value %d from min heap\n", i.Url, i.Priority)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Urls associated with 10 largest value is: ")
	for _, i := range pq {
		fmt.Printf("Url %q and long value %d \n", i.Url, i.Priority)
	}

}
