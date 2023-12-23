/*soal nomer 1*/

package main

import (
	"fmt"
	"sort"
)

func getMaximumAmount(quantity []int, m int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(quantity)))

	var totalRevenue int
	for i := 0; i < m; i++ {
		if quantity[0] <= 0 {
			break
		}
		totalRevenue += quantity[0]
		quantity[0]--

		sort.Sort(sort.Reverse(sort.IntSlice(quantity)))
	}

	return totalRevenue
}

func main() {
	quantity1 := []int{10, 10, 8, 9, 1, 6}
	m1 := 6
	maxAmount1 := getMaximumAmount(quantity1, m1)
	fmt.Println(maxAmount1) // Output: 55

	quantity2 := []int{8, 8, 8, 8}
	m2 := 4
	maxAmount2 := getMaximumAmount(quantity2, m2)
	fmt.Println(maxAmount2) // Output: 32
}
