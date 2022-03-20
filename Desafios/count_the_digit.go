package desafios

import (
	"fmt"
	"strings"
)

/*
Take an integer n (n >= 0) and a digit d (0 <= d <= 9) as an integer.
Square all numbers k (0 <= k <= n) between 0 and n.
Count the numbers of digits d used in the writing of all the k**2.
Call nb_dig (or nbDig or ...) the function taking n and d as parameters and returning this count.
*/
func NbDig(n int, d int) int {
	numbers := make([]int, 0)
	result := 0

	for i := 0; i <= n; i++ {
		numbers = append(numbers, i*i)
	}

	for _, number := range numbers {
		result += strings.Count(fmt.Sprint(number), fmt.Sprint(d))
	}

	return result
}
