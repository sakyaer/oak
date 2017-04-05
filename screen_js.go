//+build js

package oak

import (
	"errors"
	"image"

	"github.com/gopherjs/gopherjs/js"
	"golang.org/x/exp/shiny/screen"
)

type JSScreen struct{}

func (jss *JSScreen) NewBuffer(p image.Point) (screen.Buffer, error) {
	rect := image.Rect(0, 0, p.X, p.Y)
	rgba := image.NewRGBA(rect)
	buffer := &JSBuffer{
		rect,
		rgba,
	}
	return buffer, nil
}
func (jss *JSScreen) NewWindow(opts *screen.NewWindowOptions) (screen.Window, error) {
	jsc := new(JSWindow)

	document := js.Global.Get("document")
	canvas := document.Call("createElement", "canvas")
	jsc.ctx = canvas.Call("getContext", "2d")
	document.Get("body").Call("appendChild", canvas)

	return jsc, nil
}

func (jss *JSScreen) NewTexture(p image.Point) (screen.Texture, error) {
	return nil, errors.New("Not supported on JS")
}
