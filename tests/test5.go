package main

import (
	"fmt"
)

type pirate struct {
	arr []string
}

func main() {
	c := pirate{[]string{"hello", "yee", "gobble"}}
	c.Pot()
}

func (d *pirate) Pot() {
	fmt.Println(d.arr)
}
