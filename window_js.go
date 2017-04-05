//+build js

package oak

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/math/f64"

	"bitbucket.org/oakmoundstudio/oak/dlog"
	"github.com/gopherjs/gopherjs/js"
)

type JSWindow struct {
	ctx *js.Object
}

func (jsc *JSWindow) Release() {
	dlog.Info("JSWindow releasing")
}

func (jsc *JSWindow) Publish() screen.PublishResult {
	// Publish doesn't do anything on JS
	// (Publish doesn't do anything on windows either)
	return screen.PublishResult{}
}

/////////////
// EventDeque
func (jsc *JSWindow) Send(event interface{}) {
	dlog.Error("Send is not yet supported for JSWindow")
}

func (jsc *JSWindow) SendFirst(event interface{}) {
	dlog.Error("SendFirst is not yet supported for JSWindow")
}

func (jsc *JSWindow) NextEvent() interface{} {
	//Todo
	return nil
}

//////////////
// Uploader

func (jsc *JSWindow) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {
	rgba := src.RGBA()
	fmt.Println(len(rgba.Pix))
	jsUint8 := js.Global.Get("Uint8ClampedArray").New(rgba.Pix)
	fmt.Println(sr)
	jsc.ctx.Call("putImageData", js.Global.Get("ImageData").New(jsUint8, sr.Max.X, sr.Max.Y), dp.X, dp.Y)
}

func (jsc *JSWindow) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	//Todo
}

///////////////
// Drawer

func (jsc *JSWindow) Draw(src2dst f64.Aff3, src screen.Texture, sr image.Rectangle, op draw.Op, opts *screen.DrawOptions) {
	//Todo
}

// DrawUniform is like Draw except that the src is a uniform color instead
// of a Texture.
func (jsc *JSWindow) DrawUniform(src2dst f64.Aff3, src color.Color, sr image.Rectangle, op draw.Op, opts *screen.DrawOptions) {
	//Todo
}

// Copy copies the sub-Texture defined by src and sr to the destination
// (the method receiver), such that sr.Min in src-space aligns with dp in
// dst-space.
func (jsc *JSWindow) Copy(dp image.Point, src screen.Texture, sr image.Rectangle, op draw.Op, opts *screen.DrawOptions) {
	//Todo
}

// Scale scales the sub-Texture defined by src and sr to the destination
// (the method receiver), such that sr in src-space is mapped to dr in
// dst-space.
func (jsc *JSWindow) Scale(dr image.Rectangle, src screen.Texture, sr image.Rectangle, op draw.Op, opts *screen.DrawOptions) {
	//Todo
}
