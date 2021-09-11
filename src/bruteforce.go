package main

import (
	"fmt"
	"os"
	"strings"
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

		run, hide := decode.Generate(fin)

		fmt.Println(run, hide)
	}
}

//sees if string is in text
func check(u int, inpc string) {
	var chk []string
	var see string
	for c := 0; c < len(fin); c++ {
		chk = append(chk, string(rune(fin[u])))
		see = strings.Join(chk, "")
	}
	if strings.Contains(see, inpc) {
		println("Found!\n" + "l")
		os.Exit(0)
	} else {
		lock.Done()
	}
}
