//+build !js

package oak

import (
	"image/draw"

	"github.com/oakmound/shiny/screen"
)

var (
	drawLoopPublishDef = func(tx screen.Texture) {
		tx.Upload(zeroPoint, winBuffer, winBuffer.Bounds())
		windowControl.Scale(windowRect, tx, tx.Bounds(), draw.Src)
		windowControl.Publish()
	}
)
