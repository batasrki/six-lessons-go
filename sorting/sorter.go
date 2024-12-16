package sorting

type Sorter struct {
	SortType string
	MaxRange int
}

type Customer struct {
	ID           string
	NumPurchases int
}

func NewSorter(sortType string, maxRange ...int) *Sorter {
	s := &Sorter{
		SortType: sortType,
	}

	for i := range maxRange {
		s.MaxRange = maxRange[i]
	}
	return s
}

func (s *Sorter) Sort(coll []int) []int {
	switch s.SortType {
	case "bubble":
		return bubbleSort(coll)
	case "quicksort":
		return quickSort(coll)
	default:
		return coll
	}
}

func bubbleSort(coll []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(coll); i++ {
			curr := coll[i]
			prev := coll[i-1]

			if prev > curr {
				coll[i-1] = curr
				coll[i] = prev
				swapped = true
			}
		}
	}
	return coll
}

func quickSort(coll []int) []int {
	if len(coll) < 2 {
		return coll
	}
	pivot := partition(coll)
	quickSort(coll[0:pivot])
	quickSort(coll[pivot+1:])
	return coll
}

func partition(coll []int) int {
	lo := 0
	hi := len(coll) - 1
	pivot := coll[hi]
	idx := lo - 1

	for j := lo; j < hi; j++ {
		if coll[j] <= pivot {
			idx += 1
			coll[idx], coll[j] = coll[j], coll[idx]
		}
	}
	idx += 1
	coll[idx], coll[hi] = coll[hi], coll[idx]
	return idx
}

func (s *Sorter) CountingSort(coll []Customer, maxRange int) []Customer {
	counts := make([]int, maxRange+1)
	output := make([]Customer, maxRange+1)

	// determine how many numbers repeat
	for i := range coll {
		counts[coll[i].NumPurchases] += 1
	}

	// ensure each slot has the number of elements less than or equal to its value
	for i := range counts {
		if i > 0 {
			counts[i] += counts[i-1]
		}
	}

	// add to the output array in sorted order
	for i := len(coll) - 1; i >= 0; i-- {
		output[counts[coll[i].NumPurchases]] = coll[i]
		counts[coll[i].NumPurchases] -= 1
	}

	return output[1:]
}
