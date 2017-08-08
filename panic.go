//+build !js

package oak

import "runtime/debug"

func SetPanicOnFault() {
	debug.SetPanicOnFault(true)
}
