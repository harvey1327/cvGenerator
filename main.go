package main

import (
	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.SetFont("wts11", "", 14)
	pdf.AddPage()
	pdf.Cell(nil, "something")
	pdf.WritePdf("hello.pdf")
}
