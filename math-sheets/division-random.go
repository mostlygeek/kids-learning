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

	pdf := lib.NewPDF()

	for page := 0; page < 5; page++ {
		list := make([]string, 48, 48)

		// generate equations like X x _____ = Z
		// ensure that there are no duplicate equations too close to each other
		// the table generated is 4 x 12 ... make sure the previous value is not
		// the same, or the value above
		i := 0

	fillLoop:
		for {
			if i == 48 {
				break
			}

			numerator := 2 + rand.Intn(8)
			denominator := 2 + rand.Intn(8)

			if numerator == denominator {
				continue fillLoop
			}

			cell := fmt.Sprintf("%d / %d = ____", numerator*denominator, denominator)
			// make sure there are no duplicates in the last four
			if i > 0 {
				for l := 1; l <= 4; l++ {
					if i-l < 0 {
						break
					}

					if list[i-l] == cell {
						continue fillLoop
					}

				}
			}

			list[i] = cell
			i++
		}
		lib.AddTablePage(pdf, list)
	}

	if err := pdf.OutputFileAndClose("output-division-random.pdf"); err != nil {
		fmt.Println(err.Error())
	}
}
