package main

import "fmt"

func Fibo(a, b int) {
	fmt.Print(b, ", ")
	if b > 100 {
		fmt.Println()
		return
	}
	Fibo(b, a+b)
}

func NumeroMasGrande(nums ...int) int {
	bigger := 0

	for _, v := range nums {
		if v > bigger {
			bigger = v
		}
	}

	return bigger
}

func GeneradorImpares() func() uint {
	i := uint(1)
	return func() uint {
		var impar = i
		i += 2
		return impar
	}
}

func Intercambia(a *int, b *int) {
	valorB := *a
	valorA := *b

	*a = valorA
	*b = valorB
}

func main() {
	fmt.Println("Fibonacci (Recursividad)")
	Fibo(0, 1)

	fmt.Println("\nNúmero más grande (Variadic function): \n9, 3, 7, 5, 12, 15, 3, 4, 2, 9, 89, 23")
	fmt.Println(NumeroMasGrande(9, 3, 7, 5, 12, 15, 3, 4, 2, 9, 89, 23))

	fmt.Println("\nGenerador de impares (Cerraduras)")
	nextImpar := GeneradorImpares()
	fmt.Print(nextImpar(), ", ")
	fmt.Print(nextImpar(), ", ")
	fmt.Print(nextImpar(), ", ")
	fmt.Print(nextImpar(), ", ")
	fmt.Print(nextImpar(), ", ")
	fmt.Print(nextImpar(), ", ")
	fmt.Print(nextImpar(), ", ")
	fmt.Print(nextImpar(), ", ")
	fmt.Print(nextImpar(), ", ")
	fmt.Print(nextImpar())

	fmt.Println("\n\nIntercambio (Punteros)")
	a := 5
	b := 10
	Intercambia(&a, &b)
	fmt.Println("a:", a, "b:", b)
}
