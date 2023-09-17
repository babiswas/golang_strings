package main

import (
	"flag"
	"fmt"
)

func count_sub_string(str1 string) int {
	mp := make(map[int]int)
	sum := 0
	mp[sum] += 1
	subarray_count := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] == '0' {
			sum -= 1
			val, ok := mp[sum]
			if ok == true {
				subarray_count += val
				mp[sum] += 1
			} else {
				mp[sum] = 1
			}
		} else if str1[i] == '1' {
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

func count_0_1(str1 string) bool {
	count1 := 0
	count0 := 0
	for _, val := range str1 {
		if val == '0' {
			count0 += 1
		} else if val == '1' {
			count1 += 1
		}
	}
	return count1 == count0
}

func find_all_substrings(str1 string) {
	for i := 0; i < len(str1); i++ {
		for j := i; j < len(str1); j++ {
			temp := str1[i : j+1]
			if count_0_1(string(temp)) == true {
				fmt.Println(temp)
			}
		}
	}
}

func main() {
	binary_string := flag.String("b_string", "", "a string")
	flag.Parse()
	str1 := *binary_string
	fmt.Println("======================")
	find_all_substrings(str1)
	fmt.Println("=======================")
	fmt.Println(count_sub_string(str1))
	fmt.Println("=======================")
}
