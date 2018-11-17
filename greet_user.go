package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name string
	if len(os.Args) > 1 {
		name = strings.Join(os.Args[1:], " ")
	} else {
		fmt.Print("What's your name?\n    ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		name = scanner.Text()
	}
	fmt.Printf("Hello %s!\n", name)
}
