package main

import (
	"fmt"
	"os"
)

type Item struct {
	value int64
}

type Heap []Item

func (pq Heap) Len() int { return len(pq) }

func (pq Heap) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq Heap) GreaterThan(i, j int) bool { return pq[i].value > pq[j].value }

func (pq *Heap) Peek() Item {
	if pq.Len() == 0 {
		os.Exit(1)
	}
	return (*pq)[0]
}

func (pq *Heap) Push(item Item) {
	*pq = append(*pq, item)
	pq.siftUp(pq.Len() - 1)
}

func (pq *Heap) Pop() Item {
	pq.Swap(0, pq.Len()-1)

	// Pop last item.
	old := *pq
	n := pq.Len()
	item := old[n-1]
	*pq = old[:n-1]

	pq.siftDown(0)
	return item
}

func (pq *Heap) siftUp(childIdx int) {
	// While the process has not reached top.
	for childIdx > 0 {
		parentIdx := (childIdx - 1) / 2
		if pq.GreaterThan(parentIdx, childIdx) {
			pq.Swap(parentIdx, childIdx)
		}
		childIdx = parentIdx
	}
}

func (pq *Heap) siftDown(parentIdx int) {
	// While parent has at least one child.
	for parentIdx*2+1 < pq.Len() {
		childIdx := pq.smallerChildIndex(parentIdx)
		if pq.GreaterThan(parentIdx, childIdx) {
			pq.Swap(parentIdx, childIdx)
		}
		parentIdx = childIdx
	}
}

func (pq *Heap) smallerChildIndex(i int) int {
	left, right := i*2+1, i*2+2

	if right >= pq.Len() || !pq.GreaterThan(left, right) {
		return left
	} else {
		return right
	}
}

func main() {
	pq := &Heap{}

	items := []Item{{3}, {2}, {2}, {1}, {4}, {5}, {2}}
	for _, item := range items {
		pq.Push(item)
	}
	fmt.Printf("Peek: %d\n", pq.Peek())

	for pq.Len() > 0 {
		fmt.Printf("Pop: %d\n", pq.Pop())
	}
}
