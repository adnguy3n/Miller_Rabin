package main

import "fmt"

/*
 * An implementation of the Miller Rabin algorithmn to check what numbers between
 * 500,000 and 1,000,000 are prime.
 */
func main() {
	//Checks from 500,001 and 999,999 as all Prime Numbers, save for 2, are odd.
	start := 500001
	end := 999999
	res := millerRabin(start, end, []int{})
	printResult(res)
}

/*
 * Performs the Miller Rabin algorithmn on integers between 500,000 and 1,000,000.
 * Appends the number to an array if the number is prime.
 */
func millerRabin(start, end int, res []int) []int {
	// Check if prime. Current only checks if the number is odd.
	if start%2 != 0 {
		res = append(res, start)
	}

	if start < end {
		res = millerRabin(start+1, end, res)
	}

	return res
}

/*
 * Prints the numbers in the result array.
 */
func printResult(res []int) {
	if len(res) == 0 {
		return
	}

	fmt.Println(res[0])
	printResult(res[1:])
}
