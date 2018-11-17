package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name string
	printHeader()
	if len(os.Args) > 1 {
		name = strings.Join(os.Args[1:], " ")
	} else {
		fmt.Print("What's your name?\n    ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		name = scanner.Text()
		fmt.Println()
	}
	fmt.Printf("Hello %s!\n", name)
}

func printHeader() {
	fmt.Println("\n/* greet_user.go")
	fmt.Println(" *")
	fmt.Print(" * Prints \"Hello NAME!\", where NAME is either provided at runtime")
	fmt.Println(" or interactively.")
	fmt.Println(" *")
	fmt.Println(" * @author  Glanderto Smegmoso")
	fmt.Println(" * @version 1.0, Nov 17 2018")
	fmt.Println(" */\n")
}
