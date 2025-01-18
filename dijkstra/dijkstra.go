package dijkstra

import (
	"container/heap"
	"math"
)

const infiniteDistance = math.MaxInt

type Point struct {
	X, Y int
}

type item struct {
	Point
	idx      int
	priority int
}

type dHeap []*item

func (h dHeap) Len() int { return len(h) }

func (h dHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }

func (h dHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *dHeap) Push(x any) {
	item := x.(*item)
	item.idx = len(*h)

	*h = append(*h, item)
}

func (h *dHeap) Pop() any {
	old := *h
	n := len(old)

	item := old[n-1]
	*h = old[0 : n-1]

	return item
}

func Run(input [][]int, st, end Point) (int, []Point) {
	n, m := len(input), len(input[0])
	distances := make(map[Point]int)
	prev := make(map[Point]*Point)

	for i := range n {
		for j := range m {
			distances[Point{X: i, Y: j}] = infiniteDistance
		}
	}

	distances[st] = input[st.X][st.Y]

	pq := &dHeap{}
	heap.Init(pq)
	heap.Push(pq, &item{Point: st, priority: input[st.X][st.Y]})

	directions := []Point{
		{-1, 0}, // влево
		{1, 0},  // вправо
		{0, -1}, // вверх
		{0, 1},  // вниз
	}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*item).Point
		if current == end {
			break
		}

		for _, dir := range directions {
			neighbor := Point{X: current.X + dir.X, Y: current.Y + dir.Y}

			if neighbor.X < 0 || neighbor.X >= n || neighbor.Y < 0 || neighbor.Y >= m {
				continue
			}

			neighborCost := input[neighbor.X][neighbor.Y]
			if neighborCost == 0 { // стена
				continue
			}

			newDist := distances[current] + neighborCost
			if newDist < distances[neighbor] {
				distances[neighbor] = newDist
				prev[neighbor] = &current

				heap.Push(pq, &item{Point: neighbor, priority: newDist})
			}
		}
	}

	shortestDistance := distances[end]

	if shortestDistance == infiniteDistance {
		return -1, nil
	}

	path := make([]Point, shortestDistance)
	for i, cur := shortestDistance-1, &end; cur != nil; i, cur = i-1, prev[*cur] {
		path[i] = *cur
	}

	return shortestDistance, path
}
