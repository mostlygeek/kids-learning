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

	// create [0..10] + [1..9] sheet, to make adding 0 to 19 memory
	// 10 pages of [0..10] + rand
	//  5 pages of rand + rand (for practice)

	for start := 0; start <= 10; start++ {
		list := make([]string, 48, 48)

		// generate equations like X + Y = _____
		// ensure that there are no duplicate equations too close to each other
		// the table generated is 4 x 12 ... make sure the previous value is not
		// the same, or the value above
		i := 0

	fillLoop:
		for {
			if i == 48 {
				break
			}
			cell := fmt.Sprintf("%d + %d = ____", start, 1+rand.Intn(9))

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

	for page := 0; page < 5; page++ {
		list := make([]string, 48, 48)
		i := 0
	fillLoop2:
		for {
			if i == 48 {
				break
			}
			cell := fmt.Sprintf("%d + %d = ____", rand.Intn(10), 1+rand.Intn(9))

			// make sure there are no duplicates in the last four
			if i > 0 {
				for l := 1; l <= 4; l++ {
					if i-l < 0 {
						break
					}

					if list[i-l] == cell {
						continue fillLoop2
					}

				}
			}

			list[i] = cell
			i++
		}
		lib.AddTablePage(pdf, list)
	}

	if err := pdf.OutputFileAndClose("output-add-to-20.pdf"); err != nil {
		fmt.Println(err.Error())
	}
}
