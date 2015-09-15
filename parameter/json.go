package parameter

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
)

func init() {
	jsonProvider := new(JsonProvider)
	flag.BoolVar(&jsonProvider.jsonStdIn, "jin", false, "Read json from stdin")
	flag.StringVar(&jsonProvider.jsonPath, "j", "", "Read json from file")
	flag.StringVar(&jsonProvider.jsonString, "js", "", "Json as string")
	ParametersMapProviders = append(ParametersMapProviders, jsonProvider)
	ParametersSliceProviders = append(ParametersSliceProviders, jsonProvider)
}

type JsonProvider struct {
	jsonStdIn  bool
	jsonPath   string
	jsonString string
}

func (me *JsonProvider) FillMap(m map[string]interface{}) error {
	if me.jsonPath != "" {
		bytes, err := ioutil.ReadFile(me.jsonPath)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(bytes, &m); err != nil {
			return err
		}
	}
	if me.jsonStdIn {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(bytes, &m); err != nil {
			return err
		}
	}
	if me.jsonString != "" {
		if err := json.Unmarshal([]byte(me.jsonString), &m); err != nil {
			return err
		}
	}
	return nil
}

func (me *JsonProvider) FillSlice(s *[]interface{}) error {
	if me.jsonPath != "" {
		bytes, err := ioutil.ReadFile(me.jsonPath)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(bytes, s); err != nil {
			return err
		}
	}
	if me.jsonStdIn {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(bytes, s); err != nil {
			return err
		}
	}
	if me.jsonString != "" {
		if err := json.Unmarshal([]byte(me.jsonString), s); err != nil {
			return err
		}
	}
	return nil
}
