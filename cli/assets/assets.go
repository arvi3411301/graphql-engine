// Code generated by go-bindata.
// sources:
// assets/unversioned/console.html
// assets/v1.0-alpha/console.html
// assets/v1.0/console.html
// DO NOT EDIT!

package assets

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

var _assetsUnversionedConsoleHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x51\x8b\xe3\x36\x10\x7e\xbf\x5f\x31\xa8\x94\x7b\xa9\xa5\xec\xf5\xa0\x87\x93\x2c\x1c\x2d\x85\xd2\x2b\x1c\xed\xa5\xaf\xc7\xac\x3c\xb6\x67\x4f\x96\x5c\x8d\x92\xdd\x34\xe4\xbf\x17\xdb\xb1\xe3\x64\x97\x96\x2e\x2d\x79\x91\xbe\xf9\x26\x9f\xe6\xd3\x78\xb4\xaa\x53\xe3\xc0\xa1\xaf\xd6\x8a\x7c\xb6\x15\x75\xfb\x0a\x60\x55\x13\x16\xdd\x02\x60\xe5\xd8\x7f\x81\x48\x6e\xad\xd8\x06\xaf\x20\xed\x5b\x5a\x2b\x6e\xb0\x22\xd3\xfa\x4a\x41\x1d\xa9\x5c\xab\x3a\xa5\x56\x72\x63\x24\x85\x88\x15\xe9\x2a\x84\xca\x11\xb6\x2c\xda\x86\xc6\xd4\x28\xdb\x88\x59\x15\xb1\xad\xff\x70\x19\xf9\x8a\x3d\x19\x1b\xbc\x04\x47\x06\x45\x28\x89\x29\x71\xd7\x69\xe8\xfe\x6f\xcd\x49\x5f\x6c\xe4\x36\x0d\x1b\x80\x07\xf6\x45\x78\xd0\x9f\x3f\x93\xdf\xc1\x1a\x0e\x27\x18\x00\x5b\xde\xfc\xfa\x21\x87\xc3\x41\x0f\xcb\xe3\xf1\x9b\x29\x68\x1d\xff\x4e\x51\x38\xf8\x9e\x70\xde\xce\x49\x05\x26\x7c\xdf\xf2\x26\xba\x9e\x74\xde\x3e\x43\x9a\xff\xdb\x25\x34\x27\x1f\x0e\x19\x70\x09\x3e\x24\xd0\x6d\x0c\x8f\xfb\x8f\x98\x6a\x38\x1e\xaf\x09\xba\x46\x79\x6f\x2d\x89\xfc\x4c\xfb\x79\x1c\x47\x70\xa8\xab\x68\xd8\xff\x46\x36\x52\xba\x54\x01\x72\x42\x17\x79\x67\xe6\xdf\x67\x66\x40\xbe\xb8\x3e\xd1\x15\xb4\x8d\xee\x63\xa4\x92\x1f\x73\x50\x46\xcd\x4c\x1d\x2e\xef\x97\x50\x50\x0e\xca\x3a\x56\x17\x86\x6f\x36\x3f\xfd\x30\xba\xdd\xad\x9f\x31\xe6\x6c\xca\x4c\x6e\xc2\xfa\xe4\x19\xe3\xb2\x62\x5f\x40\x36\xcb\x22\x8f\x77\x8e\x3e\x91\xa3\x86\x52\x1c\xfc\xba\xc2\x26\xf6\x71\x39\x74\x96\x39\xb7\xd6\xca\x8c\x1d\xbf\xba\x0b\xc5\x7e\x6c\xbd\xb4\x77\x34\xac\x75\x83\xec\xbf\x0f\x3e\x91\x4f\x53\xd7\x15\x2c\xad\xc3\x7d\x0e\xaf\x7d\xf0\xf4\x7a\x79\x82\x43\x8b\x96\xd3\x3e\x87\xc5\x88\xa4\x88\x5e\x38\xf5\x1d\x73\x8a\x82\x7e\xb3\x10\x70\xec\x09\xe3\x40\x3b\x3e\x11\xd2\x52\x87\x87\x67\xd4\xee\x5c\xb0\x5f\x9e\xca\xdd\xbc\x40\x6e\x65\x4e\x45\x0e\xbb\x82\x77\xc0\xc5\x5a\xb9\x80\x05\xfb\x4a\x8d\x9f\x5d\x1f\xb0\x0e\x45\xd6\xaa\xc5\x8a\xb2\x91\x00\x7d\xfa\x5a\x9d\x78\x0d\xfb\xac\x26\xae\xea\x94\xc3\xcd\x62\xb1\xab\x97\xd3\x77\x5b\x74\x37\x7a\xb3\x58\x7c\xbd\xbc\xae\xa7\x74\xf4\x38\x82\xe8\xb8\xf2\x19\x27\x6a\x24\x07\x4b\x3e\x51\x1c\x43\x65\xf0\x29\x2b\xb1\x61\xb7\xcf\x41\xd0\x4b\x26\x14\xb9\x1c\xc3\xf7\x5b\x49\x5c\xee\x33\x3b\x78\x77\x9d\x3d\x95\xd2\xdd\x6b\x8b\x7e\xac\xe6\xba\x82\x93\x8e\xf0\x9f\x94\xc3\x1b\x6a\x96\x13\xde\x60\xac\xd8\x67\x29\xb4\x39\x64\xdf\xce\x23\x36\xb8\x10\x73\xf8\xea\xdd\xdb\xee\x77\xc6\x67\x9a\x1f\x06\xbf\xb4\xd6\xa3\xa3\xa6\x3b\xc5\xe4\xaf\x29\x78\x77\xea\xba\xd9\x72\xbc\x8e\x53\x4d\x6a\x3c\xf4\xac\x47\xd4\xed\x3c\xe1\x3c\xac\xfb\xaa\xa4\x26\x4a\xd7\x13\xda\x16\xfe\x5e\xb4\x75\x61\x5b\x94\x0e\x23\xf5\xf3\x19\xef\xf1\xd1\x38\xbe\x13\xd3\x97\x8f\x0f\x24\xa1\x21\xf3\x56\x7f\xa7\x17\xc6\xca\x25\xac\x1b\xf6\xda\x8a\x28\xf3\x2f\x64\x5f\xf4\x30\x74\xb3\xab\x7f\x1b\xa6\xe9\x6a\xba\xda\x7b\x71\xb0\x35\x46\xa1\xb4\x56\x9b\x4f\x3f\x66\xef\xd4\xe5\x83\x01\x12\xed\x7f\x2f\xbe\x23\x5f\x84\xa8\xef\x9f\xaa\xdf\xce\xc7\xc9\xff\x7c\x8a\xde\x82\x7f\x3a\xc3\xca\x0c\x93\x6c\x65\xba\x27\xfe\xf6\xd5\x5f\x01\x00\x00\xff\xff\xc8\x30\xa7\xed\xea\x07\x00\x00")

func assetsUnversionedConsoleHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsUnversionedConsoleHtml,
		"assets/unversioned/console.html",
	)
}

func assetsUnversionedConsoleHtml() (*asset, error) {
	bytes, err := assetsUnversionedConsoleHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/unversioned/console.html", size: 2026, mode: os.FileMode(420), modTime: time.Unix(1552991624, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsV10AlphaConsoleHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x51\x8b\xe3\x36\x10\x7e\xbf\x5f\x31\xa8\x94\x7b\xa9\xa5\xec\xf5\xa0\x87\x93\x2c\x1c\x2d\x85\xd2\x2b\x1c\xed\xa5\xaf\xc7\xac\x3c\xb6\x67\x4f\x96\x5c\x8d\x92\xdd\x34\xe4\xbf\x17\xdb\xb1\xe3\x64\x97\x96\x2e\x2d\x79\x91\xbe\xf9\x26\x9f\xe6\xd3\x78\xb4\xaa\x53\xe3\xc0\xa1\xaf\xd6\x8a\x7c\xb6\x15\x75\xfb\x0a\x60\x55\x13\x16\xdd\x02\x60\xe5\xd8\x7f\x81\x48\x6e\xad\xd8\x06\xaf\x20\xed\x5b\x5a\x2b\x6e\xb0\x22\xd3\xfa\x4a\x41\x1d\xa9\x5c\xab\x3a\xa5\x56\x72\x63\x24\x85\x88\x15\xe9\x2a\x84\xca\x11\xb6\x2c\xda\x86\xc6\xd4\x28\xdb\x88\x59\x15\xb1\xad\xff\x70\x19\xf9\x8a\x3d\x19\x1b\xbc\x04\x47\x06\x45\x28\x89\x29\x71\xd7\x69\xe8\xfe\x6f\xcd\x49\x5f\x6c\xe4\x36\x0d\x1b\x80\x07\xf6\x45\x78\xd0\x9f\x3f\x93\xdf\xc1\x1a\x0e\x27\x18\x00\x5b\xde\xfc\xfa\x21\x87\xc3\x41\x0f\xcb\xe3\xf1\x9b\x29\x68\x1d\xff\x4e\x51\x38\xf8\x9e\x70\xde\xce\x49\x05\x26\x7c\xdf\xf2\x26\xba\x9e\x74\xde\x3e\x43\x9a\xff\xdb\x25\x34\x27\x1f\x0e\x19\x70\x09\x3e\x24\xd0\x6d\x0c\x8f\xfb\x8f\x98\x6a\x38\x1e\xaf\x09\xba\x46\x79\x6f\x2d\x89\xfc\x4c\xfb\x79\x1c\x47\x70\xa8\xab\x68\xd8\xff\x46\x36\x52\xba\x54\x01\x72\x42\x17\x79\x67\xe6\xdf\x67\x66\x40\xbe\xb8\x3e\xd1\x15\xb4\x8d\xee\x63\xa4\x92\x1f\x73\x50\x46\xcd\x4c\x1d\x2e\xef\x97\x50\x50\x0e\xca\x3a\x56\x17\x86\x6f\x36\x3f\xfd\x30\xba\xdd\xad\x9f\x31\xe6\x6c\xca\x4c\x6e\xc2\xfa\xe4\x19\xe3\xb2\x62\x5f\x40\x36\xcb\x22\x8f\x77\x8e\x3e\x91\xa3\x86\x52\x1c\xfc\xba\xc2\x26\xf6\x71\x39\x74\x96\x39\xb7\xd6\xca\x8c\x1d\xbf\xba\x0b\xc5\x7e\x6c\xbd\xb4\x77\x34\xac\x75\x83\xec\xbf\x0f\x3e\x91\x4f\x53\xd7\x15\x2c\xad\xc3\x7d\x0e\xaf\x7d\xf0\xf4\x7a\x79\x82\x43\x8b\x96\xd3\x3e\x87\xc5\x88\xa4\x88\x5e\x38\xf5\x1d\x73\x8a\x82\x7e\xb3\x10\x70\xec\x09\xe3\x40\x3b\x3e\x11\xd2\x52\x87\x87\x67\xd4\xee\x5c\xb0\x5f\x9e\xca\xdd\xbc\x40\x6e\x65\x4e\x45\x0e\xbb\x82\x77\xc0\xc5\x5a\xb9\x80\x05\xfb\x4a\x8d\x9f\x5d\x1f\xb0\x0e\x45\xd6\xaa\xc5\x8a\xb2\x91\x00\x7d\xfa\x5a\x9d\x78\x0d\xfb\xac\x26\xae\xea\x94\xc3\xcd\x62\xb1\xab\x97\xd3\x77\x5b\x74\x37\x7a\xb3\x58\x7c\xbd\xbc\xae\xa7\x74\xf4\x38\x82\xe8\xb8\xf2\x19\x27\x6a\x24\x07\x4b\x3e\x51\x1c\x43\x65\xf0\x29\x2b\xb1\x61\xb7\xcf\x41\xd0\x4b\x26\x14\xb9\x1c\xc3\xf7\x5b\x49\x5c\xee\x33\x3b\x78\x77\x9d\x3d\x95\xd2\xdd\x6b\x8b\x7e\xac\xe6\xba\x82\x93\x8e\xf0\x9f\x94\xc3\x1b\x6a\x96\x13\xde\x60\xac\xd8\x67\x29\xb4\x39\x64\xdf\xce\x23\x36\xb8\x10\x73\xf8\xea\xdd\xdb\xee\x77\xc6\x67\x9a\x1f\x06\xbf\xb4\xd6\xa3\xa3\xa6\x3b\xc5\xe4\xaf\x29\x78\x77\xea\xba\xd9\x72\xbc\x8e\x53\x4d\x6a\x3c\xf4\xac\x47\xd4\xed\x3c\xe1\x3c\xac\xfb\xaa\xa4\x26\x4a\xd7\x13\xda\x16\xfe\x5e\xb4\x75\x61\x5b\x94\x0e\x23\xf5\xf3\x19\xef\xf1\xd1\x38\xbe\x13\xd3\x97\x8f\x0f\x24\xa1\x21\xf3\x56\x7f\xa7\x17\xc6\xca\x25\xac\x1b\xf6\xda\x8a\x28\xf3\x2f\x64\x5f\xf4\x30\x74\xb3\xab\x7f\x1b\xa6\xe9\x6a\xba\xda\x7b\x71\xb0\x35\x46\xa1\xb4\x56\x9b\x4f\x3f\x66\xef\xd4\xe5\x83\x01\x12\xed\x7f\x2f\xbe\x23\x5f\x84\xa8\xef\x9f\xaa\xdf\xce\xc7\xc9\xff\x7c\x8a\xde\x82\x7f\x3a\xc3\xca\x0c\x93\x6c\x65\xba\x27\xfe\xf6\xd5\x5f\x01\x00\x00\xff\xff\xc8\x30\xa7\xed\xea\x07\x00\x00")

func assetsV10AlphaConsoleHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsV10AlphaConsoleHtml,
		"assets/v1.0-alpha/console.html",
	)
}

func assetsV10AlphaConsoleHtml() (*asset, error) {
	bytes, err := assetsV10AlphaConsoleHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/v1.0-alpha/console.html", size: 2026, mode: os.FileMode(420), modTime: time.Unix(1552991619, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsV10ConsoleHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\xdf\x8f\xdb\x36\x0c\x7e\xef\x5f\x41\x68\x18\xfa\x32\x4b\xb9\xae\xc0\x0a\x27\x3e\xa0\x58\x31\x60\x58\x07\x14\x6b\xb3\xd7\x42\x27\xd3\x36\xaf\xb2\xe4\x89\x4a\xee\xbc\x20\xff\xfb\xe0\x9f\x71\x72\xb7\xde\x76\xd8\x90\x17\x89\xfc\x68\xea\xfb\x48\x51\xd9\x54\xb1\xb6\x60\xb5\x2b\x33\x81\x2e\xd9\xb1\xb8\x7e\x01\xb0\xa9\x50\xe7\xdd\x02\x60\x63\xc9\x7d\x81\x80\x36\x13\x64\xbc\x13\x10\xdb\x06\x33\x41\xb5\x2e\x51\x35\xae\x14\x50\x05\x2c\x32\x51\xc5\xd8\x70\xaa\x14\x47\x1f\x74\x89\xb2\xf4\xbe\xb4\xa8\x1b\x62\x69\x7c\xad\x2a\xcd\xbb\xa0\x93\x32\xe8\xa6\xfa\xc3\x26\xe8\x4a\x72\xa8\x8c\x77\xec\x2d\x2a\xcd\x8c\x91\x55\xa1\xf7\x5d\x0e\xd9\x7f\x56\x8d\xf9\xd9\x04\x6a\xe2\xb0\x01\xb8\x23\x97\xfb\x3b\xf9\xf9\x33\xba\x3d\x64\x70\x18\xcd\x00\xba\xa1\xed\x6f\xef\x53\x38\x1c\xe4\xb0\x3c\x1e\xbf\x9b\x9d\xc6\xd2\xef\x18\x98\xbc\xeb\x01\xa7\xed\x12\x94\xeb\xa8\xdf\x36\xb4\x0d\xb6\x07\x9d\xb6\x8f\x80\x96\x5f\x3b\x37\x2d\xc1\x87\x43\x02\x54\x80\xf3\x11\x64\x13\xfc\x7d\xfb\x41\xc7\x0a\x8e\xc7\x4b\x80\xac\x34\xbf\x35\x06\x99\x7f\xc1\x76\xe9\xd7\x93\x71\xe0\x95\xd7\xe4\x3e\xa2\x09\x18\xcf\xb3\x00\x5a\xc6\xb3\xb8\x13\xf2\xeb\x91\x09\xa0\xcb\x2f\x4f\x74\x61\xda\x05\xfb\x21\x60\x41\xf7\x29\x08\x25\x16\xa2\x0e\xc5\xfb\xd5\xe7\x98\x82\x30\x96\xc4\x99\xe0\xdb\xed\xcf\xef\x26\xb5\xbb\xf5\x23\xc2\x9c\x44\x59\xa4\x9b\x6d\x7d\xf0\x02\x71\xce\xd8\xe5\x90\x2c\xa2\xd0\xe9\x1b\x8b\x9f\xd0\x62\x8d\x31\x0c\x7a\x5d\xd8\x66\xf4\x71\x3d\x74\x96\x3a\xb5\xd6\x46\x4d\x1d\xbf\xb9\xf1\x79\x3b\xb5\x5e\x6c\x2d\x0e\x6b\x59\x6b\x72\x3f\x7a\x17\xd1\xc5\xb9\xeb\x72\xe2\xc6\xea\x36\x85\x97\xce\x3b\x7c\xb9\x1e\xcd\xbe\xd1\x86\x62\x9b\xc2\x6a\xb2\xc4\xa0\x1d\x53\xec\x3b\x66\xf4\x82\x7c\xb5\x62\xb0\xe4\x50\x87\x01\x76\x7c\x90\x48\x72\xe5\xef\x1e\xc9\x76\x63\xbd\xf9\xf2\x30\xdd\xd5\x33\xd2\x6d\xd4\x48\x72\xd8\xe5\xb4\x07\xca\x33\x61\xbd\xce\xc9\x95\x62\xba\x76\xbd\xc3\x58\xcd\x9c\x89\x46\x97\x98\x4c\x00\xe8\xc3\x33\x31\xe2\x6a\x72\x49\x85\x54\x56\x31\x85\xab\xd5\x6a\x5f\xad\xe7\x7b\x9b\x77\x15\xbd\x5a\xad\xbe\x5d\x5f\xf2\x29\x2c\xde\x4f\x46\x6d\xa9\x74\x09\x45\xac\x39\x05\x83\x2e\x62\x98\x5c\x85\x77\x31\x29\x74\x4d\xb6\x4d\x81\xb5\xe3\x84\x31\x50\x31\xb9\x6f\x77\x1c\xa9\x68\x13\x33\x68\x77\x19\x3d\x53\xe9\xea\xda\x68\x37\xb1\xb9\x64\x30\xe6\x61\xfa\x13\x53\x78\x85\xf5\x7a\xb6\xd7\x3a\x94\xe4\x92\xe8\x9b\x14\x92\xef\x97\x1e\xe3\xad\x0f\x29\x7c\xf3\xe6\x75\xf7\x3b\xd9\x17\x39\xdf\x0f\x7a\x49\x29\x27\x45\x55\x77\x8a\x59\x5f\x95\xd3\x7e\xec\xba\xc5\x72\x2a\xc7\xc8\x49\x4c\x87\x5e\xf4\x88\xb8\x5e\x06\x9c\x86\x75\xcf\x8a\x2b\xc4\x78\x39\xa1\x4d\xee\x6e\x59\x1a\xeb\x77\x79\x61\x75\xc0\x7e\x3e\xeb\x5b\x7d\xaf\x2c\xdd\xb0\xea\xe9\xeb\x3b\x64\x5f\xa3\x7a\x2d\x7f\x90\x2b\x65\xf8\xdc\x2c\x6b\x72\xd2\x30\x0b\x35\xf6\xcd\xe1\xd0\x5f\x67\x63\xe9\x63\xd4\x91\xcc\x3b\x0a\xdd\x04\x79\xfa\x4c\x8a\x7b\xbc\xea\x08\xf5\x5f\x04\x53\xe9\xc0\x18\x33\xb1\xfd\xf4\x53\xf2\x46\x9c\xbf\x02\xc0\xc1\x9c\x82\xf6\xe8\x72\x1f\xe4\xed\xc3\xa8\xeb\xe5\xdd\xfe\x9b\xe8\x3e\xe5\xd7\x63\x27\x6e\xd3\x78\xfd\xe7\x1a\x3f\xeb\x15\xec\x06\x75\xff\x10\xce\x4f\xc9\xbf\xd3\xe5\xbf\x4e\xfe\x3c\x7d\xff\x17\x09\x9e\xae\xd3\xf8\x26\x0c\x55\xda\xa8\x61\x8a\x6f\x54\xf7\xf7\xe6\xfa\xc5\x5f\x01\x00\x00\xff\xff\xfe\x73\xb6\x1b\xe6\x08\x00\x00")

func assetsV10ConsoleHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsV10ConsoleHtml,
		"assets/v1.0/console.html",
	)
}

func assetsV10ConsoleHtml() (*asset, error) {
	bytes, err := assetsV10ConsoleHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/v1.0/console.html", size: 2278, mode: os.FileMode(420), modTime: time.Unix(1552991622, 0)}
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
	"assets/unversioned/console.html": assetsUnversionedConsoleHtml,
	"assets/v1.0-alpha/console.html": assetsV10AlphaConsoleHtml,
	"assets/v1.0/console.html": assetsV10ConsoleHtml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"unversioned": &bintree{nil, map[string]*bintree{
			"console.html": &bintree{assetsUnversionedConsoleHtml, map[string]*bintree{}},
		}},
		"v1.0": &bintree{nil, map[string]*bintree{
			"console.html": &bintree{assetsV10ConsoleHtml, map[string]*bintree{}},
		}},
		"v1.0-alpha": &bintree{nil, map[string]*bintree{
			"console.html": &bintree{assetsV10AlphaConsoleHtml, map[string]*bintree{}},
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

