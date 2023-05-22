package engine

import (
	"testing"

	"github.com/mxcd/gameplan/pkg/schema"
	"github.com/stretchr/testify/assert"
)

func TestGameplanValidation(t *testing.T) {
	assert := assert.New(t)

	gameplanFilePath := "../../examples/gameplanA.yml"
	workingDirectory := "../../_work"

	gameplanInstance, err := NewGameplan(&GameplanOptions{
		WorkingDirectory: &workingDirectory,
		GameplanFilePath: gameplanFilePath,
	})
	if err != nil {
		t.Error(err)
	}

	assert.Equal("gameplanA", gameplanInstance.GameplanDefinition.Name, "Gameplan name should be gameplanA")
	println(gameplanInstance.GameplanDefinition.Name)
}

func TestGameplanExecution(t *testing.T) {
	gameplanFilePath := "../../examples/gameplanA.yml"
	workingDirectory := "../../_work"

	templateData := schema.GameplanTemplateData{
		Data: map[string]interface{}{
			"Name": "mxcd",
		},
	}

	gameplanInstance, err := NewGameplan(&GameplanOptions{
		WorkingDirectory:    &workingDirectory,
		GameplanFilePath:    gameplanFilePath,
		TemplateData:        &templateData,
		TemplateDirectories: []string{"../../examples/templates"},
	})
	if err != nil {
		t.Error(err)
	}

	err = gameplanInstance.Execute()
	if err != nil {
		t.Error(err)
	}
}
