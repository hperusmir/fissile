// Code generated by go-bindata.
// sources:
// scripts/templates/transformations.yml
// scripts/templates/transformations_code.yml
// DO NOT EDIT!

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _scriptsTemplatesTransformationsYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x50\xef\x4a\xc3\x30\x10\xff\xee\x53\x1c\xb1\x90\x9c\x73\x41\xbf\x16\x7c\x92\xa5\x4a\x69\xd3\x2d\xe0\x92\x90\xde\x70\x50\xf6\xee\x5e\xad\x75\x6b\xb5\x03\x41\xcc\xb7\xdc\xef\xee\xf7\x8f\x52\xe9\xdb\x26\xa4\x7d\x49\x2e\xf8\x36\xbf\x01\xb8\x85\xa8\x44\x2c\x69\xa7\x29\xe8\x2a\xf8\xc6\x6d\x05\x42\x48\x3c\x96\xd3\xb1\x44\x5e\x5f\x03\xdf\x56\xbb\x1c\xe4\xb3\x69\xef\xa2\x51\x1b\x21\x65\xd1\x3d\x9e\x36\x6b\xf3\x66\x74\xb1\x1a\xff\x06\x19\xcf\x24\x9f\x00\x84\x03\xc5\x03\xf1\x8d\x52\xce\xd7\xf6\x08\x1a\x1e\x10\x07\xec\x07\x4b\xfd\x3b\xeb\x2c\x09\xc8\xcf\xcd\xeb\xf4\x57\x24\x66\x32\x4c\x5f\xac\xe4\x05\x76\x41\xdc\x75\x23\xb1\x52\x4c\xb6\xb5\x90\xb9\xfa\x78\x0f\x99\x7d\xb5\x7b\xeb\x09\xf2\x27\xc6\x10\x85\x52\xe7\x19\x7f\x79\xdd\xfa\x1a\xf1\xf4\xe1\x96\xab\x4e\x21\xda\x44\xce\xb6\x7a\xda\xed\xf7\x66\xbf\x36\x8d\x1e\x93\xff\x55\xa1\xb3\xa4\xbf\x4b\x89\xc8\x9a\x0d\xf8\x01\xec\x65\x97\x42\xcf\xb2\xd7\x2e\xd9\x8a\x5e\x16\xf2\x0e\xa6\xfe\x2f\xa1\x98\xf0\x8a\xde\xe4\x7b\x00\x00\x00\xff\xff\x95\xea\x34\x61\x1d\x03\x00\x00")

func scriptsTemplatesTransformationsYmlBytes() ([]byte, error) {
	return bindataRead(
		_scriptsTemplatesTransformationsYml,
		"scripts/templates/transformations.yml",
	)
}

func scriptsTemplatesTransformationsYml() (*asset, error) {
	bytes, err := scriptsTemplatesTransformationsYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/templates/transformations.yml", size: 797, mode: os.FileMode(436), modTime: time.Unix(1444058541, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _scriptsTemplatesTransformations_codeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x29\x4a\xcc\x2b\x4e\xcb\x2f\xca\x4d\x2c\xc9\xcc\xcf\x2b\xb6\xe2\x52\x50\x50\x56\x48\xcc\x53\x48\x2d\x4a\x52\x48\xce\x4f\x49\x55\x48\xcd\x4b\x51\x48\xca\xc9\x4f\xce\xb6\x52\xb0\x51\x05\xf3\x54\xed\x80\x8a\x74\x15\x80\x3a\x92\x33\xac\x14\xd4\xe3\x62\x8a\xb5\x80\xc2\x40\x52\x45\x1d\x28\xa1\xa0\x90\x5f\x5a\x52\x50\x5a\x02\x94\xa9\xae\x06\xab\xaf\xad\x55\xc7\x34\x35\xa7\x38\x15\xd9\x58\x10\x17\x9b\xb9\x40\x71\x1c\x06\x83\x74\x80\x4c\x06\x04\x00\x00\xff\xff\x2b\x15\x9c\x5c\xc1\x00\x00\x00")

func scriptsTemplatesTransformations_codeYmlBytes() ([]byte, error) {
	return bindataRead(
		_scriptsTemplatesTransformations_codeYml,
		"scripts/templates/transformations_code.yml",
	)
}

func scriptsTemplatesTransformations_codeYml() (*asset, error) {
	bytes, err := scriptsTemplatesTransformations_codeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/templates/transformations_code.yml", size: 193, mode: os.FileMode(436), modTime: time.Unix(1444058541, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"scripts/templates/transformations.yml": scriptsTemplatesTransformationsYml,
	"scripts/templates/transformations_code.yml": scriptsTemplatesTransformations_codeYml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"scripts": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"transformations.yml": &bintree{scriptsTemplatesTransformationsYml, map[string]*bintree{
			}},
			"transformations_code.yml": &bintree{scriptsTemplatesTransformations_codeYml, map[string]*bintree{
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        // File
        if err != nil {
                return RestoreAsset(dir, name)
        }
        // Dir
        for _, child := range children {
                err = RestoreAssets(dir, filepath.Join(name, child))
                if err != nil {
                        return err
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
