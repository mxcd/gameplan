package engine

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
)

var templates = map[string]*template.Template{}

func loadTemplateDirectories(templateDirectories []string) error {

	for _, templateDirectory := range templateDirectories {
		dirInfo, err := os.Stat(templateDirectory)
		if os.IsNotExist(err) {
			return err
		} else if err != nil {
			return err
		}

		if !dirInfo.IsDir() {
			return errors.New("template directory path exists but is not a directory")
		}

		err = loadTemplatesFromDirectory(templateDirectory)
		if err != nil {
			return err
		}
	}

	return nil
}

func loadTemplatesFromDirectory(templateDirectory string) error {
	var templateFiles []string
	err := filepath.WalkDir(templateDirectory,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			templateFiles = append(templateFiles, path)
			return nil
		})
	if err != nil {
		return err
	}

	for _, templateFile := range templateFiles {
		relativePath, err := filepath.Rel(templateDirectory, templateFile)
		if err != nil {
			return err
		}
		if _, ok := templates[relativePath]; ok {
			return errors.New("duplicate template file '" + relativePath + "' name found")
		}
		templateData, err := os.ReadFile(templateFile)
		if err != nil {
			return err
		}
		tmp, err := template.New(relativePath).Parse(string(templateData))
		if err != nil {
			return err
		}
		templates[relativePath] = tmp
	}
	return nil
}

func getTemplate(templateName string) (*template.Template, error) {
	if _, ok := templates[templateName]; !ok {
		return nil, errors.New("template '" + templateName + "' not found")
	}
	return templates[templateName], nil
}
