//+build js

package oak

import (
	"fmt"
	"image"
	"image/draw"

	"github.com/oakmound/oak/dlog"
	"github.com/oakmound/oak/render"
	"golang.org/x/exp/shiny/screen"
)

var (
	imageBlack = image.Black
)

// DrawLoop
// Unless told to stop, the draw channel will repeatedly
// 1. draw black to a temporary buffer
// 2. run any functions bound to precede drawing.
// 3. draw all elements onto the temporary buffer.
// 4. run any functions bound to follow drawing.
// 5. draw the buffer's data at the viewport's position to the screen.
// 6. publish the screen to display in window.
func DrawLoopNoFPS() {
	fmt.Println("DrawLoop waiting")
	<-drawChannel
	for {
		fmt.Println("Draw Loop")
		dlog.Verb("Draw Loop")
	drawSelect:
		select {
		case <-windowUpdateCH:
			<-windowUpdateCH
		case <-drawChannel:
			dlog.Verb("Got something from draw channel")
			<-drawChannel
			dlog.Verb("Starting loading")
			for {
				draw.Draw(worldBuffer.RGBA(), winBuffer.Bounds(), imageBlack, ViewPos, screen.Src)
				draw.Draw(winBuffer.RGBA(), winBuffer.Bounds(), worldBuffer.RGBA(), ViewPos, screen.Src)

				if loadingR != nil {
					loadingR.Draw(winBuffer.RGBA())
				}
				render.DrawStaticHeap(winBuffer.RGBA())

				windowControl.Upload(zeroPoint, winBuffer, winBuffer.Bounds())
				windowControl.Publish()

				select {
				case <-drawChannel:
					break drawSelect
				case viewPoint := <-viewportChannel:
					dlog.Verb("Got something from viewport channel (waiting on draw)")
					updateScreen(viewPoint[0], viewPoint[1])
				default:
				}
			}
		case viewPoint := <-viewportChannel:
			dlog.Verb("Got something from viewport channel")
			updateScreen(viewPoint[0], viewPoint[1])
		default:
			draw.Draw(worldBuffer.RGBA(), winBuffer.Bounds(), imageBlack, ViewPos, screen.Src)

			render.PreDraw()
			render.DrawHeap(worldBuffer.RGBA(), ViewPos, ScreenWidth, ScreenHeight)
			draw.Draw(winBuffer.RGBA(), winBuffer.Bounds(), worldBuffer.RGBA(), ViewPos, screen.Src)
			render.DrawStaticHeap(winBuffer.RGBA())

			windowControl.Upload(zeroPoint, winBuffer, winBuffer.Bounds())
			windowControl.Publish()
		}
	}
}

func DrawLoopFPS() {
	DrawLoopNoFPS()
}
