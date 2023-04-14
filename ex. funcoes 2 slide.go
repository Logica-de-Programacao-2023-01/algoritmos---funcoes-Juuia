package main

import "fmt"

func calcularmedia(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("A lista está vazia")
	}

	soma := 0.
	for _, number := range numbers {
		soma += number
	}
	return soma / float64(len(numbers)), nil
}

func main() {
	var media float64

	media, err := calcularmedia([]float64{7, 8, 9, 9, 87, 7})

	if err != nil {
		fmt.Println("erro")
	} else {
		fmt.Printf("A média é %.2f", media)
	}

}
