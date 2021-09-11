package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/WizardMaggot/XE-Algorithm-Bruteforce/decode"
	"github.com/WizardMaggot/XE-Algorithm-Bruteforce/primer"
)

//make sure everything is gone through
var lock sync.WaitGroup

//final "primed" string
var fin []int

//struct input

//basic input and code exec
func main() {
	//var input string
	input := "ex.txt"
	//fmt.Printf("Enter filename> ")
	//fmt.Scanln(&input)
	if input == "" {
		fmt.Println("Enter filename\n> ")
		main()
	} else {
		f, err := os.ReadFile(input)
		if err != nil {
			fmt.Println("Failed to read file")
			main()
		}
		cf := string(f)
		fin = primer.Prime(cf)
		fmt.Printf("Key letters?\n> ")
		fmt.Scanln(&input)

		run := decode.Generate(fin)

		run(input)
	}
}
