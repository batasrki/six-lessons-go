package main

import (
	"fmt"
	"math/rand"
	"time"

	ss "github.com/batasrki/six-go-lessons/searching"
	s "github.com/batasrki/six-go-lessons/sorting"
)

func main() {
	// bubble sort
	bs := s.NewSorter("bubble")
	exampleSlice := makeRandomSlice(30000, 100000000)

	fmt.Printf("Made an array of %d items for bubble sort\n", len(exampleSlice))
	startTime := time.Now()
	bs.Sort(exampleSlice)

	fmt.Println("Time it took to sort the collection: ", time.Since(startTime))
	printSlice(exampleSlice, 40)

	// quick sort
	qs := s.NewSorter("quicksort")
	exampleSlice = makeRandomSlice(30000, 100000000)

	fmt.Printf("\nMade an array of %d items for quicksort\n", len(exampleSlice))
	startTime = time.Now()
	qs.Sort(exampleSlice)

	fmt.Println("Time it took to sort the collection: ", time.Since(startTime))
	printSlice(exampleSlice, 40)

	// counting sort
	maxRange := 1000000
	numOfCustomers := 1000000
	cs := s.NewSorter("countingsort", maxRange)
	customerSlice := makeCustomerSlice(numOfCustomers, maxRange)

	fmt.Printf("\nMade an array of %d items for counting sort\n", len(customerSlice))
	startTime = time.Now()
	customerSlice = cs.CountingSort(customerSlice, maxRange)

	fmt.Println("Time it took to sort the collection: ", time.Since(startTime))
	printCustomerSlice(customerSlice, 40)

	// linear search
	ls := ss.NewLinearSearch()
	exampleSlice = makeRandomSlice(10000000, 1000000)

	startTime = time.Now()
	lst := ls.SearchInts(exampleSlice, 1001)

	fmt.Println("\nTime it took to search the collection: ", time.Since(startTime))
	fmt.Printf("Found item %d at index %d after %d lookups\n\n", 1001, lst.IndexOfFoundItem, lst.ItemsSearched)

	// binary search
	bins := ss.NewBinarySearch()
	exampleSlice = qs.Sort(exampleSlice)
	// exampleSlice = makeRandomSlice(300000, 1000000)

	startTime = time.Now()
	bins.SearchInts(exampleSlice, 1001)

	fmt.Println("Time it took to search the collection: ", time.Since(startTime))
	fmt.Printf("Found item %d at index %d after %d lookups", 1001, bins.IndexOfFoundItem, bins.ItemsSearched)

}

func makeCustomerSlice(numOfCustomers int, maxNumOfPurchases int) []s.Customer {
	coll := make([]s.Customer, numOfCustomers)

	for i := range coll {
		coll[i] = s.Customer{
			ID:           fmt.Sprintf("C%d", i),
			NumPurchases: rand.Intn(maxNumOfPurchases),
		}
	}
	return coll
}

func makeRandomSlice(numItems, max int) []int {
	coll := make([]int, numItems)

	for i := 0; i < len(coll); i++ {
		coll[i] = rand.Intn(max)
	}

	return coll
}

func printSlice(slice []int, maxNumItems int) {
	n := 0

	if len(slice) > maxNumItems {
		n = maxNumItems
	} else {
		n = len(slice)
	}

	fmt.Println(slice[0:n])
}

func printCustomerSlice(slice []s.Customer, maxNumItems int) {
	n := 0

	if len(slice) > maxNumItems {
		n = maxNumItems
	} else {
		n = len(slice)
	}

	fmt.Println(slice[0:n])
}
