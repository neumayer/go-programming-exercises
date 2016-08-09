package arrays

import "github.com/neumayer/go-programming-exercises/utils"

type quicksort struct {
}

func quicklysort(array []int, lower int, upper int) {
	pivotvalue := array[lower+(upper-lower)/2]
	i := lower
	j := upper
	for i < j {
		for array[i] < pivotvalue {
			i++
		}
		for array[j] > pivotvalue {
			j--
		}
		if i <= j {
			arrayutils.Swap(array, i, j)
			i++
			j--
		}
	}
	if lower < j {
		quicklysort(array, lower, j)
	}
	if i < upper {
		quicklysort(array, i, upper)
	}
}

func (s quicksort) Sort(array []int) {
	l := len(array)
	quicklysort(array, 0, l-1)
}
