package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mostlygeek/kids-learning/lib"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// generate 5 pages of random cvc words
	pages := 5
	listSize := pages * 48 /* 48 cells in the table */
	list := make([]string, listSize, listSize)

	min := 2
	max := 9

	// generate equations like X x Y = _____
	for i := 0; i < listSize; i++ {
		list[i] = fmt.Sprintf("%d x %d = ____", min+rand.Intn(max-min+1), min+rand.Intn(max-min+1))
	}

	pdf, err := lib.TableLayout(pages, list)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	if err := pdf.OutputFileAndClose("output.pdf"); err != nil {
		fmt.Println(err.Error())
	}
}
