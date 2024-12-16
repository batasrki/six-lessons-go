package sorting

import (
	"reflect"
	"testing"
)

func TestNewSorter(t *testing.T) {
	type args struct {
		sortType string
		maxRange int
	}
	tests := []struct {
		name string
		args args
		want *Sorter
	}{
		{"bubble sort", args{"bubble", 0}, NewSorter("bubble")},
		{"quick sort", args{"quicksort", 0}, NewSorter("quicksort")},
		{"counting sort", args{"countingsort", 1}, NewSorter("countingsort", 1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSorter(tt.args.sortType, tt.args.maxRange); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSorter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSorter_Sort(t *testing.T) {
	type fields struct {
		SortType string
		MaxRange int
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
		{"bubble sort", fields{"bubble", 0}, args{[]int{1, 4, 3, 5, 10, 2}}, []int{1, 2, 3, 4, 5, 10}},
		{"quicksort", fields{"quicksort", 0}, args{[]int{1, 4, 3, 5, 10, 2}}, []int{1, 2, 3, 4, 5, 10}},
		{"unsuppported", fields{"unsupported", 0}, args{[]int{1, 4, 3, 5, 10, 2}}, []int{1, 4, 3, 5, 10, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sorter{
				SortType: tt.fields.SortType,
				MaxRange: tt.fields.MaxRange,
			}
			if got := s.Sort(tt.args.coll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sorter.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSorter_CountingSort(t *testing.T) {
	type fields struct {
		SortType string
		MaxRange int
	}
	type args struct {
		coll     []Customer
		maxRange int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Customer
	}{
		{"counting sort", fields{"countingsort", 10}, args{inputCustomers(), 10}, expectedCustomerOrder()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sorter{
				SortType: tt.fields.SortType,
				MaxRange: tt.fields.MaxRange,
			}
			if got := s.CountingSort(tt.args.coll, tt.args.maxRange); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sorter.CountingSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// {1, 4, 3, 9, 5, 9, 2, 7, 6, 8}
func inputCustomers() []Customer {
	return []Customer{
		{
			ID:           "C0",
			NumPurchases: 1,
		},
		{
			ID:           "C1",
			NumPurchases: 4,
		},
		{
			ID:           "C2",
			NumPurchases: 3,
		},
		{
			ID:           "C3",
			NumPurchases: 9,
		},
		{
			ID:           "C4",
			NumPurchases: 5,
		},
		{
			ID:           "C6",
			NumPurchases: 9,
		},
		{
			ID:           "C7",
			NumPurchases: 2,
		},
		{
			ID:           "C8",
			NumPurchases: 7,
		},
		{
			ID:           "C9",
			NumPurchases: 6,
		},
		{
			ID:           "C10",
			NumPurchases: 8,
		},
	}
}

// []int{0, 1, 2, 3, 4, 5, 6, 7, 9, 10, 10}
func expectedCustomerOrder() []Customer {
	return []Customer{
		{
			ID:           "C0",
			NumPurchases: 1,
		},
		{
			ID:           "C7",
			NumPurchases: 2,
		},
		{
			ID:           "C2",
			NumPurchases: 3,
		},
		{
			ID:           "C1",
			NumPurchases: 4,
		},
		{
			ID:           "C4",
			NumPurchases: 5,
		},
		{
			ID:           "C9",
			NumPurchases: 6,
		},
		{
			ID:           "C8",
			NumPurchases: 7,
		},
		{
			ID:           "C10",
			NumPurchases: 8,
		},
		{
			ID:           "C3",
			NumPurchases: 9,
		},
		{
			ID:           "C6",
			NumPurchases: 9,
		},
	}
}
