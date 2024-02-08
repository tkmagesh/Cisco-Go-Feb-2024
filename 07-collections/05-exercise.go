/* Write a function "genPrimes" that generates the prime numbers between the given start & end

Use the "genPrimes" function in "main" function to generate the prime numbers between 2 and 100 and print them one after another in the main() function
*/

package main

import "fmt"

func main() {

	primeNos := genPrimes(2, 100)
	for _, no := range primeNos {
		fmt.Println(no)
	}
}

func genPrimes(start, end int) []int {
	result := make([]int, 0)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			result = append(result, no)
		}
	}
	return result
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
