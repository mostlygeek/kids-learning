package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/mostlygeek/kids-learning/lib"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Error:  Needed output filename")
		return
	}

	outputFile := os.Args[1]

	wordlist := make([]string, 0, 48)
	reader := bufio.NewReader(os.Stdin)
	count := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		wordlist = append(wordlist, line)
		count++

		if count == 48 {
			break
		}
	}

	num := len(wordlist)
	switch {
	case num == 0:
		fmt.Println("Error: No words provided on stdin")
		return
	case num < 48:
		fmt.Printf("Warning: Only %d words provided, 48 or more is optimal\n", num)
	}

	pdf := lib.NewPDF()

	for page := 0; page < 5; page++ {
		lib.Shuffle(wordlist)
		lib.AddTablePage(pdf, wordlist)
	}

	if err := pdf.OutputFileAndClose(outputFile); err != nil {
		fmt.Println(err.Error())
	}
}
