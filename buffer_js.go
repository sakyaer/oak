//+build js

package oak

import (
	"image"

	"bitbucket.org/oakmoundstudio/oak/dlog"
)

type JSBuffer struct {
	rect image.Rectangle
	rgba *image.RGBA
}

func (jsb *JSBuffer) Release() {
	dlog.Info("JSBuffer pretending to release")
}

func (jsb *JSBuffer) Size() image.Point {
	return jsb.rect.Max
}

func (jsb *JSBuffer) Bounds() image.Rectangle {
	return jsb.rect
}

func (jsb *JSBuffer) RGBA() *image.RGBA {
	return jsb.rgba
}
