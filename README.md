# Filler
Filler is a simple templating tool, based on [golang templates](http://golang.org/pkg/text/template/), that you can run from the command line to substitute values in a template from different sources, such as (for the moment):
* Json documents
* Environment variables
* Inline key pairs

You can for example receive Json data from external tools and use it to to programmatically adapt your configuration files, for example

`aws iam list-users | filler -jin -ts "user_arns={{range .Users}}{{.Arn}},{{end}}"`
