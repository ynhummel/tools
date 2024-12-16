package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines := flag.Bool("l", false, "print the newline counts")
	bytes := flag.Bool("c", false, "print the byte counts")
	multibytes := flag.Bool("m", false, "print the character counts")
	words := flag.Bool("w", false, "print the word counts")

	flag.Parse()

	content := make([]byte, 0)
	fileName := ""
	var err error

	if len(flag.Args()) == 0 {
		content, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fileName := flag.Args()[0]
		content, err = os.ReadFile(fmt.Sprintf("./%s", fileName))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

	if flag.NFlag() == 0 {
		fmt.Printf("%6d %6d %6d ", lineCount(content), wordCount(content), byteCount(content))
	}

	if *lines {
		fmt.Printf("%d ", lineCount(content))
	}

	if *words {
		fmt.Printf("%d ", wordCount(content))
	}

	if *multibytes {
		fmt.Printf("%d ", characterCount(content))
	}

	if *bytes {
		fmt.Printf("%d ", byteCount(content))
	}

	fmt.Println(fileName)
	os.Exit(0)
}

func lineCount(content []byte) int {
	lineCount := 0
	for _, b := range content {
		if b == '\n' {
			lineCount++
		}
	}

	if content[len(content)-1] != '\n' {
		lineCount++
	}

	return lineCount
}

func wordCount(content []byte) int {
	words := strings.Fields(string(content))
	return len(words)
}

func byteCount(content []byte) int {
	return len(content)
}

func characterCount(content []byte) int {
	characters := []rune(string(content))
	return len(characters)
}
