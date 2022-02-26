package main

import (
	"fmt"
	"strings"
)

const (
	// StupidYaml2GoStructName yaml2go library use it as root struct name
	StupidYaml2GoStructName = "Yaml2Go"

	// packageModel
	packageModel = "package %s\n\n"
)

func StructReplace(structContent, structName, packageName string) string {
	result := strings.ReplaceAll(structContent, StupidYaml2GoStructName, structName)
	return fmt.Sprintf(packageModel, packageName) + result
}
