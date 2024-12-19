package searching

import (
	"reflect"
	"testing"
)

func TestNewLinearSearch(t *testing.T) {
	tests := []struct {
		name string
		want *LinearSearch
	}{
		{"basic", &LinearSearch{ItemsSearched: 0, IndexOfFoundItem: -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinearSearch(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinearSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinearSearch_SearchInts(t *testing.T) {
	type args struct {
		coll      []int
		searchFor int
	}
	tests := []struct {
		name string
		ls   *LinearSearch
		args args
		want *LinearSearch
	}{
		{"found", NewLinearSearch(), args{[]int{1, 3, 4, 2, 6, 10, 2}, 10}, &LinearSearch{ItemsSearched: 6, IndexOfFoundItem: 5}},
		{"not found", NewLinearSearch(), args{[]int{1, 3, 4, 2, 6, 10, 2}, 11}, &LinearSearch{ItemsSearched: 7, IndexOfFoundItem: -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ls.SearchInts(tt.args.coll, tt.args.searchFor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinearSearch.SearchInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
