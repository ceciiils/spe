/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 07/11/23, 15:09
 * Copyright (c) 2023
 */

package main

import (
	"fmt"
)

func blueOcean(list1, list2 []int) []int {
	var result []int

	list2Map := make(map[int]bool)
	for _, value := range list2 {
		list2Map[value] = true
	}

	for _, value := range list1 {
		if !list2Map[value] {
			result = append(result, value)
		}
	}

	return result
}

func main() {
	result := blueOcean([]int{1, 5, 5, 5, 3}, []int{5})
	result2 := blueOcean([]int{1, 2, 3, 4, 6, 10}, []int{1})

	fmt.Printf("%v\n", result)
	fmt.Printf("%v\n", result2) // Should print [1 2 4]
}
