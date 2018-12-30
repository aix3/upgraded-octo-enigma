package quicksort

// QuickSort is a quick sort impl
func QuickSort(array []int, start, end int) {
	if len(array) <= 1 {
		return
	}
	first := array[start]

	var i = end
	var j = start
	var c int

out:
	for ; i >= j; i-- {
		if i == j {
			swap(array, start, i)
			c = i
			break
		}
		if array[i] < first {
			for ; j <= i; j++ {
				if i == j {
					swap(array, start, i)
					c = i
					break out
				}
				if array[j] > first {
					swap(array, j, i)
					break
				}
			}
		}
	}

	if c > start {
		QuickSort(array, start, c-1)
	}

	if c < end {
		QuickSort(array, c+1, end)
	}
}

func swap(array []int, i, j int) {
	if i == j {
		return
	}
	// fmt.Println("Before swap: ", array[i], array[j])
	array[i] ^= array[j]
	array[j] ^= array[i]
	array[i] ^= array[j]
	// fmt.Println("After swap: ", array[i], array[j])
}
