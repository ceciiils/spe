/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 07/11/23, 14:53
 * Copyright (c) 2023
 */

package main

import (
	"fmt"
	"math"
	"strconv"
)

func isNarcissisticNumber(num int) bool {
	numStr := strconv.Itoa(num)
	numDigits := len(numStr)

	sum := 0
	for _, digit := range numStr {
		digitValue, _ := strconv.Atoi(string(digit))
		sum += int(math.Pow(float64(digitValue), float64(numDigits)))
	}
	return sum == num
}

func main() {
	numbersToCheck := []int{111, 153}
	for _, num := range numbersToCheck {
		fmt.Println(isNarcissisticNumber(num))
	}
}
