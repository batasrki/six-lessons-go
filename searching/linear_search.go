package searching

type LinearSearch struct {
	ItemsSearched    int
	IndexOfFoundItem int
}

func NewLinearSearch() *LinearSearch {
	return &LinearSearch{
		ItemsSearched:    0,
		IndexOfFoundItem: -1,
	}
}

func (ls *LinearSearch) SearchInts(coll []int, searchFor int) *LinearSearch {
	for i := range coll {
		ls.ItemsSearched = i + 1
		if coll[i] == searchFor {
			ls.IndexOfFoundItem = i
			break
		}
	}
	return ls
}
