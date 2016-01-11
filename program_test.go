package main

import (
	"testing"
	"fmt"
)

func Test_multiply_basic(t *testing.T) {

	t.Parallel()

	result, err := multiply(2,3)

	if err != nil {
		t.Fatal("We got some error",err)
	}

	if result != 6 {
		t.Fatal("2 x 3 Should be equals to 6")
	}
	
	fmt.Println("Test_multiply_basic OK")

}

func Test_multiply_big_number(t *testing.T) {
	
	_, err := multiply(11,3)
	
	if err==nil{
		t.Fatal("Invalid big number")
	}
	
	fmt.Println("Test_multiply_big_number OK")
}
