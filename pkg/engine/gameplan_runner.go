package engine

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/mxcd/gameplan/pkg/schema"
)

func (g *GameplanInstance) Execute() error {
	workingDirectoryInfo, err := os.Stat(g.WorkingDirectory)
	if os.IsNotExist(err) {
		err := os.MkdirAll(g.WorkingDirectory, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else if !workingDirectoryInfo.IsDir() {
		return errors.New("working directory path exists but is not a directory")
	}

	for _, step := range g.GameplanDefinition.Steps {
		switch step.Type {
		case schema.GameplanStepTypeCommand:
			err := executeCommand(g, step)
			if err != nil {
				return err
			}
		case schema.GameplanStepTypeTemplate:
			err := executeTemplate(g, step)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func executeCommand(g *GameplanInstance, step schema.GameplanStep) error {
	// deconstruct command string into command and arguments
	// TODO: allow for templating of command string with data
	commandComponents := strings.Split(step.CommandOptions.Command, " ")
	mainCommand := commandComponents[0]
	commandArguments := commandComponents[1:]

	cmd := exec.Command(mainCommand, commandArguments...)

	cmd.Dir = g.WorkingDirectory

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func executeTemplate(g *GameplanInstance, step schema.GameplanStep) error {
	template, err := getTemplate(step.TemplateOptions.Template)
	if err != nil {
		return err
	}

	destinationFilePath := path.Join(g.WorkingDirectory, step.TemplateOptions.Destination)

	writer, err := os.Create(destinationFilePath)
	if err != nil {
		return err
	}
	defer writer.Close()

	template.Execute(writer, g.TemplateData)

	return nil
}
