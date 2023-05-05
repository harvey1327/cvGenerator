package latex

import (
	"fmt"
	"os"
	"text/template"

	"github.com/harvey1327/resumehack/internal/model"
)

type templateGeneration struct {
	templateFolder string
	templateName   string
	outputName     string
}

func TemplateGeneration() templateGeneration {
	return templateGeneration{
		templateFolder: "template",
		templateName:   "template.tmpl",
		outputName:     "cv.tex",
	}
}

func (tg templateGeneration) Generate(pageData model.PageData) error {
	folderPath := fmt.Sprintf("./%s/%s/%s", tg.templateFolder, pageData.Meta.Format, pageData.Meta.Version)
	inputPath := fmt.Sprintf("%s/%s", folderPath, tg.templateName)
	outputPath := fmt.Sprintf("%s/%s", folderPath, tg.outputName)

	t := template.Must(template.New(tg.templateName).Delims("[[", "]]").ParseFiles(inputPath))

	output, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	err = t.Execute(output, pageData)
	if err != nil {
		return err
	}
	return nil
}
