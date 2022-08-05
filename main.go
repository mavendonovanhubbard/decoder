package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"encoding/base64"
)

func main() {
	fmt.Println("Enter base64 encoded text:")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	
	rawDecodedText, err := base64.StdEncoding.DecodeString(input)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Decoded text: %s\n", rawDecodedText)
}
