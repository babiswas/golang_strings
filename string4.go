package main

import (
	"flag"
	"fmt"
)

func find_binary_string(binary_string string) int {
	mp := make(map[int]int)
	sum := 0
	subarray_count := 0
	mp[sum] = 1
	for i := 0; i < len(binary_string); i++ {
		if binary_string[i] == '0' {
			sum += -1
			val, ok := mp[sum]
			if ok == true {
				subarray_count += val
				mp[sum] += 1
			} else {
				mp[sum] = 1
			}
		} else if binary_string[i] == '1' {
			sum += 1
			val, ok := mp[sum]
			if ok == true {
				subarray_count += val
				mp[sum] += 1
			} else {
				mp[sum] = 1
			}
		}
	}
	return subarray_count
}

func main() {
	pattern := flag.String("bin_string", "", "a binary string")
	flag.Parse()
	binary_string := *pattern
	count_sub_array := find_binary_string(binary_string)
	message := fmt.Sprintf("The number of binary string is %d", count_sub_array)
	fmt.Println(message)
}
