//+build js

package render

import "bitbucket.org/oakmoundstudio/oak/dlog"

func BatchLoad(baseFolder string) error {
	dlog.Info("BatchLoad not supported on JS")
	return nil
}
