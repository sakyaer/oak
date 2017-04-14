//+build js

package oak

func AddCommand(s string, fn func([]string)) {
	AddCheat(s, fn)
}

func AddCheat(s string, fn func([]string)) {}

func DebugConsole(resetCh, skipScene chan bool) {}
