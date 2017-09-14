//+build !js

package oak

import (
	"github.com/oakmound/oak/dlog"
	"golang.org/x/exp/shiny/driver"
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
