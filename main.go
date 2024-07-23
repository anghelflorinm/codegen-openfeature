package main

import (
	generator "codegen/generators"
	"codegen/generators/golang"
	_ "embed"
	"fmt"
)

func generate(gen generator.Generator) error {
	data := &generator.BaseTmplData{
		OutputDir: ".",
		Flags: []*generator.FlagTmplData{
			{
				Name:         "my_open_feature_flag",
				Type:         generator.BoolType,
				DefaultValue: "true",
				Docs:         "A flag for Open Feature.",
			},
		},
	}
	input := generator.Input{
		BaseData: data,
	}
	return gen.Generate(input)
}

func main() {
	params := golang.Params{
		File:      "./test_folder/experiment_plags.go",
		GoPackage: "experimentflags",
	}
	gen := golang.NewGenerator(params)
	err := generate(gen)
	if err != nil {
		fmt.Printf("error generating flag accesssors: %v", err)
	}
}
