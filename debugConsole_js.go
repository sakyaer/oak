//+build js

package oak

var (
	viewportLocked = false
)

func AddCommand(s string, fn func([]string)) {
	AddCheat(s, fn)
}

func AddCheat(s string, fn func([]string)) {}

func debugConsole(resetCh, skipScene chan bool) {}
