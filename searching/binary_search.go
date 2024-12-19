package searching

type BinarySearch struct {
	ItemsSearched    int
	IndexOfFoundItem int
	IndexChecked     int
}

func NewBinarySearch() *BinarySearch {
	return &BinarySearch{
		ItemsSearched:    0,
		IndexOfFoundItem: -1,
	}
}

func (bs *BinarySearch) SearchInts(coll []int, searchFor int) *BinarySearch {
	bs.ItemsSearched++
	currIdx := len(coll) / 2
	bs.IndexChecked += currIdx

	if len(coll) < 2 {
		return bs
	}

	if searchFor == coll[currIdx] {
		bs.IndexOfFoundItem = bs.IndexChecked
	} else if searchFor > coll[currIdx] {
		bs.SearchInts(coll[currIdx:], searchFor)
	} else {
		bs.SearchInts(coll[0:currIdx], searchFor)
	}

	return bs
}
