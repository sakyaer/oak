//+build js

package oak

import "golang.org/x/exp/shiny/screen"

var (
	drawLoopPublishDef = func(tx screen.Texture) {
		windowControl.Upload(zeroPoint, winBuffer, winBuffer.Bounds())
	}
)
