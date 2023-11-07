/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 07/11/23, 15:05
 * Copyright (c) 2023
 */

package main

import (
	"fmt"
)

func findNeedle(haystack []string, needle string) int {
	for i, str := range haystack {
		if str == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack := []string{"apple", "banana", "cherry", "date", "fig"}
	needle := "cherry"

	index := findNeedle(haystack, needle)
	fmt.Println(index)
}
