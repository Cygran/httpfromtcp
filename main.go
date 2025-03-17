package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	messages, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer messages.Close()
	for {
		b := make([]byte, 8)
		n, err := messages.Read(b)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("Error reading file: ", err)
			break
		}
		fmt.Printf("read: %s\n", b[:n])
	}
}
