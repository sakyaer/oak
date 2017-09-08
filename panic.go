//+build !js

package oak

import "runtime/debug"

func setPanicOnFault() {
	debug.SetPanicOnFault(true)
}
