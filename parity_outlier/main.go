/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 07/11/23, 14:57
 * Copyright (c) 2023
 */

package main

import "fmt"

func findOutlier(arr []int) int {
	evenCount, oddCount := 0, 0
	var even, odd int

	for _, num := range arr {
		if num%2 == 0 {
			evenCount++
			even = num
		} else {
			oddCount++
			odd = num
		}

		if evenCount > 0 && oddCount > 0 {
			break
		}
	}

	if evenCount == 1 {
		return even
	} else if oddCount == 1 {
		return odd
	}

	return 0
}

func main() {
	var data [][]int
	arr1 := []int{2, 4, 0, 100, 4, 11, 2602, 36}
	data = append(data, arr1)

	arr2 := []int{160, 3, 1719, 19, 11, 13, -21}
	data = append(data, arr2)

	arr3 := []int{11, 13, 15, 19, 9, 13, -21}
	data = append(data, arr3)

	for _, datum := range data {
		outlier := findOutlier(datum)
		if outlier == 0 {
			fmt.Println("false")
		} else {
			fmt.Printf("%d\n", outlier)
		}
	}
}
