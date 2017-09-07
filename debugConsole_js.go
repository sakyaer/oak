//+build js

package oak

import (
	"fmt"
	"io"
)

var (
	viewportLocked = false
)

func AddCommand(s string, fn func([]string)) {
	AddCheat(s, fn)
}

func AddCheat(s string, fn func([]string)) {}

func defaultDebugConsole() {
	fmt.Println("Ditching JS debug console")
}

func debugConsole(resetCh, skipScene chan bool, input io.Reader) {}
