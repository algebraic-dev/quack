package main

import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
	parser "github.com/algebraic-sofia/quack/internal"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var events []parser.Event

	for scanner.Scan() {
		parsed, err := parser.Validate(scanner.Text())

		if err != nil {
			continue
		}

		result, err := parser.Parse(parsed)

		if err != nil {
			panic(err)
		}

		events = append(events, result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	analytics := parser.CollectData(events)

	bt, err := json.MarshalIndent(analytics, "", "   ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bt))
}
