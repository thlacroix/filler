package parameter

import (
	"bufio"
	"flag"
	"os"
)

func init() {
	listinProvider := new(ListinProvider)
	flag.BoolVar(&listinProvider.useStdIn, "lin", false, "Use of stdin for list input")
	ParametersSliceProviders = append(ParametersSliceProviders, listinProvider)
}

type ListinProvider struct {
	useStdIn bool
}

func (me *ListinProvider) FillSlice(s *[]interface{}) error {
	if me.useStdIn {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			*s = append(*s, scan.Text())
		}
		if err := scan.Err(); err != nil {
			return err
		}
	}
	return nil
}
