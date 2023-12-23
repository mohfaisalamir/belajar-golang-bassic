/*soal nomer 2*/
package main

import (
	"fmt"
	"sort"
)

func getMaximumAmount(quantity []int, m int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(quantity)))

	var maximumAmount int
	for i := 0; i < m; i++ {
		if i < len(quantity) {
			maximumAmount += quantity[i]
			quantity[i] -= 1
		}
	}

	return maximumAmount
}

func main() {
	quantity := []int{-5, 4, -10, -1, -5, 8, -3}
	m := 3
	fmt.Println(getMaximumAmount(quantity, m)) // Output: 55
}
