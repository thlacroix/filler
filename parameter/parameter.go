package parameter

import (
	"os"
	"text/template"
)

var ParametersMapProviders = make([]ParametersMapProvider, 0)

type ParametersMapProvider interface {
	FillMap(map[string]interface{}) error
}

type SliceParametersProvider interface {
	FillSlice([]interface{}) error
}

type Parameters interface {
	ProcessProviders() error
	ExecuteTemplate(*template.Template) error
}

type ParameterMap struct {
	data map[string]interface{}
}

func NewParameterMap() *ParameterMap {
	return &ParameterMap{make(map[string]interface{})}
}

func (m *ParameterMap) processProvider(p ParametersMapProvider) error {
	return p.FillMap(m.data)
}

func (m *ParameterMap) ProcessProviders() error {
	for _, p := range ParametersMapProviders {
		if err := m.processProvider(p); err != nil {
			return err
		}
	}
	return nil
}

func (m *ParameterMap) ExecuteTemplate(t *template.Template) error {
	return t.Execute(os.Stdout, m.data)
}
