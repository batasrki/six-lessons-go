package searching

type BinarySearch struct {
	ItemsSearched    int
	IndexOfFoundItem int
}

func NewBinarySearch() *BinarySearch {
	return &BinarySearch{
		ItemsSearched:    0,
		IndexOfFoundItem: -1,
	}
}

func (bs *BinarySearch) SearchInts(coll []int, searchFor int) *BinarySearch {
	lo := 0
	hi := len(coll) - 1

	for lo <= hi {
		bs.ItemsSearched++
		mid := (lo + hi) / 2

		if coll[mid] < searchFor {
			lo = mid + 1
		} else if coll[mid] > searchFor {
			hi = mid - 1
		} else {
			bs.IndexOfFoundItem = mid
			return bs
		}
	}

	return bs
}
