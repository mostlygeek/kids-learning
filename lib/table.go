package lib

import (
	"math/rand"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewPDF() *gofpdf.Fpdf {
	return gofpdf.New("P", "mm", "Letter", "")
}

// Adds a page with a 4x12 table, where each item in list is its own cell
func AddTablePage(pdf *gofpdf.Fpdf, list []string) *gofpdf.Fpdf {
	leftMargin := 10.0
	topMargin := 8.0

	// cell sizes for table (makes a 4x12 table)
	cellWidth := 49.0
	cellHeight := 18.0

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.SetLeftMargin(leftMargin)
	pdf.SetTopMargin(topMargin)

	pdf.SetFont("Arial", "B", 16)
	// add Name
	pdf.SetXY(leftMargin, topMargin+12)
	pdf.CellFormat(98, 12, "Name: _________________________", "0", 0, "RM", false, 0, "")

	// add Date
	pdf.SetXY(leftMargin+98, topMargin)
	pdf.CellFormat(98, 12, "Date: ___________________", "0", 0, "RM", false, 0, "")

	// add Time
	pdf.SetXY(leftMargin+98, topMargin+12)
	pdf.CellFormat(98, 12, "Time: ___________________", "0", 0, "RM", false, 0, "")

	pdf.SetXY(leftMargin, 40)
	pdf.SetFont("Arial", "B", 20)
	lm := 0

	i := 0
	l := len(list)
	for row := 0; row < 12; row++ {
		for col := 0; col < 4; col++ {
			if col < 3 {
				lm = 0
			} else {
				lm = 1
			}

			pdf.CellFormat(cellWidth, cellHeight, list[i%l], "1", lm, "CM", false, 0, "")
			i++
		}
	}

	return pdf
}

// Shuffle randomizes order of input list
func Shuffle(list []string) {
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
}
