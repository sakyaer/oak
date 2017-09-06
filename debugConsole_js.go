//+build js

package oak

var (
	viewportLocked = false
)

func AddCommand(s string, fn func([]string)) {
	AddCheat(s, fn)
}

func AddCheat(s string, fn func([]string)) {}

func defaultDebugConsole() {}

func debugConsole(resetCh, skipScene chan bool, input io.Reader) {}
