//+build !js

package file

import "os"
import "io"

func Open(s string) (io.ReadCloser, error) {
	return os.Open(s)
}

func Getwd() (string, error) {
	return os.Getwd()
}
