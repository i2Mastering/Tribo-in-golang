// Ike Iloegbu Project 3 CSC 6303
// Visual Studio Code

package main

import (
	"fmt"
	"math/big"
)

/**
 * tribonacci generates the first n numbers in the Tribonacci sequence.
 * The sequence starts with three 1's and each subsequent number is the sum
 * of the previous three numbers.
 *
 * This implementation uses the math/big package to handle very large numbers,
 * avoiding integer overflow.
 *
 * @param n the number of elements in the Tribonacci sequence to generate
 * @return a slice of *big.Int containing the Tribonacci sequence up to n elements
 */
func tribonacci(n int) []*big.Int {
	if n <= 0 {
		return []*big.Int{}
	} else if n == 1 {
		return []*big.Int{big.NewInt(0)}
	} else if n == 2 {
		return []*big.Int{big.NewInt(0), big.NewInt(1)}
	} else {
		fibSeq := []*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(1)}
		for i := 3; i < n; i++ {
			nextNum := new(big.Int).Add(fibSeq[i-1], fibSeq[i-2])
			nextNum.Add(nextNum, fibSeq[i-3])
			fibSeq = append(fibSeq, nextNum)
		}
		return fibSeq
	}
}

/**
 * The main function simulates different types of input and demonstrates
 * type handling similar to Python's error messaging.
 *
 * If the input is a string, it outputs a TypeError message related to string comparison.
 * If the input is an integer, it calculates and prints the Tribonacci sequence.
 * If the input is a float or any other type, it outputs a TypeError message
 * indicating a float cannot be interpreted as an integer.
 */
func main() {
	var input interface{} = 5

	switch val := input.(type) {
	case string:
		fmt.Print("TypeError: '<=' not supported between instances of 'str' and 'int'")
	case int:
		fmt.Print(tribonacci(val))
	default:
		fmt.Print("TypeError: 'float' object cannot be interpreted as an integer")
	}
}
