package sorting

type Sorter struct {
	SortType string
}

func NewSorter(sortType string) *Sorter {
	return &Sorter{
		SortType: sortType,
	}
}

func (s *Sorter) Sort(coll []int) []int {
	switch s.SortType {
	case "bubble":
		return bubbleSort(coll)
	case "quicksort":
		return quickSort(coll, 0, len(coll)-1)
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

func quickSort(coll []int, lo int, hi int) []int {
	if lo >= hi || lo < 0 {
		return coll
	}
	pivot := partition(coll, lo, hi)
	quickSort(coll, lo, pivot-1)
	quickSort(coll, pivot+1, hi)
	return coll
}

func partition(coll []int, lo int, hi int) int {
	pivot := coll[hi]
	idx := lo

	for j := lo; j < hi-1; j++ {
		if coll[j] <= pivot {
			coll[idx], coll[j] = coll[j], coll[idx]
			idx += 1
		}
	}
	coll[idx], coll[hi] = coll[hi], coll[idx]
	return idx
}
