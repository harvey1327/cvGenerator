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
	templatePath := fmt.Sprintf("./template/%s/%s/template.tex", pageData.Meta.Format, pageData.Meta.Version)
	t := template.Must(template.New("test.template").Delims("[[", "]]").ParseFiles("test.template"))
	err = t.Execute(os.Stdout, pageData)
	if err != nil {
		panic(err)
	}
}
