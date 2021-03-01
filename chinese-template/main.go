package main

import (
	"fmt"

	"github.com/mostlygeek/kids-learning/lib"
)

const (
	letterWidth  = 216.0
	letterHeight = 279.0
)

// generate a template for writing chinese characters
func main() {

	pdf := lib.NewPDF()
	leftMargin := 12.0
	topMargin := 12.0

	boundWidth := letterWidth - (2 * leftMargin) - 15
	boundHeight := letterHeight - (2 * topMargin)

	// draw outside box
	square := 24.0      // dimensions of the square
	inner := square / 3 // for inner lines

	numCol := int(boundWidth / square)
	numRow := int(boundHeight / square)

	// set a box's left/top margin so they are nicely spaced and don't touch
	lMargin := (boundWidth - (square * float64(numCol))) / float64(numCol)
	tMargin := (boundHeight - (square * float64(numRow))) / float64(numRow)

	fmt.Println("Each sheet has :", numRow*numCol, "squares")

	for page := 0; page < 2; page++ {
		pdf.AddPage()
		pdf.SetLeftMargin(leftMargin)
		pdf.SetTopMargin(topMargin)

		for row := 0; row < numRow; row++ {
			for col := 0; col < numCol; col++ {

				frow := float64(row)
				fcol := float64(col)

				top := topMargin + frow*square + tMargin*frow
				left := leftMargin + fcol*square + lMargin*fcol

				pdf.SetLineWidth(0.1)
				pdf.SetDrawColor(150, 150, 150)
				pdf.Line(left+inner, top, left+inner, top+square)
				pdf.Line(left+(2*inner), top, left+(2*inner), top+square)

				pdf.Line(left, top+inner, left+square, top+inner)
				pdf.Line(left, top+2*inner, left+square, top+2*inner)

				pdf.SetLineWidth(0.3)
				pdf.SetDrawColor(0, 0, 0)

				pdf.Rect(left, top, square, square, "")

				//
			}
		}

		// draw top/bottom lines on side of page for kids to write their name
		pdf.SetDrawColor(0, 0, 0)
		pdf.Line(letterWidth-leftMargin+4, topMargin, letterWidth-leftMargin+4, letterHeight-topMargin)
		pdf.Line(letterWidth-leftMargin-10, topMargin, letterWidth-leftMargin-10, letterHeight-topMargin)

		// middle dashed line
		pdf.SetDrawColor(150, 150, 150)
		dashLength := 2
		numDashes := int(letterHeight-topMargin) / dashLength
		fmt.Println(numDashes)
		for num := 0; num < numDashes; num++ {
			if num%2 != 0 {
				continue // white space
			}

			start := topMargin + float64(num*dashLength)
			end := start + float64(dashLength)

			// guard against drawing too long of a dashed line
			if end > letterHeight-topMargin {
				break
			}

			pdf.Line(letterWidth-leftMargin-3, start, letterWidth-leftMargin-3, end)
		}
	}

	if err := pdf.OutputFileAndClose("chinese-template.pdf"); err != nil {
		fmt.Println(err.Error())
	}

}
