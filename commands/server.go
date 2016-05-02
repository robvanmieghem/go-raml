package commands

import (
	"github.com/Jumpscale/go-raml/codegen"

	log "github.com/Sirupsen/logrus"
)

// ServerCommand is executed to generate a go server from a RAML specification
type ServerCommand struct {
	Language         string // target language
	Dir              string //target dir
	RamlFile         string //raml file
	PackageName      string //package name in the generated go source files
	NoMainGeneration bool   //do not generate a main.go file
}

// Execute generates a Go server from an RAML specification
func (command *ServerCommand) Execute() error {
	log.Infof("Generating a %v server", command.Language)
	return codegen.GenerateServer(command.RamlFile, command.Dir, command.PackageName,
		command.Language, "apidocs", !command.NoMainGeneration)
}
