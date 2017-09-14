//+build js

package oak

import "github.com/oakmound/oak/dlog"

func initDriver(firstScene, imageDir, audioDir string) {
	dlog.Info("Init asset load")
	go loadAssets(imageDir, audioDir)
	dlog.Info("Init JS Driver")
	lifecycleLoop(new(JSScreen), firstScene)
}
