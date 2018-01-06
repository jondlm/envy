package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlTemplate(t *testing.T) {
	type io struct {
		EnvVars  map[string]string
		Template string
		Output   string
	}

	tests := []io{
		io{
			EnvVars:  map[string]string{"ENVY_TEST_VALUE1": "Darkness, my old friend."},
			Template: "Hello {{ .ENVY_TEST_VALUE1 }}",
			Output:   "Hello Darkness, my old friend.",
		},
		io{
			EnvVars:  map[string]string{"ENVY_TEST_VALUE2": "que no"},
			Template: "Por {{ .ENVY_TEST_VALUE2 }}?",
			Output:   "Por que no?",
		},
	}

	for _, test := range tests {

		for key, value := range test.EnvVars {
			err := os.Setenv(key, value)
			assert.Nil(t, err)
		}

		sourceFile, err := ioutil.TempFile("", "")
		assert.Nil(t, err)

		destFile, err := ioutil.TempFile("", "")
		assert.Nil(t, err)

		defer os.Remove(sourceFile.Name())
		defer os.Remove(destFile.Name())

		_, err = sourceFile.WriteString(test.Template)
		assert.Nil(t, err)
		sourceFile.Close()

		sourceName := sourceFile.Name()
		destName := destFile.Name()
		force := true

		err = TemplateFile(&sourceName, &destName, &force)
		assert.Nil(t, err)

		result, err := ioutil.ReadFile(destName)
		assert.Nil(t, err)

		assert.Equal(t, test.Output, string(result))
	}
}
