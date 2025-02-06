package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

/*
 * An implementation of the Miller Rabin algorithmn to check what numbers between
 * 500,000 and 1,000,000 are probably prime.
 */
func main() {
	// Checks from 500,001 and 999,999 as all Prime Numbers, save for 2, are odd.
	start := 500001
	end := 999999

	res := getPrimes(start, end, []int{})
	printResult(res)
	fmt.Println("/***********************************************************/")
	//fmt.Println(res)
	fmt.Println(len(res))

	// Runs the same calculations but uses Golang's implementation of Miller Rabin.
	res2 := getProbablePrimes(start, end, []int{})
	fmt.Println(len(res2))
}

/*
 * Performs the Miller Rabin algorithmn on integers in range (start, end).
 * Appends the number to an array if the number is probably prime.
 */
func getPrimes(start, end int, res []int) []int {
	if millerRabin(start) {
		res = append(res, start)
	}

	if start < end {
		res = getPrimes(start+1, end, res)
	}

	return res
}

/*
 * Implementation of the Miller Rabin Algorithm.
 */
func millerRabin(num int) bool {
	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	// Find q and k such that num - 1 = 2^k * q
	k, q := 0, num-1
	for q%2 == 0 {
		k++
		q /= 2
	}

	randomInt := rand.Intn(num-1) + 1
	randPowModNum := powMod(randomInt, q, num)

	// Check if randomInt^q mod n = 1 or
	// for j = 0 to k - 1 if randomInt^(2jq) = num - 1.
	if randPowModNum != 1 && randPowModNum != num-1 {
		// Calculate randomInt^(2jq) for j =0 to k - 1.
		j := 0
		for ; j < k; j++ {
			randPowModNum = powMod(randPowModNum, 2, num)
			if randPowModNum == num-1 {
				break
			}
		}

		if j == k {
			return false
		}
	}

	return true
}

/*
 * Modular Exponentation function.
 */
func powMod(base, exp, mod int) int {
	res := 1
	base = base % mod

	for exp > 0 {
		if exp%2 == 1 {
			res = (res * base) % mod
		}

		exp = exp / 2
		base = (base * base) % mod
	}

	return res
}

/*
 * Performs the Miller Rabin algorithmn on integers in range (start, end).
 * Appends the number to an array if the number is probably prime.
 * Uses the Golang's implementation of the Miller Rabin algorithm in math/big.
 * Used to compare with the above implementation of the Miller Rabin Algorithm.
 */
func getProbablePrimes(start, end int, res []int) []int {
	if big.NewInt(int64(start)).ProbablyPrime(40) {
		res = append(res, start)
	}

	if start < end {
		res = getPrimes(start+1, end, res)
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
