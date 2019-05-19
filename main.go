package main

import (
	"fmt"
	"log"

	"github.com/rciurlea/cn/mat"
)

func main() {
	A := mat.New(
		5, 5,
		11, 23, 31, 40, 51,
		62, 17, 18, 39, 10,
		21, 12, 33, 14, 45,
		26, 17, 48, 19, 30,
		23, 20, 23, 24, 15,
	)
	L, U, err := mat.LU(A)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("A =", A)
	fmt.Println("L =", L)
	fmt.Println("U =", U)
	fmt.Println("L * U =", L.Mul(U))
}
