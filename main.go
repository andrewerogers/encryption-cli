package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your password: ")
		in, _ := reader.ReadString('\n')
		out, err := encrypt(in)

		if err != nil {
			fmt.Printf("> An error occurred")
			break
		}

		fmt.Printf("> Your encrypted pasword is %s\n", out)
	}
}

func encrypt(s string) (string, error) {
	// todo: research open source libraries to use for this. 
	return "", nil
}
