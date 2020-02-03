package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

const (
	fontsize = 16
	spacer   = 3
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "need one argument: email")
		os.Exit(1)
	}
	email := os.Args[1]
	if !strings.Contains(email, "@") {
		fmt.Fprintln(os.Stderr, "provided argument is not an email")
		os.Exit(1)
	}
	err := GeneratePdf("emailrecap.pdf", email)
	if err != nil {
		panic(err)
	}
}

func GeneratePdf(filename, email string) error {

	pdf := gofpdf.New("P", "mm", "A5", "")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", fontsize)
	lineSz := pdf.PointToUnitConvert(fontsize)
	tr := pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"

	// ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
	pdf.ImageOptions(
		"logo.png",
		10, 10,
		20, 20,
		true,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0,
		"",
	)
	pdf.MultiCell(0, lineSz+spacer, tr(""), "", "C", false) // blank line
	pdf.SetFont("Helvetica", "B", fontsize)                 // bold font
	pdf.MultiCell(0, lineSz+spacer, tr("Télétransmission des compte-rendus d'imagerie médicale"), "", "C", false)
	pdf.SetFont("Helvetica", "", fontsize-2)                // regular font
	pdf.MultiCell(0, lineSz+spacer, tr(""), "", "C", false) // blank line
	pdf.MultiCell(0, lineSz+spacer, tr("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec neque sed ante convallis rhoncus. Ut sodales, magna at efficitur placerat, dui ex accumsan mi, quis pulvinar orci magna ut nisl. Aliquam semper, orci ut pharetra molestie, erat purus vulputate ipsum, ut egestas leo enim a augue. Nulla et risus nunc. Donec massa lacus, posuere a quam in, sodales rhoncus orci. Cras eleifend, ex ut gravida rutrum, nisi ipsum congue ligula, pellentesque condimentum quam turpis in sapien."), "", "", false)
	pdf.MultiCell(0, lineSz+spacer, tr(""), "", "C", false) // blank line
	pdf.SetFont("Helvetica", "", fontsize-1)                // regular font
	pdf.MultiCell(0, lineSz+spacer, tr("Vous recevrez vos comptes-rendus d'imagerie médicale à l'adresse email suivante :"), "", "C", false)
	pdf.SetFont("Helvetica", "B", fontsize) // bold font
	pdf.MultiCell(0, lineSz+spacer, tr(email), "", "C", false)

	return pdf.OutputFileAndClose(filename)
}
