package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputFilePath = "messages.txt"

func main() {
	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("could not open %s: %s\n", inputFilePath, err)
	}

	fmt.Printf("Reading data from %s\n", inputFilePath)
	fmt.Println("=====================================")

	currentLine := ""
	for {
		b := make([]byte, 8)
		n, err := f.Read(b)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s\n", err.Error())
			break
		}
		str := string(b[:n])
		parts := strings.Split(str, "\n")
		for i, part := range parts {
			if i < len(parts)-1 {
				currentLine += part
				fmt.Printf("read: %s\n", currentLine)
				currentLine = ""
			} else {
				currentLine += part
			}
		}
	}
	if currentLine != "" {
		fmt.Printf("read: %s\n", currentLine)
	}
}
