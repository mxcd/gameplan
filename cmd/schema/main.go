package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"os"

	"github.com/invopop/jsonschema"
	schema "github.com/mxcd/gameplan/pkg/schema"
)

func main() {
	var outputFilePath string
	flag.StringVar(&outputFilePath, "o", "schema.json", "Output file path")

	jsonSchema := jsonschema.Reflect(&schema.GameplanDefinition{})
	data, err := jsonSchema.MarshalJSON()
	if err != nil {
		panic(err)
	}

	var obj map[string]interface{}
	json.Unmarshal(data, &obj)

	indentedData, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err)
	}

	file, err := os.Create("gameplan-schema.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	w.Write(indentedData)
	w.Flush()
}
