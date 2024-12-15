package main

import (
	"fmt"
	"math/rand"
	"time"

	s "github.com/batasrki/six-go-lessons/sorting"
)

func main() {
	bs := s.NewSorter("bubble")
	exampleSlice := makeRandomSlice(30000, 100000000)

	fmt.Printf("Made an array of %d items\n", len(exampleSlice))
	startTime := time.Now()
	bs.Sort(exampleSlice)

	fmt.Println("Time it took to sort the collection: ", time.Since(startTime))
	printSlice(exampleSlice, 40)
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
