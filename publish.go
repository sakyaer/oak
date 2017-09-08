//+build !js

package oak

import "golang.org/x/exp/shiny/screen"

var (
	drawLoopPublishDef = func(tx screen.Texture) {
		tx.Upload(zeroPoint, winBuffer, winBuffer.Bounds())
		windowControl.Scale(windowRect, tx, tx.Bounds(), screen.Src, nil)
		windowControl.Publish()
	}
)
