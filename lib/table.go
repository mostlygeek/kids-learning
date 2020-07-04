package lib

import "github.com/jung-kurt/gofpdf"

func Table54(list []string) (*gopdf.Fpdf, error) {
	leftMargin := 10.0
	topMargin := 8.0

	pdf := gofpdf.New("P", "mm", "Letter", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.SetLeftMargin(leftMargin)
	pdf.SetTopMargin(topMargin)

	// add Name
	//pdf.SetFont("Arial", "B", 24)
	//pdf.Text(leftMargin, 24.0, "Name: _____________")

	// add Date
	pdf.SetFont("Arial", "B", 16)
	pdf.SetXY(leftMargin+98, topMargin)
	pdf.CellFormat(98, 12, "Date: ___________________", "0", 0, "RM", false, 0, "")
	pdf.SetXY(leftMargin+98, topMargin+12)
	pdf.CellFormat(98, 12, "Time: ___________________", "0", 0, "RM", false, 0, "")
	// add Time

	// create a table 4x18 (54) CVC words
	for row := 0; row < 18; row++ {
		for col := 0; col < 4; col++ {

		}
	}

	// draw table of cvc words
	cellWidth := 49.0
	cellHeight := 12.15

	pdf.SetXY(leftMargin, 40)
	pdf.SetFont("Arial", "B", 15)
	lm := 0
	for row := 0; row < 18; row++ {
		for col := 0; col < 4; col++ {
			if col < 3 {
				lm = 0
			} else {
				lm = 1
			}

			pdf.CellFormat(cellWidth, cellHeight, "CVC", "1", lm, "CM", false, 0, "")
		}
	}

	return pdf, nil
}
