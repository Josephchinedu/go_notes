// mini golang note taking script.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	var note string
	var notes []string
	var done bool
	var err error

	fmt.Println("Enter notes, or 'done' to exit:")

	for {
		fmt.Scanln(&input)
		if input == "done" {
			done = true
			break
		}
		note = strings.TrimSpace(input)
		notes = append(notes, note)
	}

	if done {
		fmt.Println("Notes:")
		fmt.Println(notes)
	} else {
		err = fmt.Errorf("Error: no notes entered")
		fmt.Println(err)
	}
}
