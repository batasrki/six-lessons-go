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
