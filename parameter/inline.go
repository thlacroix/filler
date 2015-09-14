package parameter

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

func init() {
	inlineProvider := &InlineProvider{make(map[string]string)}
	flag.Var(inlineProvider.kp, "v", "Key pair parameter, key=pair")
	ParametersMapProviders = append(ParametersMapProviders, inlineProvider)
}

type keyPair map[string]string

type InlineProvider struct {
	kp keyPair
}

func (kp keyPair) Set(v string) error {
	values := strings.SplitN(v, "=", 2)
	if len(values) != 2 {
		return errors.New(fmt.Sprint("Wrong format for variable ", v))
	}
	kp[values[0]] = values[1]
	return nil
}

func (kp keyPair) String() string {
	return fmt.Sprint(map[string]string(kp))
}
func (me *InlineProvider) FillMap(m map[string]interface{}) error {
	for k, v := range me.kp {
		m[k] = v
	}
	return nil
}
