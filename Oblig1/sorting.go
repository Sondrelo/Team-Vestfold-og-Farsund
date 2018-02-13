package algorithms

// Les https://en.wikipedia.org/wiki/Bubble_sort
func Bubble_sort_modified(list []int) {
	n := len(list)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n-1; i++ {
			if list[i-1] > list[i] {
				// swap values using Go's tuple assignment
				list[i], list[i-1] = list[i-1], list[i]
				swapped = true
			}
		}
	}


}
