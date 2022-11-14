package leetcode

func kClosest(points [][]int, k int) [][]int {
	h := MinHeap{}
	for _, point := range points {
		distance := point[0]*point[0] + point[1]*point[1]
		h.Push(distance, point)
	}

	var output [][]int
	for i := 0; i < k; i++ {
		output = append(output, h.Pop().value)
	}

	return output
}

type Item struct {
	key   int
	value []int
}

type MinHeap struct {
	Data []Item
}

func (h *MinHeap) Push(key int, value []int) {
	// append to last
	h.Data = append(h.Data, Item{key, value})

	// compare with parents and swap
	idx := len(h.Data) - 1

	for idx > 0 {
		parentIdx := (idx+1)/2 - 1
		if h.Data[parentIdx].key > h.Data[idx].key {
			h.Data[parentIdx], h.Data[idx] = h.Data[idx], h.Data[parentIdx]
		}
		idx = (idx+1)/2 - 1
	}
}

func (h *MinHeap) Pop() Item {
	min := h.Data[0]

	h.Data[0] = h.Data[len(h.Data)-1]

	h.Data = h.Data[0 : len(h.Data)-1]

	idx := 0

	for !nodeIsLeaf(idx, h.Data) {
		nextIdx := h.findMinChild(idx)
		if nextIdx == -1 {
			return min
		}
		// do swap
		h.Data[idx], h.Data[nextIdx] = h.Data[nextIdx], h.Data[idx]
		idx = nextIdx
	}
	return min
}

// Determine current index reach the end of array
func nodeIsLeaf(idx int, data []Item) bool {
	leftIdx := idx*2 + 1
	rightIdx := idx*2 + 2

	if leftIdx > len(data) || rightIdx > len(data) {
		return true
	}

	return false
}

// Find min child of current node, return (true, pos of node) if min child existed
// if min child not existed, return false, -1
func (h *MinHeap) findMinChild(idx int) int {
	leftIdx := idx*2 + 1
	rightIdx := idx*2 + 2

	var nextIdx int
	if rightIdx < len(h.Data) && h.Data[rightIdx].key < h.Data[leftIdx].key {
		nextIdx = rightIdx
	} else {
		nextIdx = leftIdx
	}

	if h.Data[idx].key > h.Data[nextIdx].key {
		return nextIdx
	}

	return -1
}
