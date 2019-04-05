package GoLangExamples

import (
	"fmt"

	"github.com/ehteshamz/greetings"
)

//ConcatSlice return contents of slice concatenated together
func ConcatSlice(sliceToConcat []byte) string {
	concatenatedString := ""
	for i := 0; i < len(sliceToConcat); i++ {
		if i != len(sliceToConcat)-1 {
			concatenatedString = concatenatedString + string(sliceToConcat[i]) + "-"
		} else {
			concatenatedString = concatenatedString + string(sliceToConcat[i])
		}
	}
	return concatenatedString
}

//Encrypt encrypts the provided slice
func Encrypt(sliceToEncrypt []byte, ceaserCount int) {

	for i, v := range sliceToEncrypt {
		sliceToEncrypt[i] = byte(int(v) + ceaserCount)
		fmt.Printf(string(sliceToEncrypt[i]))
	}
}

//EZGreetings greets to the name passed
func EZGreetings(name string) {
	fmt.Printf(greetings.PrintGreetings(name))
}
