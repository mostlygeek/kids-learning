package main

// create a multiplication / division test for addison

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

	pdf := lib.NewPDF()

	list := make([]string, 48, 48)

	// multiplication test
	i := 0
	for {
		if i == 48 {
			break
		}
		cell := fmt.Sprintf("%d x %d = ____", 2+rand.Intn(10), 2+rand.Intn(10))

		// make sure there are no duplicates in the last four
		if i > 0 {
			for l := 1; l <= 4; l++ {
				if i-l < 0 {
					break
				}

				if list[i-l] == cell {
					continue
				}

			}
		}

		list[i] = cell
		i++
	}
	lib.AddTablePage(pdf, list)

	// division test
	i = 0
	for {
		if i == 48 {
			break
		}
		numerator := 2 + rand.Intn(8)
		denominator := 2 + rand.Intn(8)
		answer := numerator * denominator
		cell := fmt.Sprintf("%d / %d = ____", answer, denominator)

		// make sure there are no duplicates in the last four
		if i > 0 {
			for l := 1; l <= 4; l++ {
				if i-l < 0 {
					break
				}

				if list[i-l] == cell {
					continue
				}

			}
		}

		list[i] = cell
		i++
	}
	lib.AddTablePage(pdf, list)

	if err := pdf.OutputFileAndClose("multiply-test.pdf"); err != nil {
		fmt.Println(err.Error())
	}
}
