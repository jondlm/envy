package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/jawher/mow.cli"
)

// This gets set by `go build -ldflags "-X main.version=1.0.0"`
var version string

func main() {
	app := cli.App("envy", "template a single SRC go template file, using environment variables, to DST")
	app.Spec = "[-f] SRC DST"

	app.Version("v version", version)

	var (
		sourcePath      = app.StringArg("SRC", "", "Source path of your go template file")
		destinationPath = app.StringArg("DST", "", "Destination path for the resulting, templated file")
		force           = app.BoolOpt("f force", false, "Force overwrite the DST file even if it already exists")
	)

	app.Action = func() {
		err := TemplateFile(sourcePath, destinationPath, force)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	}

	app.Run(os.Args)
}

// TemplateFile takes a go template file at `sourcePath` and writes the result
// to `destinationPath`
func TemplateFile(sourcePath *string, destinationPath *string, force *bool) error {
	if _, err := os.Stat(*sourcePath); os.IsNotExist(err) {
		return fmt.Errorf("\"%s\" was not found", *sourcePath)
	}

	if !*force {
		if _, err := os.Stat(*destinationPath); err == nil {
			return fmt.Errorf("\"%s\" already exists", *destinationPath)
		}
	}

	dest, err := os.Create(*destinationPath)
	if err != nil {
		return err
	}

	defer dest.Close()

	tpl, err := template.ParseFiles(*sourcePath)
	if err != nil {
		return err
	}

	// Don't allow missing keys
	tpl.Option("missingkey=error")

	var values = make(map[string]interface{})
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		values[parts[0]] = parts[1]
	}

	err = tpl.Execute(dest, values)
	if err != nil {
		return err
	}

	fmt.Printf("File written to `%v`\n", *destinationPath)
	return nil
}
