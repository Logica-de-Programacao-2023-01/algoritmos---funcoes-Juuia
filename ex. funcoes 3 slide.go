package main

import "fmt"

func main() {
	result, err := count("abc")
	if err != nil {
		fmt.Printf("Ocorreu um erro %s", err)
	} else {
		fmt.Println(result)
	}
}

func count(s string) (int, error) {
	chars := len(s)
	if chars == 0 {
		return 0, fmt.Errorf("A strimg está vazia")
	}
	return chars, nil
}
