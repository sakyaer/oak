//+build js

package oak

import (
	"image"
	"image/color"
	"image/draw"
	"sync"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/math/f64"

	"github.com/gopherjs/gopherjs/js"
	"github.com/oakmound/oak/dlog"
)

type JSWindow struct {
	ctx       *js.Object
	jsUint8   *js.Object
	imgData   *js.Object
	events    []interface{}
	eventLock sync.Mutex
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
	jsc.eventLock.Lock()
	jsc.events = append(jsc.events, event)
	jsc.eventLock.Unlock()
}

func (jsc *JSWindow) SendFirst(event interface{}) {
	dlog.Error("SendFirst is not yet supported for JSWindow")
}

func (jsc *JSWindow) NextEvent() interface{} {
	if len(jsc.events) > 0 {
		jsc.eventLock.Lock()
		ev := jsc.events[0]
		jsc.events = jsc.events[1:]
		jsc.eventLock.Unlock()
		return ev
	}
	return nil
}

//////////////
// Uploader

func (jsc *JSWindow) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {
	jsc.jsUint8 = js.Global.Get("Uint8ClampedArray").New(src.RGBA().Pix, sr.Max.X, sr.Max.Y)
	jsc.imgData = js.Global.Get("ImageData").New(jsc.jsUint8, sr.Max.X, sr.Max.Y)
	jsc.ctx.Call("putImageData", jsc.imgData, dp.X, dp.Y)
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
