package main

import (
	"fmt"
	"strconv"
	"strings"
)

type fizzBuzzer struct {
	Words []string
}

func Generate(number int) *fizzBuzzer {
	result := make([]string, number)
	for i := 1; i <= number; i++ {
		if IsFizz(i) || IsBuzz(i) {
			if IsFizz(i) {
				result[i-1] += fizzWording
			}
			if IsBuzz(i) {
				result[i-1] += buzzWording
			}
			continue
		}
		result[i-1] += strconv.Itoa(i)
	}
	return &fizzBuzzer{
		Words: result,
	}
}

func (fb *fizzBuzzer) Print() {
	fmt.Println(strings.Join(fb.Words, " "))
}
func (fb *fizzBuzzer) Println() {
	fmt.Println(strings.Join(fb.Words, "\n"))
}
func (fb *fizzBuzzer) PrintWithCustomSeparator(sep string) {
	fmt.Println(strings.Join(fb.Words, sep))
}
