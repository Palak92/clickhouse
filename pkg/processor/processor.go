package processor

import (
	"container/heap"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/palak92/clickhouse/pkg/urlheap"
)

type Processor struct {
	pq   *urlheap.PriorityQueue
	urls map[int64][]string
}

func New() *Processor {
	pq := make(urlheap.PriorityQueue, 0)
	heap.Init(&pq)

	m := make(map[int64][]string)
	return &Processor{
		pq:   &pq,
		urls: m,
	}
}

func (p Processor) Add(text string) error {
	s := strings.Fields(text)
	val, err := strconv.ParseInt(s[1], 10, 64)
	if err != nil {
		return fmt.Errorf("error while parsing int :%w", err)
	}
	url := s[0]
	urls, ok := p.urls[val]
	if ok {
		urls = append(urls, url)
	}
	item := &urlheap.Item{
		Priority: val,
	}
	heap.Push(p.pq, item)
	p.urls[val] = []string{url}
	if p.pq.Len() > 10 {
		i := heap.Pop(p.pq).(*urlheap.Item)
		delete(p.urls, i.Priority)
		log.Printf("popped the value %d from min heap\n", i.Priority)
	}
	return nil
}

func (p Processor) TopUrls() *[]string {
	var urls []string
	for _, v := range p.urls {
		urls = append(urls, v...)
	}
	return &urls
}
