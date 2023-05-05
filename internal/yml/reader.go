package yml

import (
	"os"
	"path/filepath"

	"github.com/harvey1327/resumehack/internal/model"
	"gopkg.in/yaml.v2"
)

func PopulateStructWithYML(filePath string, pageData *model.PageData) error {
	filename, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, pageData)
	if err != nil {
		return err
	}
	return nil
}
