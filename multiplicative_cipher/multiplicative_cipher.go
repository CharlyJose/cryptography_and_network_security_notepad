package main

import (
	"fmt"
	"os"
	"strconv"

	math "github.com/charlyjose/math"
)

func main() {
	A, _ := strconv.Atoi(os.Args[1])
	B, _ := strconv.Atoi(os.Args[2])
	a := math.Factors(A)
	b := math.Factors(B)
	fmt.Printf("\nFactors of A: %d", a)
	fmt.Printf("\nFactors of B: %d", b)
	d := math.GCD(A, B)
	fmt.Printf("\nGCD of A, B: %d", d)
}
