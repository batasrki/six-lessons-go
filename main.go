package main

import (
	"fmt"
	"math/rand"
	"time"

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
	cs := s.NewSorter("countingsort", maxRange)
	customerSlice := makeCustomerSlice(maxRange)

	fmt.Printf("\nMade an array of %d items for counting sort\n", len(customerSlice))
	startTime = time.Now()
	customerSlice = cs.CountingSort(customerSlice, maxRange)

	fmt.Println("Time it took to sort the collection: ", time.Since(startTime))
	printCustomerSlice(customerSlice, 40)
}

func makeCustomerSlice(maxRange int) []s.Customer {
	coll := make([]s.Customer, maxRange)

	for i := range coll {
		coll[i] = s.Customer{
			ID:           fmt.Sprintf("C%d", i),
			NumPurchases: rand.Intn(maxRange),
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
