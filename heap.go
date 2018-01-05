package heap

type Heap struct {
	h     []int
	count int
}

func NewMinHeap(cap int) Heap {
	heap := make([]int, cap)
	return Heap{heap, 0}
}

// Insert inserts item into the heap, returning true for success or false for failure
func (heap *Heap) Insert(item int) bool {
	// Check if overflowed length and resize
	if heap.count == cap(heap.h) {
		return false
	}

	// Place at the next open slot
	heap.h[heap.count] = item
	heap.count++

	// Percolate up
	heap.percolateUp()
	return true
}

func (heap *Heap) percolateUp() {
	if heap.count <= 0 {
		panic("percolateUp should never be called on empty heap")
	}

	elementPos := heap.count - 1
	parentPos := (elementPos - 1) / 2

	for elementPos != parentPos { // Reached top of heap
		if heap.h[elementPos] < heap.h[parentPos] {
			heap.h[parentPos], heap.h[elementPos] = heap.h[elementPos], heap.h[parentPos]
			elementPos = parentPos
			parentPos = (elementPos - 1) / 2
		} else {
			return // Element is as far as it should go up the heap
		}
	}
}

// RemoveMin removes and returns the min item from the heap.
// The bool return is true if RemoveMin succeeded, else false.
func (heap *Heap) RemoveMin() (int, bool) {
	// Check if we have elements to remove
	if heap.count <= 0 {
		return 0, false
	}

	// Remove element
	min := heap.h[0]
	heap.count--

	// Replace with last element
	heap.h[0] = heap.h[heap.count]

	// Percolate down
	heap.percolateDown()
	return min, true
}

func (heap *Heap) percolateDown() {
	pos := 0
	left, right := 1, 2
	swapPos := pos // Which child index to swap with

	for left < heap.count || right < heap.count { // At least one valid child
		if left < heap.count && heap.h[left] < heap.h[pos] { // Valid left swap
			swapPos = left
		}

		if right < heap.count && heap.h[right] < heap.h[pos] {
			if swapPos == pos { // Left side was not a good swap
				swapPos = right
			} else if heap.h[right] < heap.h[left] { // Both children good swaps, but right is better (swap with smaller child)
				swapPos = right
			}
		}

		if swapPos == pos { // No more percolating to do
			return
		}

		heap.h[pos], heap.h[swapPos] = heap.h[swapPos], heap.h[pos]
		pos = swapPos
		left, right = (pos*2)+1, (pos*2)+2
	}
}

// Sort uses the heap sort algorithm to return a sorted copy of orig
func Sort(orig []int) []int {
	// Insert all items into min heap
	heap := NewMinHeap(len(orig))
	for _, item := range orig {
		heap.Insert(item)
	}

	sorted := make([]int, len(orig))

	i := 0
	min, ok := heap.RemoveMin()
	for ok {
		sorted[i] = min
		i++
		min, ok = heap.RemoveMin()
	}

	return sorted
}
