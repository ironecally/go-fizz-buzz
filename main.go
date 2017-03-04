package main

import "fmt"

func main() {
	var number int
	fmt.Print("Please insert number :")
	fmt.Scanf("%d", &number)
	Generate(number).PrintWithCustomSeparator(",")
}
