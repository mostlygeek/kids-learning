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

	// create a 2 x [1 -> 10], 3 x [1 -> 10], ... 9 x [1 -> 10]
	for start := 2; start <= 9; start++ {
		for page := 0; page < 5; page++ {
			list := make([]string, 48, 48)

			// generate equations like X x Y = _____
			// ensure that there are no duplicate equations too close to each other
			// the table generated is 4 x 12 ... make sure the previous value is not
			// the same, or the value above
			i := 0

		fillLoop:
			for {
				if i == 48 {
					break
				}
				cell := fmt.Sprintf("%d x %d = ____", start, 1+rand.Intn(10))

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
	}

	if err := pdf.OutputFileAndClose("output.pdf"); err != nil {
		fmt.Println(err.Error())
	}
}
