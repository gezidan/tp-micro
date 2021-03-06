// Code generated by go-bindata.
// sources:
// __tp-micro__tpl__.go
// DO NOT EDIT!

package test

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	name    string
	size    int64
	mode    os.FileMode
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

var ___tpMicro__tpl__Go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x52\x51\x6f\xda\x30\x10\x7e\xb6\x7f\xc5\xa9\x4f\x94\x87\xa5\xdd\xaa\x3e\xa0\x69\x12\x2b\xd3\xa8\x04\x85\x0d\x78\x98\xa6\x29\x98\xe4\x08\x5e\x13\x3b\xb5\x2f\x55\x51\xc5\x7f\x9f\x6c\x27\x24\x61\x3c\x04\xdb\xdf\xdd\x77\xdf\xdd\x77\x51\x04\xa5\x48\x9e\x45\x86\x10\xc7\xeb\xe5\x2c\x8e\x41\x5a\xa0\x03\x42\x69\xf4\x5f\x4c\x08\x08\x8b\x32\x17\x84\xfc\x22\x8e\xf3\x28\x82\x38\x1e\x2f\x1f\xe3\x78\xb9\x99\xb9\x4c\x83\x99\xb4\x84\x06\xdc\x1d\x8c\xae\x08\xcd\xc8\x85\x41\x74\xd0\x05\x86\x53\x21\xe8\x10\xa5\xf2\x55\xa6\xc8\xe9\x58\xe2\x05\x87\x54\x84\x66\x2f\x12\x84\x77\xce\xa6\xba\xc0\xc1\xd0\x92\xa9\x12\x7a\x3f\x5d\xc3\xd0\x3d\xfc\x44\x5b\xe5\xc4\xd9\x5c\xd0\x81\x9f\xfa\x32\x56\xd3\xbe\x8c\xd5\xb4\x2f\xc3\x92\xa0\x8b\xaa\x3e\xa5\x57\x75\x45\x82\x06\x43\xf7\x1d\x9b\xec\xfa\x5c\x62\xfe\x6b\xf5\x63\x16\xc7\xf3\xc5\xe4\x9b\x13\x9a\x18\x14\x84\x50\x1c\xed\x4b\x0e\x85\x4e\x31\x6f\x88\x2f\x02\x83\x7c\x47\xbc\xb1\x68\x38\x9b\xe9\x8c\xb3\x09\xbe\xca\x04\x5b\xee\xc5\xd3\xf7\xc5\xff\xdc\x5a\x65\x3a\xdd\x5d\xb0\xf7\x43\x5b\xf6\x39\x92\xa8\x09\xdd\x68\x20\xd1\x8a\x8c\xce\x73\x34\x21\xd5\x3f\xf6\x1a\x8d\x22\x98\x78\x23\xe0\x20\x54\xea\x02\x59\xb8\x0f\x86\xe1\xdf\xf5\x0f\xf5\xb9\x1e\x7b\x28\xd0\xfa\x00\xce\x59\x30\x01\xf4\x75\x3a\x58\x2b\xee\x41\x2b\x42\xe5\x5f\xa4\xca\x20\x8a\x80\xf0\xcd\xb3\xf9\x9c\x41\x47\xcc\xd8\x64\x10\xf6\x03\x44\x29\x41\x98\xac\x91\xe5\x90\x96\xd2\x65\x84\x38\x95\x72\xc6\xc6\xb0\xcf\xb5\xa0\xfb\xbb\x16\xb1\xda\x70\xc6\xbe\x36\x00\x6c\x4b\x61\x44\x31\xba\xfa\x6c\x84\xca\x70\x04\x37\x1f\x6e\x6e\x47\xb7\x37\xee\xf7\xe5\x6a\xcb\xd9\xa9\xa3\xa2\xee\xa0\x23\xa4\xee\x91\xf5\xf0\xbe\x9c\x97\x4a\x93\x44\x45\x9c\xb1\x87\x56\xce\x89\x5f\xfb\xa1\xd5\x1b\x05\x6e\x09\x9b\x89\xfb\xf6\xfc\x08\x5a\xb4\x61\x5c\x5b\x67\xd7\xfd\x9d\x9f\x96\x2c\xd0\x92\x28\x4a\x5b\x3b\xe0\x76\x09\x2a\xf7\x91\x6a\xaf\x03\x85\x7f\x6b\xf3\x1f\x53\x80\x9a\x01\xb6\xcf\x78\x1c\x5d\x95\x46\xba\x3e\x9f\x44\x81\x8d\x13\x01\xa8\x94\x07\xc6\x19\xfa\x8c\x4f\x1f\xcf\xce\xcc\x74\x4f\x12\xbe\x35\x1e\x9e\x23\xc2\x2a\x77\xd7\x7c\xf3\x38\xe9\xd3\x87\xba\x4d\x82\x5b\xd4\x4e\xf8\x54\xef\x76\x47\xf8\xfd\xa7\xa6\x65\x6b\x91\x59\x68\xef\x27\xfe\x2f\x00\x00\xff\xff\x9f\x42\x19\x52\xa4\x04\x00\x00")

func __tpMicro__tpl__GoBytes() ([]byte, error) {
	return bindataRead(
		___tpMicro__tpl__Go,
		"__tp-micro__tpl__.go",
	)
}

func __tpMicro__tpl__Go() (*asset, error) {
	bytes, err := __tpMicro__tpl__GoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "__tp-micro__tpl__.go", size: 1188, mode: os.FileMode(420), modTime: time.Unix(1530118210, 0)}
	a := &asset{bytes: bytes, info: info}
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
	if err != nil {
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
	"__tp-micro__tpl__.go": __tpMicro__tpl__Go,
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
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"__tp-micro__tpl__.go": &bintree{__tpMicro__tpl__Go, map[string]*bintree{}},
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
