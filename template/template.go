package template

import (
	"flag"
	"io/ioutil"
	"text/template"
)

var templatePath, templateString string

func init() {
	flag.StringVar(&templatePath, "t", "", "Path to the template")
	flag.StringVar(&templateString, "ts", "", "Template string")
}

func GetTemplate() (*template.Template, error) {
	if templatePath != "" {
		bytes, err := ioutil.ReadFile(templatePath)
		if err != nil {
			return nil, err
		}
		templateString = string(bytes)
	}
	return template.New("File").Parse(string(templateString))
}
