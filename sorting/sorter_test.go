package sorting

import (
	"reflect"
	"testing"
)

func TestNewSorter(t *testing.T) {
	type args struct {
		sortType string
	}
	tests := []struct {
		name string
		args args
		want *Sorter
	}{
		{"bubble sort", args{"bubble"}, NewSorter("bubble")},
		{"quick sort", args{"quicksort"}, NewSorter("quicksort")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSorter(tt.args.sortType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSorter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSorter_Sort(t *testing.T) {
	type fields struct {
		SortType string
	}
	type args struct {
		coll []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{"bubble sort", fields{"bubble"}, args{[]int{1, 4, 3, 5, 10, 2}}, []int{1, 2, 3, 4, 5, 10}},
		{"quicksort", fields{"quicksort"}, args{[]int{1, 4, 3, 5, 10, 2}}, []int{1, 2, 3, 4, 5, 10}},
		{"unsuppported", fields{"unsupported"}, args{[]int{1, 4, 3, 5, 10, 2}}, []int{1, 4, 3, 5, 10, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sorter{
				SortType: tt.fields.SortType,
			}
			if got := s.Sort(tt.args.coll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sorter.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
