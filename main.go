package main

import (
	"github.com/PrasadG193/yaml2go"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "yaml2go-cli",
		Short: "yaml2go-cli is a cli-tool for yaml to go struct",
		Run: func(cmd *cobra.Command, args []string) {
			baFile, err := ioutil.ReadFile(*inputFile)
			if err != nil {
				panic(err)
			}

			y2g := yaml2go.New()
			strStruct, err := y2g.Convert(StupidYaml2GoStructName, baFile)
			if err != nil {
				panic(err)
			}

			strStruct = StructReplace(strStruct, *structName, *packageName)
			err = ioutil.WriteFile(*outputFile, []byte(strStruct), 0644)
			if err != nil {
				panic(err)
			}

			// adapter for unix
			err = os.Chmod(*outputFile, 0644)
			if err != nil {
				panic(err)
			}
		},
	}

	inputFile   *string
	outputFile  *string
	structName  *string
	packageName *string
)

func init() {
	inputFile = rootCmd.PersistentFlags().StringP("input", "i", "", "input yaml file path")
	outputFile = rootCmd.PersistentFlags().StringP("output", "o", "", "output go file path")
	structName = rootCmd.PersistentFlags().StringP("struct", "s", "Default", "struct name")
	packageName = rootCmd.PersistentFlags().StringP("package", "p", "main", "package name")

	err := rootCmd.MarkPersistentFlagRequired("input")
	if err != nil {
		panic(err)
	}

	err = rootCmd.MarkPersistentFlagRequired("output")
	if err != nil {
		panic(err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
