package engine

import (
	"os"

	"github.com/mxcd/gameplan/pkg/schema"
	"gopkg.in/yaml.v3"
)

type GameplanInstance struct {
	WorkingDirectory    string
	GameplanFilePath    string
	GameplanDefinition  *schema.GameplanDefinition
	TemplateData        *schema.GameplanTemplateData
	TemplateDirectories []string
}

type GameplanOptions struct {
	WorkingDirectory    *string
	GameplanFilePath    string
	TemplateData        *schema.GameplanTemplateData
	TemplateDirectories []string
}

func NewGameplan(options *GameplanOptions) (*GameplanInstance, error) {

	workingDirectory := options.WorkingDirectory
	if *workingDirectory == "" || *workingDirectory == "." || workingDirectory == nil {
		path, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		workingDirectory = &path
	}

	gameplanInstance := &GameplanInstance{
		WorkingDirectory: *workingDirectory,
		GameplanFilePath: options.GameplanFilePath,
		TemplateData:     options.TemplateData,
	}

	err := loadTemplateDirectories(options.TemplateDirectories)
	if err != nil {
		return nil, err
	}

	err = gameplanInstance.Validate()
	if err != nil {
		return nil, err
	}

	err = gameplanInstance.Load()
	if err != nil {
		return nil, err
	}

	return gameplanInstance, nil
}

func (g *GameplanInstance) Load() error {
	gameplanDefinition := &schema.GameplanDefinition{}

	yamlText, err := os.ReadFile(g.GameplanFilePath)
	if err != nil {
		return err
	}

	var yamlData interface{}
	err = yaml.Unmarshal(yamlText, &yamlData)
	if err != nil {
		return err
	}
  
	err = yaml.Unmarshal(yamlText, &gameplanDefinition)
	if err != nil {
		return err
	}

	g.GameplanDefinition = gameplanDefinition
	return nil
}

func (g *GameplanInstance) SetWorkingDirectory(workingDirectory string) {
	g.WorkingDirectory = workingDirectory
}
