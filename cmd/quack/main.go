package main

import (
	"fmt"
	parser "github.com/algebraic-sofia/quack/internal"
)

func main() {
	event, err := parser.Validate("  21:42 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT")

	if err != nil {
		panic(err)
	}

	fmt.Println(event)
}
