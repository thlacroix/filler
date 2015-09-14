package main

import (
	"flag"
	"log"

	"github.com/thlacroix/filler/parameter"
	"github.com/thlacroix/filler/template"
)

func main() {
	flag.Parse()
	var p parameter.Parameters = parameter.NewParameterMap()
	if err := p.ProcessProviders(); err != nil {
		log.Fatal(err)
	}
	t, err := template.GetTemplate()
	if err != nil {
		log.Fatal(err)
	}
	if err := p.ExecuteTemplate(t); err != nil {
		log.Fatal(err)
	}
}
