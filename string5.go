package main

import (
	"flag"
	"fmt"
)

func find_all_ways(str1 string) int {
	count := 0
	result := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] == '0' {
			count += 1
		}
	}
	fmt.Println("The number of 0's in the string is ", count)
	K := count / 3
	mp := make(map[int]int)
	sum := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] == '0' {
			sum += 1
		}
		test, ok := mp[K]
		if (sum == 2*K) && ok && (i < len(str1)) {
			result += test
		}
		mp[sum] += 1
	}
	return result
}
func main() {
	bin_str := flag.String("bin_str", "", "a string")
	flag.Parse()
	temp_str := *bin_str
	count_ways := find_all_ways(temp_str)
	message := fmt.Sprintf("the number of possible ways %d", count_ways)
	fmt.Println(message)
}
