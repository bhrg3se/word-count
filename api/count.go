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

		//if word is not found insert
		if !ok {
			item = &Item{
				Value: word,
				Count: 1,
			}
			wordmap[word] = item
			pq.Push(item)

		} else { //if word is found update the "count"
			item.Count += 1
			wordmap[word] = item
			pq.update(item, word, item.Count)
		}
	}

	result := make([]Item, 0)

	// pop top 10 values into a new slice
	for pq.Len() > 0 && len(result) < 10 {
		item := heap.Pop(&pq).(*Item)
		result = append(result, *item)
	}

	return result
}
