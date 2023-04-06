package main

import "fmt"


func main() {
	const count int = 1000
	const x, y = 3, 5

	var result int = Original(count, x, y)
	
	fmt.Printf("Sum of all numbers: %d\n", result)

}

func Original(count int, x int, y int)  int {
	var multiplier int

	// iterate
	
	for i := 1; i < count; i++ {
		if (i % x == 0) || (i % y == 0) {
			multiplier += i
		}
	}

	return multiplier
	
}

func Modeled(count int, x int, y int) int {
	// formula using arithmetic progression
	/*
		(n)th last term; l = a + (n -1)d = d+(int(x/d)-1)*d=d*int(x/d)
		simplified; S = d * int(x/d) * (1 + int(x/d))/2

		T = S3 + S5 + S15 ,(15) is divisible by both 3 and 5
	*/
	
	return 0
}