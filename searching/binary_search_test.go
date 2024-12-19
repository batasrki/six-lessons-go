package searching

import (
	"reflect"
	"testing"
)

func TestNewBinarySearch(t *testing.T) {
	tests := []struct {
		name string
		want *BinarySearch
	}{
		{"basic", &BinarySearch{ItemsSearched: 0, IndexOfFoundItem: -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinarySearch(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearch_SearchInts(t *testing.T) {
	type args struct {
		coll      []int
		searchFor int
	}
	tests := []struct {
		name string
		bs   *BinarySearch
		args args
		want *BinarySearch
	}{
		{"found", NewBinarySearch(), args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10}, &BinarySearch{IndexOfFoundItem: 9, ItemsSearched: 4, IndexChecked: 9}},
		{"not found", NewBinarySearch(), args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11}, &BinarySearch{IndexOfFoundItem: -1, ItemsSearched: 5, IndexChecked: 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bs.SearchInts(tt.args.coll, tt.args.searchFor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinarySearch.SearchInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
