package main

import (
	"fmt"
	"strings"
)

// taken from https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/#cheat-sheet

func main() {

	fmt.Println(strings.Join([]string{"Hallo", "Du", "Da"}, " "))

	//In this example, fmt.Printf formats and writes to standard output:
	fmt.Printf("Binary: %b\\%b", 4, 5) // Prints `Binary: 100\101`

	//Newline
	fmt.Println()

	//As a special case, the verb %%, which consumes no argument, produces a percent sign:
	fmt.Printf("%d %%", 50) // Prints `50 %`

	fmt.Println()

	//Use fmt.Sprintf to format a string without printing it:
	s := fmt.Sprintf("Binary: %b\\%b", 4, 5) // s == `Binary: 100\101`
	Var_dump(s)

	//If you try to compile and run this incorrect line of code
	fmt.Printf("Binary: %b\\%b", 4) // An argument to Printf is missing.
	fmt.Println("")

	s = fmt.Sprintf("Das ist mein %v Test mit dem hier %v und das da %v", 12, "2", 3.7896)
	fmt.Println(s)

	//To catch this type of errors early, you can use the vet command – it can find calls whose arguments do not align with the format string.
	/**
	C:\Users\sam\go\src\GOLearn\fmt.sPrintF> go vet SPrintF.go
	# command-line-arguments
	.\SPrintF.go:23:2: Printf format %b reads arg #2, but call has 1 arg
	*/
}

func Var_dump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}
