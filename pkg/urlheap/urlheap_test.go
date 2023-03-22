package urlheap

import (
	"container/heap"
	"testing"
)

func TestPop(t *testing.T) {
	items := map[string]int{
		"url1": 3, "url2": 2, "url3": 4,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			Url:      value,
			Priority: int64(priority),
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	want := &Item{
		Url:      "url0",
		Priority: 1,
	}

	heap.Push(&pq, want)
	got := heap.Pop(&pq).(*Item)

	if want.Priority != got.Priority {
		t.Errorf("heap.pop() = item.priority got %v, want %v", want.Priority, got.Priority)
	}

}
