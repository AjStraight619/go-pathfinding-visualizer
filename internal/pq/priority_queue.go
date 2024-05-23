package pq

import (
	"container/heap"

	"github.com/AjStraight619/pathfinding-go/internal/grid"
)

type PriorityQueue []*grid.Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].F < pq[j].F
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {

	*pq = append(*pq, x.(*grid.Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(node *grid.Node, f int) {
	for i, n := range *pq {
		if n == node {
			(*pq)[i].F = f
			heap.Fix(pq, i)
			break
		}
	}
}
