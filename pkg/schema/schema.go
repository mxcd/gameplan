package schema

type GameplanDefinition struct {
	Version     string         `json:"version" yaml:"version"`
	Name        string         `json:"name" yaml:"name"`
	Description string         `json:"description" yaml:"description"`
	Steps       []GameplanStep `json:"steps" yaml:"steps"`
}

type GameplanStepType string

const (
	GameplanStepTypeCommand  GameplanStepType = "command"
	GameplanStepTypeTemplate GameplanStepType = "template"
)

type GameplanStep struct {
	Name            string                       `json:"name,omitempty" yaml:"name,omitempty"`
	Description     string                       `json:"description,omitempty" yaml:"description,omitempty"`
	Type            GameplanStepType             `json:"type" yaml:"type"`
	CommandOptions  *GameplanStepCommandOptions  `json:"commandOptions,omitempty" yaml:"commandOptions,omitempty"`
	TemplateOptions *GameplanStepTemplateOptions `json:"templateOptions,omitempty" yaml:"templateOptions,omitempty"`
}

type GameplanStepCommandOptions struct {
	Command string `json:"command,omitempty" yaml:"command,omitempty"`
}

type GameplanStepTemplateOptions struct {
	Template    string `json:"template,omitempty" yaml:"template,omitempty"`
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`
}

type GameplanTemplateData struct {
	Data map[string]interface{}
}
