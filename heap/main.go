package main

type Heap struct {
	array []int
	size  int
}

func (h *Heap) Heapify(arr []int, i int) {
}

func (h *Heap) Poll() int {
	elem := h.array[0]
	// bubble
	return elem
}

func (h *Heap) Push(val int) {
	h.array = append(h.array, val)
	// bubble
}

func (h *Heap) percolateUp(index int) {
	parentIndex := (index - 1) / 2

	for index > 0 && h.array[index] < h.array[parentIndex] {
		h.array[index], h.array[parentIndex] = h.array[parentIndex], h.array[index]
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
}

func (h *Heap) percolateDown(index int) {
	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		smallestIndex := index

		if leftChildIndex < h.size && h.array[leftChildIndex] < h.array[smallestIndex] {
			smallestIndex = leftChildIndex
		}

		if rightChildIndex < h.size && h.array[rightChildIndex] < h.array[smallestIndex] {
			smallestIndex = rightChildIndex
		}

		if smallestIndex == index {
			break
		}

		h.array[index], h.array[smallestIndex] = h.array[smallestIndex], h.array[index]
		index = smallestIndex
	}
}

func (h *Heap) Peek() int {
	if len(h.array) == 0 {
		return 0
	}
	return h.array[0]
}

func problem() {}

func main() {
}
