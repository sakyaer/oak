// +build js

package oak

import "golang.org/x/exp/shiny/screen"

func InitDriver() {
	go lifecycleLoop(new(JSScreen))
}

func WindowController(s screen.Screen, ScreenWidth, ScreenHeight int) (screen.Window, error) {
	return s.NewWindow(&screen.NewWindowOptions{
		Width:  ScreenWidth,
		Height: ScreenHeight,
		Title:  conf.Title,
	})
}
