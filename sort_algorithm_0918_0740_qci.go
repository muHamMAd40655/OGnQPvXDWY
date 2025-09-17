// 代码生成时间: 2025-09-18 07:40:08
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Sorter is an interface that groups the Sort method.
type Sorter interface {
    Sort() []int
}

// IntSlice attaches the methods of Sorter to []int,
// creating a Sorter value.
type IntSlice []int

// Sort implements the Sorter interface for IntSlice.
func (p IntSlice) Sort() []int {
    // Randomized quicksort algorithm
    return p.quicksort(0, len(p)-1)
}

// quicksort is an implementation of the quicksort algorithm.
func (p IntSlice) quicksort(left, right int) []int {
    if left < right {
        // Partition the array and get the index of the pivot element
        p, pivotIndex := p.partition(left, right)
        
        // Recursively sort the elements on the left and right of the pivot
        p.quicksort(left, pivotIndex-1)
        p.quicksort(pivotIndex+1, right)
    }
    return p
}

// partition rearranges the elements in the range [left, right]
// so that all elements less than the pivot are on the left,
// all elements greater than the pivot are on the right,
// and the pivot element is in its final sorted position.
// It returns the new array and the index of the pivot element.
func (p IntSlice) partition(left, right int) (IntSlice, int) {
    pivotValue := p[right]
    var storeIndex int = left
    for i := left; i < right; i++ {
        // If the current value is less than the pivot, it swaps the values
        if p[i] < pivotValue {
            p[i], p[storeIndex] = p[storeIndex], p[i]
            storeIndex++
        }
    }
    // Swaps the pivot element with the element at the storeIndex
    p[right], p[storeIndex] = p[storeIndex], p[right]
    return p, storeIndex
}

func main() {
    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())
    
    // Create a slice with random numbers
    numbers := make(IntSlice, 10)
    for i := range numbers {
        numbers[i] = rand.Intn(100)
    }
    
    fmt.Println("Original slice: ", numbers)
    
    // Sort the slice
    sortedNumbers := numbers.Sort()
    
    fmt.Println("Sorted slice: ", sortedNumbers)
}
