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

// func sorted(coll []int) bool {
// 	sorted := false
// 	for i := 1; i < len(coll); i++ {
// 		if coll[i-1] < coll[i] {
// 			sorted = true
// 		} else {
// 			sorted = false
// 			break
// 		}
// 	}
// 	return sorted
// }
