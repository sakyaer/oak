// Copyright 2015 Hajime Hoshi
// Modifications Copyright 2017 Patrick Stephen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// We use this implementation because I can't really
// think of another way of doing this.
// At the time of this writing (April 2017) this is the only file
// in our code with a license-- this should change?

//+build js

package file

import (
	"bytes"
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"os"

	"github.com/gopherjs/gopherjs/js"
)

type fileInfo struct {
	size int64
}

func (fi fileInfo) Name() string {
	return "js_dummy"
}

func (fi fileInfo) Size() int64 {
	return fi.size
}

func (fi fileInfo) Mode() os.FileMode {
	return 0
}

func (fi fileInfo) ModTime() time.Time {
	return time.Time{}
}

func (fi fileInfo) IsDir() bool {
	return false
}

func (fi fileInfo) Sys() interface{} {
	return nil
}

type file struct {
	*bytes.Reader
}

func (f *file) Close() error {
	return nil
}

func (f *file) Stat() (os.FileInfo, error) {
	fi := fileInfo{f.Size()}
	return fi, nil
}

func Open(path string) (File, error) {
	fmt.Println("Open JS")
	var err error
	var content *js.Object
	ch := make(chan struct{})

	req := js.Global.Get("XMLHttpRequest").New()
	req.Call("open", "GET", path, true)
	req.Set("responseType", "arraybuffer")
	req.Call("addEventListener", "load", func() {
		defer close(ch)
		fmt.Println("UHHH")
		status := req.Get("status").Int()
		if 200 <= status && status < 400 {
			content = req.Get("response")
			return
		}
		err = errors.New(fmt.Sprintf("http error: %d", status))
	})
	req.Call("addEventListener", "error", func() {
		defer close(ch)
		fmt.Println("error")
		err = errors.New(fmt.Sprintf("XMLHttpRequest error: %s", req.Get("statusText").String()))
	})
	req.Call("send")

	fmt.Println("Waiting on channel")
	<-ch
	if err != nil {
		return nil, err
	}

	data := js.Global.Get("Uint8Array").New(content).Interface().([]uint8)
	f := &file{bytes.NewReader(data)}
	fmt.Println("Open JS END")
	return f, nil
}

func Getwd() (string, error) {
	win := js.Global.Get("window")
	loc := win.Get("location").Get("pathname").String()
	fmt.Println("Working directory:", loc)
	return filepath.Dir(loc), nil
}
