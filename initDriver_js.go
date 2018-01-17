//+build js

package oak

import "github.com/oakmound/oak/dlog"

func initDriver(firstScene, imageDir, audioDir string) {
	dlog.Info("Init JS Driver")
	InitDriver(lifecycleLoop)
}
