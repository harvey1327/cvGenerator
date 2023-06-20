package main

import (
	"fmt"
	"log"

	"github.com/harvey1327/resumehack/internal/latex"
	"github.com/harvey1327/resumehack/internal/model"
	"github.com/harvey1327/resumehack/internal/sensativedata"
	"github.com/harvey1327/resumehack/internal/yml"
)

func main() {

	resumeLocation := "./resume"

	//accept commandline flags
	arguments := sensativedata.GetData()

	// Parse yml into struct
	log.Println("Converting yml into struct")
	var pageData model.PageData
	err := yml.PopulateStructWithYML(fmt.Sprintf("%s/resume.yml", resumeLocation), &pageData)
	if err != nil {
		log.Fatalf("error converting yml into struct: %s", err.Error())
	}
	log.Println("yml successfully converted")

	log.Println("applying arguments")
	arguments.Apply(&pageData)

	// Convert struct into latex template
	log.Println("Starting template generation")
	ltx := latex.TemplateGeneration(pageData.Meta, pageData.Info.Name)
	err = ltx.Generate(pageData)
	if err != nil {
		log.Fatalf("error generating latex template: %s", err.Error())
	}
	log.Println("template successfully generated")

	//Convert latex template to pdf
	log.Println("Starting PDF generation")
	output, err := ltx.ConvertToPdf(resumeLocation)
	if err != nil {
		log.Fatalf("error generating PDF, please read log: %s", err.Error())
		log.Println(string(output))
	}
	log.Println(string(output))
	log.Println("PDF successfully generated")
}
