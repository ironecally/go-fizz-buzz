package main

import "fmt"

var fizzWording, buzzWording string

func init() {
	fizzWording = "fizz"
	buzzWording = "buzz"
}

//SetWording: set current word for fizz and buzz condition
func SetWording(fizzStr, buzzStr string) {
	fizzWording = fizzStr
	buzzWording = buzzStr
}

//GetWording: Get current word for fizz and buzz condition
func GetWording() {
	fmt.Println("fizzWording :", fizzWording)
	fmt.Println("buzzWording :", buzzWording)
}

//IsFizz: will check if number is multiplication of 3
func IsFizz(number int) bool {
	if number%3 == 0 {
		return true
	}
	return false
}

//IsBuzz: will check if number is multiplication of 5
func IsBuzz(number int) bool {
	if number%5 == 0 {
		return true
	}
	return false
}
