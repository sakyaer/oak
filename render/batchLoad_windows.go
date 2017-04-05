//+build windows,!js

package render

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"bitbucket.org/oakmoundstudio/oak/dlog"
)

func BatchLoad(baseFolder string) error {

	// dir2 := filepath.Join(dir, "textures")
	folders, _ := ioutil.ReadDir(baseFolder)

	aliasFile, err := ioutil.ReadFile(filepath.Join(baseFolder, "alias.json"))
	aliases := make(map[string]string)
	if err == nil {
		err = json.Unmarshal(aliasFile, &aliases)
		if err != nil {
			dlog.Error("Alias file unparseable: ", err)
		} else {
			dlog.Verb(aliases)
		}
	}

	for i, folder := range folders {

		dlog.Verb("folder ", i, folder.Name())
		if folder.IsDir() {

			var frameW int
			var frameH int

			if folder.Name() == "raw" {
				frameW = 0
				frameH = 0
			} else if result := regexpTwoNumbers.Find([]byte(folder.Name())); result != nil {
				vals := strings.Split(string(result), "x")
				dlog.Verb("Extracted dimensions: ", vals)
				frameW, _ = strconv.Atoi(vals[0])
				frameH, _ = strconv.Atoi(vals[1])
			} else if result := regexpSingleNumber.Find([]byte(folder.Name())); result != nil {
				val, _ := strconv.Atoi(string(result))
				frameW = val
				frameH = val
			} else {
				if aliased, ok := aliases[folder.Name()]; ok {
					if result := regexpTwoNumbers.Find([]byte(aliased)); result != nil {
						vals := strings.Split(string(result), "x")
						dlog.Verb("Extracted dimensions: ", vals)
						frameW, _ = strconv.Atoi(vals[0])
						frameH, _ = strconv.Atoi(vals[1])
					} else if result := regexpSingleNumber.Find([]byte(aliased)); result != nil {
						val, _ := strconv.Atoi(string(result))
						frameW = val
						frameH = val
					} else {
						return errors.New("Alias value not parseable as a frame width and height pair.")
					}
				} else {
					return errors.New("Alias name not found in alias file.")
				}
			}

			files, _ := ioutil.ReadDir(filepath.Join(baseFolder, folder.Name()))
			for _, file := range files {
				if !file.IsDir() {
					n := file.Name()
					switch n[len(n)-4:] {
					case ".png":
						dlog.Verb("loading file ", n)
						buff := loadPNG(baseFolder, filepath.Join(folder.Name(), n))
						w := buff.Bounds().Max.X
						h := buff.Bounds().Max.Y

						dlog.Verb("buffer: ", w, h, " frame: ", frameW, frameH)

						if frameW == 0 || frameH == 0 {
							continue
						} else if w < frameW || h < frameH {
							dlog.Error("File ", n, " in folder", folder.Name(), " is too small for these folder dimensions")
							return errors.New("File in folder is too small for these folder dimensions")

							// Load this as a sheet if it is greater
							// than the folder size's frame size
						} else if w != frameW || h != frameH {
							dlog.Verb("Loading as sprite sheet")
							_, err = LoadSheet(baseFolder, filepath.Join(folder.Name(), n), frameW, frameH, defaultPad)
							if err != nil {
								dlog.Error(err)
							}
						}
					default:
						dlog.Error("Unsupported file ending for batchLoad: ", n)
					}
				}
			}
		} else {
			dlog.Verb("Not Folder")
		}

	}
	return nil
}
