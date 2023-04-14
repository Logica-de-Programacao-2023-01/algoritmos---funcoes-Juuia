package main

import "fmt"

func divide(a, b float64) (float64, error) {

	if b == 0 {

		return 0, fmt.Errorf("cannot divide by zero")

	}
	return a / b, nil
}

func main() {
	var a, b float64
	fmt.Println("Digite dois números: ")
	fmt.Scan(&a)
	fmt.Scan(&b)

	resultado, err := divide(a, b)
	if err != nil {
		fmt.Printf("Ocorreu um erro ao dividir a e b: %s\n", err)

		return
	}

	fmt.Println(resultado)
}
