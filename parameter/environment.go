package parameter

import (
	"flag"
	"os"
	"strings"
)

func init() {
	envProvider := new(EnvironmentProvider)
	flag.BoolVar(&envProvider.useEnvironment, "env", false, "Use of environment")
	ParametersMapProviders = append(ParametersMapProviders, envProvider)
}

type EnvironmentProvider struct {
	useEnvironment bool
}

func (me *EnvironmentProvider) FillMap(parameterMap map[string]interface{}) error {
	if me.useEnvironment == false {
		return nil
	}
	getenvironment := func(data []string, getkeyval func(item string) (key, val string)) {
		for _, item := range data {
			key, val := getkeyval(item)
			parameterMap[key] = val
		}
	}
	getenvironment(os.Environ(), func(item string) (key, val string) {
		splits := strings.Split(item, "=")
		key = splits[0]
		val = splits[1]
		return
	})
	return nil
}
