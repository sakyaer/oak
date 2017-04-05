//+build !js

package file

import "os"

func Open(s string) (File, error) {
	return os.Open(s)
}

func Getwd() (string, error) {
	return os.Getwd()
}
