package main

import (
	"flag"
	"fmt"
	"strings"
)

func sub_strings(s1 string) []string {
	temp_list := make([]string, 0)
	for i := 0; i < len(s1); i++ {
		temp_list = append(temp_list, s1[:i])
		temp_list = append(temp_list, string(s1[i]))
	}
	return temp_list
}

func count_all_substrings(s1 string, l []string) int {
	count := 0
	for _, val := range l {
		if strings.Contains(s1, val) {
			count += 1
		}
	}
	return count
}

func main() {
	str1 := flag.String("string1", "", "a string")
	str2 := flag.String("string2", "", "a string")
	flag.Parse()
	s1 := *str1
	s2 := *str2
	message := fmt.Sprintf("Create all the substring of %s", s1)
	fmt.Println(message)
	all_substrings := sub_strings(s1)
	fmt.Println("All substring of s1 is", " ", all_substrings)
	count := count_all_substrings(s2, all_substrings)
	message = fmt.Sprintf("The number of substrings is %d", count)
	fmt.Println(message)
}
