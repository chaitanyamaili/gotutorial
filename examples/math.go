package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	i1, i2, i3, i4 := 12, 45, 68, 11
	intSum := i1 + i2 + i3 + i4
	fmt.Println("Integer Sum:", intSum)

	f1, f2, f3, f4 := 12.123456, 45.5, 68.7, 11.0
	floatSum := f1 + f2 + f3 + f4
	fmt.Println("Float Sum:", floatSum)

	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(23.5)
	b2.SetFloat64(65.1)
	b3.SetFloat64(76.3)
	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("BigSum = %.10g\n", &bigSum)

	circleRadius := 15.49
	circumference := circleRadius * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference)
}
