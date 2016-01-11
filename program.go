package main

import (
	"fmt"
	"errors"
)

func main() {
	
	fmt.Println("Pruebas de Golang")
	
	result, err := multiply(2,3)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println("Multiplicación: ", result)
	
	createWebserver()
		
}

func multiply(a int, b int) (int, error) {
	
	if a>10 || b > 10 {
		return 0, errors.New("Numbers too big")
	}
	
	return a * b, nil
}

