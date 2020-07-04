package main

import (
	"fmt"

	"github.com/mostlygeek/kids-learning/lib"
)

func main() {

	pdf, _ := lib.Table54([]string{"bob", "ben"})
	err := pdf.OutputFileAndClose("cvc.pdf")
	if err != nil {
		fmt.Println(err.Error())
	}
}
