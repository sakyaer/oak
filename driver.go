//+build !js

package oak

import (
	"github.com/oakmound/oak/dlog"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
)

func initDriver(firstScene, imageDir, audioDir string) {
	dlog.Info("Init Scene Loop")
	go sceneLoop(firstScene)
	dlog.Info("Init asset load")
	go loadAssets(imageDir, audioDir)
	dlog.Info("Init Console")
	go defaultDebugConsole()
	dlog.Info("Init Main Driver")
	driver.Main(lifecycleLoop)
}

// A Driver is a function which can take in our lifecycle function
// and initialize oak with the OS interfaces it needs.
type Driver func(f func(screen.Screen))

// InitDriver is the driver oak will call during initialization
var InitDriver = DefaultDriver

// Driver alternatives
var (
	DefaultDriver = driver.Main
	// disabled for https://github.com/golang/go/issues/23451
	// we also need a way to say "you can use this if you have
	// a C compiler, but still compile without using this if you
	// don't"
	// GLDriver = gldriver.Main
)
