package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/WizardMaggot/XE-Algorithm-Bruteforce/pirate"
	"github.com/WizardMaggot/XE-Algorithm-Bruteforce/primer"
)

//make sure everything is gone through
var lock sync.WaitGroup

//final "primed" string
var fin []int

//allows input to be used everywhere
var input string

//basic input and code exec
func main() {
	//input
	fmt.Printf("Enter filename> ")
	fmt.Scanln(&input)
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

		fmt.Println(fin)

		fmt.Printf("Key letters?\n> ")
		//changes to wanted letters afterwards
		fmt.Scanln(&input)

		run()
		defer lock.Wait()
	}
}

//bruteforces text and check() if string is inside
func run() {
	for i := pirate.Min(fin) - 255; i < pirate.Max(fin)+255; i++ {
		lock.Add(1)
		go check(i, input)
	}
}

//sees if string is in text
func check(rot int, inpc string) {
	var chk []string
	var see string
	for c := 0; c < len(fin); c++ {
		chk = append(chk, string(rune(fin[c]-rot)))
		see = strings.Join(chk, "")
	}
	if strings.Contains(see, inpc) {
		fmt.Println(see)
		os.Exit(0)
	} else {
		lock.Done()
	}
}
