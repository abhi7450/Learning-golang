package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// alias
var pl = fmt.Println

func main() {
	// User Input
	pl("What is your name? ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello ", name)
	} else {
		log.Fatal(err)
	}

}
