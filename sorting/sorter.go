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
	numOfItems := len(coll)

	counts := make([]int, maxRange)
	output := make([]Customer, numOfItems)

	// determine how many numbers repeat
	for i := 0; i < numOfItems; i++ {
		counts[coll[i].NumPurchases]++
	}

	// ensure each slot has the number of elements less than or equal to its value
	for i := 1; i < maxRange; i++ {
		counts[i] += counts[i-1]
	}

	// add to the output array in sorted order
	for i := numOfItems - 1; i >= 0; i-- {
		output[counts[coll[i].NumPurchases]-1] = coll[i]
		counts[coll[i].NumPurchases] -= 1
	}

	return output
}
