package latex

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"github.com/harvey1327/resumehack/internal/model"
)

type templateGeneration struct {
	folderPath      string
	templateName    string
	outputLatexFile string
}

func TemplateGeneration(metaData model.Meta) templateGeneration {
	folderPath := fmt.Sprintf("./template/%s/%s", metaData.Format, metaData.Version)
	return templateGeneration{
		folderPath:      folderPath,
		templateName:    "template.tex",
		outputLatexFile: fmt.Sprintf("%s/cv.tex", folderPath),
	}
}

func (tg templateGeneration) Generate(pageData model.PageData) error {
	t := template.Must(template.New(tg.templateName).Delims("[[", "]]").ParseFiles(fmt.Sprintf("%s/%s", tg.folderPath, tg.templateName)))
	output, err := os.Create(tg.outputLatexFile)
	if err != nil {
		return err
	}
	err = t.Execute(output, pageData)
	if err != nil {
		return err
	}
	return nil
}

func (tg templateGeneration) ConvertToPdf() ([]byte, error) {
	cmd := exec.Command("pdflatex", fmt.Sprintf("--output-directory=%s", tg.folderPath), tg.outputLatexFile)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}
