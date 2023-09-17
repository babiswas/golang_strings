package main

import (
	"flag"
	"fmt"
	"strings"
)

func find_replace_str(s1 string, s2 string, s3 string) string {
	message := fmt.Sprintf("Before replacement the string is %s", s1)
	fmt.Println(message)
	index := strings.Index(s1, s2)
	for index != -1 {
		end_index := index + len(s2) - 1
		s1 = s1[:index] + s3 + s1[end_index+1:]
		index = strings.Index(s1, s2)
	}
	return s1
}

func find_replace(s1 string, s2 string, s3 string) string {
	message := fmt.Sprintf("Before replacement the string is %s", s1)
	result := s1
	fmt.Println(message)
	for strings.Contains(result, s2) {
		result = strings.Replace(result, s2, s3, 1)
	}
	return result
}

func main() {
	fmt.Println("===========================")
	fmt.Println("Find and replace strings:")
	string1 := flag.String("string1", "", "a string")
	string2 := flag.String("string2", "", "a string")
	string3 := flag.String("string3", "", "a string")
	flag.Parse()
	str1 := *string1
	str2 := *string2
	str3 := *string3
	fmt.Println("Finding and replacing strings:")
	fmt.Println(find_replace(str1, str2, str3))
	fmt.Println(find_replace_str(str1, str2, str3))
}
