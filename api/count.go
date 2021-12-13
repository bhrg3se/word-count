package api

import (
	"bufio"
	"container/heap"
	"io"
)

func countWords(reader io.Reader) []Item {

	wordmap := make(map[string]*Item)

	//initialize a priority queue
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	// loop through words and keep pushing them into priority queue
	for scanner.Scan() {
		word := scanner.Text()

		item, ok := wordmap[word]
		if !ok {
			item = &Item{
				Value: word,
				Count: 1,
			}
			wordmap[word] = item
			pq.Push(item)

		} else {
			item.Count += 1
			wordmap[word] = item
			pq.update(item, word, item.Count)
		}
	}

	result := make([]Item, 0)

	for pq.Len() > 0 && len(result) < 10 {
		item := heap.Pop(&pq).(*Item)
		result = append(result, *item)
	}

	return result
}
