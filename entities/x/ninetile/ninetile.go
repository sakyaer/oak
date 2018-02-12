package ninetile

import (
	"image/draw"

	"github.com/oakmound/oak/render"
)

const (
	rot90 = iota
	rot180
	rot270
)

type NineTile struct {
	render.LayeredPoint
	w, h           float64
	tileW, tileH   int
	corners        []*render.Sprite
	rotatedCorners [][]*render.Sprite
	sides          []*render.Sprite
	rotatedSides   [][]*render.Sprite
	centers        []*render.Sprite
	Pattern        [][]int
	Scale          bool
}

func NewNineTile(opts ...Option) (*NineTile, error) {
	n := new(NineTile)
	for _, o := range opts {
		o(n)
	}
	if err := n.validate(); err != nil {
		return nil, err
	}
	return n, nil
}

func (n *NineTile) Draw(buff draw.Image) {
	n.DrawOffset(buff, 0, 0)
}

func (n *NineTile) DrawOffset(buff draw.Image, xOff, yOff float64) {

	nw := len(n.Pattern) - 1
	nh := len(n.Pattern[0]) - 1
	nx := n.X()
	ny := n.Y()

	// draw four corners
	n.corners[n.Pattern[0][0]].DrawOffset(buff, nx+xOff, ny+yOff)
	n.corners[n.Pattern[nw][0]].DrawOffset(buff, n.w*float64(nw)+nx+xOff, ny+yOff)
	n.corners[n.Pattern[0][nh]].DrawOffset(buff, nx+xOff, n.h*float64(nh)+ny+yOff)
	n.corners[n.Pattern[nw][nh]].DrawOffset(buff, n.w*float64(nw)+nx+xOff, n.h*float64(nh)+ny+yOff)
	// draw sides
	for x := 1; x < nw-1; x++ {
		n.rotatedSides[rot90][n.Pattern[x][0]].DrawOffset(buff, float64(x)*n.w+nx+xOff, ny+yOff)
		n.rotatedSides[rot270][n.Pattern[x][nh]].DrawOffset(buff, float64(x)*n.w+nx+xOff, n.h*float64(nh)+ny+yOff)
	}
	for y := 1; y < nh-1; y++ {
		n.sides[n.Pattern[0][y]].DrawOffset(buff, nx+xOff, float64(y)*n.h+ny+yOff)
		n.rotatedSides[rot180][n.Pattern[nw][y]].DrawOffset(buff, n.w*float64(nw)+nx+xOff, float64(y)*n.h+ny+yOff)
	}
	// draw centers
	for x := 1; x < nw-1; x++ {
		for y := 1; y < nh-1; y++ {
			n.centers[n.Pattern[x][y]].DrawOffset(buff, float64(x)*n.w+nx+xOff, float64(y)*n.h+ny+yOff)
		}
	}
}

func (n *NineTile) validate() error {
	// todo
	return nil
}
