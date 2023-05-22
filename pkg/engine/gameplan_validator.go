package engine

import (
	"fmt"
	"os"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
)

func (g *GameplanInstance) Validate() error {

	yamlText, err := os.ReadFile(g.GameplanFilePath)
	if err != nil {
		return err
	}

	var yamlData interface{}
	err = yaml.Unmarshal(yamlText, &yamlData)
	if err != nil {
		return err
	}

	schemaData, err := os.ReadFile("../../gameplan-schema.json")
	if err != nil {
		return err
	}
	schemaString := string(schemaData)

	compiler := jsonschema.NewCompiler()
	err = compiler.AddResource("gameplan-schema.json", strings.NewReader(schemaString))
	if err != nil {
		return err
	}

	schema, err := compiler.Compile("gameplan-schema.json")
	if err != nil {
		return err
	}

	if err := schema.Validate(yamlData); err != nil {
		return err
	}

	fmt.Println("validation successfull")

	return nil
}
