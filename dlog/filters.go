//+build !nolog

package dlog

import (
	"fmt"
	"strings"
)

var (
	debugFilter = ""
)

func checkFilter(f string, in ...interface{}) bool {
	ret := false
	for _, elem := range in {
		ret = ret || strings.Contains(fmt.Sprintf("%s", elem), debugFilter)
	}
	return ret || strings.Contains(f, debugFilter)
}

// SetDebugFilter sets the string which determines
// what debug messages get printed. Only messages
// which contain the filer as a pseudo-regex
func SetDebugFilter(filter string) {
	debugFilter = filter
}
