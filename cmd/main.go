package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/harvey1327/resumehack/internal/model"
	"github.com/harvey1327/resumehack/internal/yml"
)

func main() {
	// Parse yml into struct
	var pageData model.PageData
	err := yml.PopulateStructWithYML("./resume.yml", &pageData)
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", pageData)

	// Convert struct into template
	templateName := "template.tmpl"
	templatePath := fmt.Sprintf("./template/%s/%s/%s", pageData.Meta.Format, pageData.Meta.Version, templateName)
	t := template.Must(template.New(templateName).Delims("[[", "]]").ParseFiles(templatePath))

	output, err := os.Create(fmt.Sprintf("./template/%s/%s/cv.tex", pageData.Meta.Format, pageData.Meta.Version))
	if err != nil {
		panic(err)
	}
	err = t.Execute(output, pageData)
	if err != nil {
		panic(err)
	}
}
