//+build !nolog
//+build js

package dlog

import (
	"bytes"
	"fmt"

	"runtime"
	"strconv"
)

var (
	byt = bytes.NewBuffer(make([]byte, 0))
)

// dLog, the primary function of the package,
// prints out and writes to file a string
// containing the logged data separated by spaces,
// prepended with file and line information.
// It only includes logs which pass the current filters.
// Todo: use io.Multiwriter to simplify the writing to
// both logfiles and stdout
func dLog(console, override bool, in ...interface{}) {
	//(pc uintptr, file string, line int, ok bool)
	_, f, line, ok := runtime.Caller(2)
	if ok {
		// Note on errors: these functions all return
		// errors, but they are always nil.
		byt.WriteRune('[')
		byt.WriteString(f)
		byt.WriteRune(':')
		byt.WriteString(strconv.Itoa(line))
		byt.WriteString("]  ")
		for _, elem := range in {
			byt.WriteString(fmt.Sprintf("%v ", elem))
		}
		byt.WriteRune('\n')

		if console {
			fmt.Print(byt.String())
		}

		byt.Reset()
	}
}

// FileWrite runs dLog, but JUST writes to file instead
// of also to stdout.
func FileWrite(in ...interface{}) {
}

// CreateLogFile creates a file in the 'logs' directory
// of the starting point of this program to write logs to
func CreateLogFile() {
}
