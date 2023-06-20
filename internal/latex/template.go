package latex

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/harvey1327/resumehack/internal/model"
)

type templateGeneration struct {
	folderPath      string
	templateName    string
	outputLatexFile string
}

func TemplateGeneration(metaData model.Meta, name string) templateGeneration {
	folderPath := fmt.Sprintf("./template/%s/%s", metaData.Format, metaData.Version)
	return templateGeneration{
		folderPath:      folderPath,
		templateName:    "template.tex",
		outputLatexFile: fmt.Sprintf("%s/%s_CV.tex", folderPath, strings.ReplaceAll(name, " ", "_")),
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

func (tg templateGeneration) ConvertToPdf(outputPath string) ([]byte, error) {
	cmd := exec.Command("pdflatex", fmt.Sprintf("--output-directory=%s", outputPath), tg.outputLatexFile)
	return cmd.Output()
}
