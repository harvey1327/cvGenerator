package main

import (
	"log"

	"github.com/harvey1327/resumehack/internal/latex"
	"github.com/harvey1327/resumehack/internal/model"
	"github.com/harvey1327/resumehack/internal/yml"
)

func main() {
	// Parse yml into struct
	log.Println("Converting yml into struct")
	var pageData model.PageData
	err := yml.PopulateStructWithYML("./resume.yml", &pageData)
	if err != nil {
		log.Fatalf("error converting yml into struct: %s", err.Error())
	}
	log.Println("yml successfully converted")

	// Convert struct into latex template
	log.Println("Starting template generation")
	ltx := latex.TemplateGeneration(pageData.Meta)
	err = ltx.Generate(pageData)
	if err != nil {
		log.Fatalf("error generating latex template: %s", err.Error())
	}
	log.Println("template successfully generated")

	//Convert latex template to pdf
	log.Println("Starting PDF generation")
	output, err := ltx.ConvertToPdf()
	if err != nil {
		log.Fatalf("error generating PDF: %s", err.Error())
	}
	log.Println(string(output))
	log.Println("PDF successfully generated")
}
