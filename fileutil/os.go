//+build !js

package fileutil

import "os"

func OpenOS(s string) (File, error) {
	return os.Open(s)
}

func Getwd() (string, error) {
	return os.Getwd()
}
