package ninetile

import (
	"image"

	"github.com/oakmound/oak/render"
)

type Option func(*NineTile)

func Sides(imgs ...*image.RGBA) Option {
	return func(n *NineTile) {
		sprs := make([]*render.Sprite, len(imgs))
		for i, img := range imgs {
			sprs[i] = render.NewSprite(0, 0, img)
		}
		n.sides = sprs
		// todo: set rotated sides
	}
}
