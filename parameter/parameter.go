package parameter

import (
	"flag"
	"os"
	"text/template"
)

func init() {
	flag.BoolVar(&useList, "l", false, "Input is a list instead of a map")
}

var useList bool = false
var ParametersMapProviders = make([]ParametersMapProvider, 0)
var ParametersSliceProviders = make([]ParametersSliceProvider, 0)

type ParametersMapProvider interface {
	FillMap(map[string]interface{}) error
}

type ParametersSliceProvider interface {
	FillSlice(*[]interface{}) error
}

type Parameters interface {
	ProcessProviders() error
	ExecuteTemplate(*template.Template) error
}

type ParameterMap struct {
	data map[string]interface{}
}

type ParameterSlice struct {
	data []interface{}
}

func GetParameters() Parameters {
	if useList == false {
		return NewParameterMap()
	} else {
		return NewParameterSlice()
	}
}

func NewParameterMap() *ParameterMap {
	return &ParameterMap{make(map[string]interface{})}
}

func NewParameterSlice() *ParameterSlice {
	return &ParameterSlice{make([]interface{}, 0)}
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

func (s *ParameterSlice) processProvider(p ParametersSliceProvider) error {
	return p.FillSlice(&s.data)
}

func (s *ParameterSlice) ProcessProviders() error {
	for _, p := range ParametersSliceProviders {
		if err := s.processProvider(p); err != nil {
			return err
		}
	}
	return nil
}

func (s *ParameterSlice) ExecuteTemplate(t *template.Template) error {
	return t.Execute(os.Stdout, s.data)
}
