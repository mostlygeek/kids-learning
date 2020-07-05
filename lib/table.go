package lib

import "github.com/jung-kurt/gofpdf"

// TableLayout iterates through list and puts each string into
// a cell in a 4x12 table.  If there are multiple pages it will
// cycle through the list until all cells are filled.
func TableLayout(pages int, list []string) (*gofpdf.Fpdf, error) {

	leftMargin := 10.0
	topMargin := 8.0

	// cell sizes for table (makes a 4x12 table)
	cellWidth := 49.0
	cellHeight := 18.0

	pdf := gofpdf.New("P", "mm", "Letter", "")

	// create a table of 48 cells,  4x12
	i := 0
	l := len(list)
	for page := 1; page <= pages; page++ {
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 12)
		pdf.SetLeftMargin(leftMargin)
		pdf.SetTopMargin(topMargin)

		// add Date
		pdf.SetFont("Arial", "B", 16)
		pdf.SetXY(leftMargin+98, topMargin)
		pdf.CellFormat(98, 12, "Date: ___________________", "0", 0, "RM", false, 0, "")

		// add Time
		pdf.SetXY(leftMargin+98, topMargin+12)
		pdf.CellFormat(98, 12, "Time: ___________________", "0", 0, "RM", false, 0, "")

		pdf.SetXY(leftMargin, 40)
		pdf.SetFont("Arial", "B", 20)
		lm := 0

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
	}

	return pdf, nil
}
