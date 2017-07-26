// Code generated by go-bindata.
// sources:
// data/locations.json
// data/profiles.json
// data/rates.json
// DO NOT EDIT!

package data

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

var _dataLocationsJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x95\x3d\x8a\x23\x41\x0c\x85\xe3\xee\x53\x98\x8e\xbd\x46\xbf\x25\xc9\x37\xd8\x33\x2c\x1b\x2c\xec\xc0\x04\x66\x9c\x38\x1b\xe6\xee\x83\x99\xcc\x4f\x4a\x1b\xbe\x2a\x95\xf4\xe9\xf5\x9f\x7d\xfb\xdc\xb7\xed\x78\xbf\x3f\xde\x6e\xbf\xff\x1f\xd7\xd3\xc1\xc7\xf9\xf9\xe5\xf6\xef\x71\x5c\x4f\x1a\x97\xc8\x15\x3f\x5f\xee\x1f\xc7\xf5\xf4\x8b\x45\x2e\xc6\x2c\xfb\xf6\x75\x46\x5a\x80\x76\x03\x9a\xc8\x7b\x5a\x81\xd6\x86\x0e\xee\x69\x7b\xa5\x4b\xd7\x2b\xad\xa5\xd4\xd3\x8e\x77\x33\xbe\x3b\x87\xbb\x17\x76\x4d\xb1\x72\x1e\xde\x1d\xd8\x35\xbc\x9b\x7c\xb8\x3b\x81\xb6\x42\x3a\xa3\xa7\x0b\x2b\x6f\xe8\x1a\x68\x26\x68\xba\x37\x0f\xa7\x09\x07\xdb\x8a\x1b\xdc\x27\xbc\xd1\x2d\x71\x68\x94\x03\x8e\xbe\x05\x35\xae\xeb\x80\xbf\x0a\x97\xb4\xb4\x19\x3b\x9b\x58\xec\xc3\x19\xa8\x5d\x34\xeb\x96\x83\xb4\x0c\xde\x45\x36\x2f\xf0\x09\x47\xf1\x9a\xe1\xf3\xb4\x32\x8c\xe6\x2d\xd8\x38\xa3\x9c\xfa\x87\xea\x55\xb3\xee\x35\x45\x0d\xa8\x97\x21\x86\xed\x17\x77\x5a\xc3\x11\x18\x76\x59\x29\xdd\x13\xc6\x2a\x1a\x05\xdd\x70\x84\x92\x16\x53\x15\xa8\x61\x75\xe1\x43\xa1\x6c\xc3\x11\x90\x7d\x19\x52\xcd\x28\x49\xa3\x86\x23\xd0\xc4\x95\xd1\x44\xa0\xae\x29\x80\xa5\x49\xc1\x30\x6c\xa7\x92\x4c\x59\x26\x60\x64\x91\x52\x13\x48\x6e\x34\x58\x29\x68\x65\x45\x34\xa9\x52\x5e\x53\x2f\xd0\x4c\xc9\x40\x37\x85\x4b\x87\x23\x14\xed\x34\x5b\xde\x54\xb1\xa6\xbf\x82\xa2\x9d\xa2\xda\xac\x77\xe6\x9a\xaa\x40\x3b\x6b\x55\x53\x85\x9a\x3d\xed\xdc\xff\xee\xdf\x01\x00\x00\xff\xff\x8d\x40\xdf\xff\x13\x08\x00\x00")

func dataLocationsJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataLocationsJson,
		"data/locations.json",
	)
}

func dataLocationsJson() (*asset, error) {
	bytes, err := dataLocationsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/locations.json", size: 2067, mode: os.FileMode(436), modTime: time.Unix(1500835266, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataProfilesJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x5c\x5f\x73\xe3\x38\x72\x7f\x9e\xf9\x14\x5d\xf3\xe2\xbb\x2a\xd1\x25\x8a\xd4\xbf\xc9\xc3\x95\xed\x59\xcf\x4e\x6e\x7c\x71\xd6\xbe\x99\xdc\xa5\xf2\xd0\x22\x5b\x22\x46\x20\xa0\x05\x40\x69\x79\xa9\xab\xca\xd7\xc8\xd7\xcb\x27\x49\x35\x48\xd9\xfa\x03\xda\xba\xad\xac\x3d\x79\xd9\x5a\x8b\x24\x08\xa2\x7f\xe8\xfe\xf5\xaf\x1b\xf3\xef\x6f\xdf\xfc\xe7\xdb\x37\x6f\xde\x89\xfc\xdd\x7b\x78\x17\xbf\xeb\xf1\x1f\x0a\x4b\xe2\x3f\xaf\xa4\x98\x3b\xf8\x51\x3b\x92\xcd\x85\x55\xa1\x15\xfd\xa9\x2a\x67\x64\xf8\xfa\xef\xd2\x78\xf8\x7b\x18\x8f\x87\x51\x3a\xee\xf7\x9b\x5b\x72\xb2\x99\x11\x2b\x27\xb4\xe2\x5b\x2e\x60\x14\x95\x42\x55\x8e\x60\x83\x72\x09\x73\xa3\x4b\xf8\xb3\x12\x5a\xc1\xdd\xcf\x15\x1a\x02\x54\x39\xa4\xd0\xdc\x63\x9b\xeb\x08\x37\x95\x12\x70\x43\xce\x68\xb0\x0e\x79\xb0\x1e\xb8\x42\x58\x90\xd5\x2f\x95\xa9\xa1\xe0\x39\x41\x4e\x56\x2c\x14\xe5\x30\xab\xe1\xb6\x10\x52\xac\x56\x04\x77\x0e\x4d\xb6\x84\x39\xa1\xab\x0c\x59\x40\x05\x68\x9c\xad\x61\x5e\x19\x25\xf8\x37\xc8\xb4\x94\x94\xf1\xa8\x20\x14\xb8\x82\x40\xea\xd9\xac\xee\x81\x50\x99\xac\x72\xa1\x16\xb0\xd1\x66\xc9\xc3\xde\xa1\x5c\x63\xae\x0d\x7c\x40\x29\xce\x9b\x4f\xc4\x3c\x37\x64\xed\xbb\xf7\xc0\x4b\xf7\xe6\x9d\x75\x86\xc8\x3d\x2e\x4b\x3a\x1d\xfa\x1b\x1f\xae\xb4\xcb\xf9\x91\xd0\xd4\x70\xe7\xda\x8b\x99\x70\x35\xff\x7c\x87\x0a\xae\x0d\xaa\x4c\xd8\x4c\x3f\x3c\x88\xae\x31\xc1\xc5\xf6\x6e\x5d\x29\x67\xfc\x03\x7f\x56\xc2\x51\xce\x1f\xea\xc8\xb6\x97\x57\xda\x3a\x94\x57\x3a\xf7\x4f\x4d\xd3\xb8\x3f\x68\xaf\x48\x74\xef\xde\x43\x32\x3e\x1f\x4f\x46\xe3\xf6\x27\x6f\x9b\x28\x1e\x0c\xce\xd3\x38\x1e\xbc\x7d\xf3\xe6\xef\x6f\xdf\xfc\xbd\xb7\x07\x85\xc1\x3e\x14\xbe\xc2\xf1\x34\xbb\xe0\x30\x8e\x86\x49\x07\x1c\x3e\x93\xb5\xe0\x0a\xb6\x09\xcc\xa4\xce\x5a\x40\xb0\x09\xfe\x42\x66\x86\x70\x59\x91\x42\xb8\x22\xe5\xc8\xc0\x5c\x1b\x7f\xe9\xc2\x38\xdb\xda\xdf\x19\x52\xf9\xd6\xfe\xc2\x02\x42\x3c\x78\x1a\x60\x27\x1a\x2d\x9e\xc4\x41\xa3\x25\x26\x7f\x31\x93\x25\xc7\x26\x1b\xa6\xc7\x26\xeb\xf7\x87\x41\x93\x25\xfb\x26\xf3\xfb\x16\xfe\x4a\xce\xe1\x13\xe6\x1a\xa6\x49\x34\x19\x0e\x87\x1d\xbb\x37\x39\x5e\x5c\xb6\xc8\xad\xde\x90\x94\x70\xe7\x97\x0a\x32\x9c\x49\x8a\x32\x34\xe0\x2a\xa3\xd0\xe8\x4a\xe5\x7e\x5f\x5f\x5e\xfc\x74\x0f\x06\x85\x3c\xd8\xc5\x85\x58\xb5\x26\x9c\xee\xef\xfc\x3d\xcf\x90\xe9\x72\x26\x14\xf1\xed\x8b\x22\x72\x94\x15\x20\x75\xbe\xf0\xdb\x53\xb8\xa2\xdd\xd6\x4e\x57\x59\x41\xf6\x44\x33\x0f\xc3\x5b\x73\xe8\x8a\xd7\xb4\x72\x12\xb2\xf2\x38\x0e\x5a\x39\x0d\x59\xf9\x8b\x70\x28\xe9\x09\x33\x0f\xc6\x93\x28\xe9\x72\xd2\xf7\x6c\x93\x0d\x3a\x32\x73\xa3\x95\x6b\x4d\xe3\xd7\xf8\x12\x6b\xb8\x34\x22\x5f\x10\xac\x05\x6d\x2c\xef\xb9\xa4\xd9\xba\xf6\x11\x0e\xd7\x42\xf1\x32\xa1\x84\x0f\xc2\x3a\x23\x32\xe7\xcd\x8f\x90\x86\xe1\x73\x4d\xc6\xd4\x70\x59\x09\xc9\xce\xf6\x44\xd3\x4d\x82\x96\xbb\x11\xd6\x7a\xcc\xbc\x90\xf5\x86\x47\xd6\x9b\x26\xa3\x23\xeb\x25\xd3\xa4\x1f\xb4\xde\x70\xdf\x7a\xb7\x85\x26\x25\x7e\x39\x21\xc6\x8e\xa2\x38\x99\x74\x38\x55\x9d\x21\xcf\xbc\x8d\x65\xf7\xa4\x72\x32\x52\x0b\x05\x8a\xc4\xa2\x98\x69\x53\x68\x9d\xf7\xd8\x57\xf6\x8f\xed\x81\x9d\xbb\xd4\xf8\xf8\x5b\x6a\xa7\x8d\xdf\x78\x04\x05\x5a\x28\xb4\xe5\x97\x95\xa8\x6a\x30\xec\xc1\xcb\xca\x8a\x4c\xa0\xb2\xde\xe6\xda\x15\x64\x20\x23\x49\x33\x23\x9c\x20\x0b\x56\xa8\x8c\xfc\xcc\xe2\xe9\xb0\x6f\xcf\xe1\x93\xfb\x9f\xff\xfa\x6f\xdb\x8d\x8e\x42\x58\xa7\x8d\xc8\xe0\xa3\x21\x74\x70\x51\x92\x11\x19\x2a\xb8\xe1\x17\xc1\x8f\x28\x25\x28\xb1\x28\x5c\x26\xab\xd9\x89\xd8\x19\xf5\xc3\xde\xfd\x87\x3c\x7f\xb9\x88\x3c\x0d\x6c\xfc\x38\x10\x91\x27\xe1\x8d\x3f\xda\x87\xce\x9d\x3b\x87\x9f\x68\x21\xec\xc9\x91\x79\x30\x49\xa3\xb4\xdf\xe1\x03\x1e\x87\xbb\xa9\x2c\x55\x25\xdc\xeb\x0d\x99\x26\xc4\xa6\x83\x88\x2d\x52\xf7\x20\x9d\xa4\x30\x77\x60\x97\xb5\xcd\x0c\xae\xf8\x86\x06\x74\x77\xba\x72\x05\xe8\x39\xdc\xa0\x59\x92\x83\x7c\xeb\x07\xf4\x7c\x7f\x7e\x3d\xb8\x42\x29\xe6\xda\x28\x81\x3d\xc0\xfc\x1b\x66\xa4\x1c\x38\xbd\x47\x00\x3e\xa2\xc9\x49\xd9\x1e\xdc\x68\x9b\x69\x45\x2d\x23\xe8\xc1\x2d\x66\x97\x1c\x79\xb6\xbe\xc3\x63\xce\x4f\x60\xf7\x25\xdb\x6f\xe0\xf9\xe8\x9c\x8c\x62\x0a\x71\x2a\x11\x18\x84\x43\xc4\x4b\x12\x81\x00\x52\x46\x49\x20\x44\xc4\x61\x22\x30\xde\x47\xca\x7d\x41\xbc\x8e\x96\xcc\x13\xe0\x98\x4e\x46\x51\x9a\x76\x80\xe3\x83\xc8\xa1\xd6\x15\x94\xc2\x5a\xfe\x1f\x03\x19\x4a\xe9\x97\xdf\xfa\xd5\xb7\x5a\x2d\x66\xc2\xe4\xbc\xe2\xfe\xfa\x82\x14\x19\xef\x4c\xfe\x00\x37\x58\xcf\x88\x7f\x3e\xf3\x41\x5d\xcd\x45\xee\x4d\x5e\xa0\xf3\xbf\xae\x09\x66\x44\x0a\x66\x92\xac\xa5\xbc\x09\x3c\x3c\xea\x5a\x8b\x8c\x78\x48\xe6\x8b\x6a\x41\xb2\x07\xb3\xca\x41\x81\x6b\x52\x67\x0e\x4a\xcc\x09\x84\x83\x15\x5a\x07\x4b\x34\xa8\x97\xd4\x78\x06\x68\x89\xa3\x2d\x18\xc4\xe7\x70\xad\x8d\xab\x14\x3a\x92\x75\xeb\xdb\xbe\x88\x8c\x7d\x0c\x2a\x58\x19\xbd\x22\xe3\x6a\x88\x22\xf8\x56\x59\xf7\x04\x21\x89\x22\xd0\xf3\x39\x19\x0b\xa5\x36\xd4\xf0\x58\xff\x48\xa6\xff\xc6\xce\x50\x97\x27\xd3\x90\x34\x88\xb1\xf4\x75\x69\xc8\x30\xe0\x8d\xfa\xc3\xb0\x37\x9a\x1c\x04\x32\x34\x19\x0c\x87\x27\xfb\xa2\x64\x3a\x88\x26\x5d\xbe\xe8\x93\x3b\xb3\x50\xa0\xc9\xd9\x2d\x88\x12\x17\x42\x51\x03\x17\xdc\xe5\x26\x7a\x4d\x06\xe2\x5e\xbf\xdf\x6f\xd6\x1e\x32\x5d\xc9\x1c\xe6\x44\x12\xa4\x58\x72\x6a\x59\xbb\x82\x51\x2a\x77\xd2\x8e\x6f\x58\x46\x2b\xcc\x96\x94\xc3\xbd\x5e\xd6\x1a\x6c\x35\xdb\x60\x0d\x4c\x5f\xf3\xca\xf0\xed\xa6\xb2\x05\x14\xba\x32\xe7\x70\x59\x39\x60\xff\x63\x1d\xad\x40\x28\x2b\xf2\x26\x94\x95\x68\xad\x58\x13\xb8\x8d\x8e\x24\xad\xf9\x8d\x9c\x42\x32\x58\xf9\xb2\xa2\x8d\xac\xc1\x90\xd2\x6b\x1f\x9a\xb7\xcb\xf3\xb5\x56\x79\x81\xe5\x81\xb7\x12\x6a\x1f\x67\x1b\x21\x25\x94\xb8\xf4\xe8\xce\x24\x31\xb1\xe6\xaf\xe7\xd0\x4a\x67\xbb\xe8\x23\xa5\xab\x45\xe1\x3f\x9f\xd7\x6a\xc6\x01\xb3\xa0\xf3\xf3\xf3\xd6\x2d\x2a\xb0\xba\x3c\x35\x05\xea\xe0\xc6\x57\xb5\x11\x12\x6e\x70\xa1\xc4\xcb\xd1\xac\x40\xf6\x9a\x4e\x03\xe8\x9c\x8c\x83\xe8\x9c\xee\xa3\xf3\x8f\x6c\xd5\x8f\xa4\xcd\x82\x9e\xa7\x5a\x93\x38\x1a\xf6\x87\x61\x64\xfe\xe0\xf3\x52\xef\x41\x66\xba\x72\xe2\xe7\x8a\x5a\x44\xca\x96\x84\x35\x66\x82\x82\xd0\xf8\xe0\xb7\x6b\xd8\x9e\x8f\x56\xec\xf2\xa4\x04\x51\x96\x94\x0b\xef\x93\xa0\x62\xc6\x66\x1d\x5f\x15\xce\x02\x4f\xfb\x9f\xfc\x7f\x73\xc0\x79\xf3\x46\x62\x1a\xee\x84\x2d\xa0\xd4\x0a\x4d\x56\xc0\xee\x47\x7d\xe9\xc1\xa2\x22\xeb\x2c\x30\x7e\xbc\x63\x74\x7a\x67\x27\x18\x5d\xa3\x74\x35\xcc\x6a\xc0\x2c\xd3\x65\xa9\x73\x74\xfc\xbc\x75\x38\x9f\xf7\x60\x8d\x92\x1c\x48\xac\x54\x6e\x6a\xb0\x64\xd6\x22\xa3\x86\xd0\x0d\xd2\x88\xb7\x42\x83\xb1\xf6\xca\x89\x80\x4a\xc2\x5e\xee\x06\xed\x0b\x12\xf6\x90\x0e\x12\x42\xd2\x34\x8c\xa4\xb8\x7f\xa8\x89\x55\x33\xf8\xd7\x0a\x8d\xe3\x38\xf0\x1c\x98\xd8\xcd\x8d\xbb\xa2\xea\x05\xaf\x7e\xcd\x88\x09\x0c\xca\x4e\x61\xdf\x4b\xac\x24\xb2\x4d\x38\x0e\xb7\xa4\xeb\x01\x64\x07\x0c\x6b\x2f\x82\xdd\x1b\x54\xb6\xe5\xd0\x70\x5b\x1b\x2c\x45\x93\x9c\x67\x52\x5b\x0f\x93\x5b\x41\x06\x92\xe9\x39\xf8\x3c\x30\x8d\xac\x43\xf3\xa8\xb0\x3c\xdc\xf6\x59\x97\x33\xf4\x24\xc8\x27\xfd\x3c\xc4\x85\xcc\xd0\x19\xfc\x1b\x7c\xb2\x12\x55\x7e\xaa\x3e\x36\x08\xe3\xe2\x4a\xe2\x8b\x91\xf1\x38\x3e\xce\xe3\x86\x21\x8a\xd5\xef\x80\xc5\x81\x54\xfa\xd1\xa0\x72\x70\x2b\xf1\x6f\xf8\x2c\x28\xd2\x24\x8d\x92\xc9\x24\x79\x06\x14\x47\x43\xfe\x76\x08\xd8\x37\x6d\x0b\x84\x63\x04\xb4\x40\xf1\xcf\xfe\x59\x89\x35\x19\x2b\x5c\x7d\xf4\xfa\x53\x61\x30\xea\x90\x49\xfd\x77\x5f\xac\xe9\x65\xfc\xc3\xe4\x18\x08\x71\x08\x08\xc3\x0e\x20\x1c\x08\xa5\x3f\x0a\xe9\xf4\xe1\xd6\xdd\x0d\x02\x4f\x66\xf9\x71\x14\x77\x79\x8b\x7d\x9b\x34\x1c\x24\x5b\x42\x8e\x9c\xc2\x1e\x23\x80\xd4\x37\x5d\xb3\x87\xe7\x80\xc2\xcb\x54\x59\x26\xeb\xcc\x65\x5a\x8e\x22\xd1\x2c\xc8\x3a\x66\xb5\x62\x0e\x4a\xbb\xc3\x5f\xfd\xcb\xec\x16\x6a\x6c\x86\x73\xb8\x6e\x24\x83\x74\x14\xcd\xa5\x66\x96\xcd\x04\xbb\xe7\x21\xc9\x29\xba\x56\xb2\xde\x61\x6c\x04\x76\x45\x99\xc3\xac\x92\x68\x5a\x01\xa9\x7d\xfb\x0c\xeb\x53\xa3\x48\x92\x04\x71\xf2\x2f\x67\xd7\x68\x4c\x23\x43\xbe\x5a\x28\x19\x4e\x02\x09\x7c\x7f\x12\x86\xca\x81\x40\xfb\x15\xcd\x46\x64\xcb\x93\x39\xf3\x74\x30\x89\xc6\xd3\x0e\x78\x5c\x6b\x03\x99\x64\x56\x9a\xb5\x28\xd9\x09\xf4\x5a\x79\x33\x32\x61\x94\xb4\x40\xe5\x7a\x0f\xf4\x25\x6a\x6e\xb6\xe4\x98\x0f\xf4\xbc\x6d\x82\x13\x6b\x7d\x90\x2d\x3c\xc9\x9e\x11\x18\x9f\x69\x55\xab\x26\xe1\x43\x29\xa9\x66\xcf\x41\x8f\x19\xd5\x43\x0d\xe6\x91\x7f\xc7\xfd\x7e\xc3\x54\x5a\xca\xce\xae\xc4\x56\xc2\x91\xed\x01\xe6\xda\xa8\x6d\xfe\xf7\x59\x57\xc2\xc2\xbf\x7d\xf9\xd4\xd4\x6d\x2c\xf3\xf8\xe6\xf6\x9c\x32\x6d\x7a\x20\x98\x8a\x95\x42\x2d\xe4\xf6\x91\xb2\x49\xf5\xb1\x24\xd5\x28\x4f\x9e\xfa\xcc\x25\x3a\x9b\x19\x4e\x30\xef\xbf\xd8\x1e\x6c\x84\x21\x3f\x9f\x4f\x3c\x80\xe2\x48\x96\x65\xfc\x37\x8f\xcd\xcc\xfa\xe4\x2a\x4f\xff\xfb\xac\xf2\x8c\xfb\x01\x48\xc6\x49\x18\x92\x41\x35\xf9\xaf\x9a\x4e\x06\xe5\x70\x14\x47\x71\xdc\x95\xc8\xf1\x12\x93\x6d\xc5\xc9\x6d\xf2\xc6\xb9\xcc\x99\x6d\x72\xfe\xb5\x76\x4d\xfd\x8e\x81\xc3\xce\xa5\x24\x73\x66\xe1\x63\xc5\x89\x56\x2b\x2e\xcc\x18\x2e\x2d\xa6\x1d\x5c\x0b\x5b\x90\x29\x51\x9d\x59\xf8\x5a\xa0\x99\xf7\x80\x3c\x12\x6b\x42\xd3\xc0\x1c\x8c\xde\xfc\x01\x3e\xcd\xc1\xea\x06\xd0\x97\x3c\xc2\x57\x9e\x88\x51\x70\x5f\x59\x76\x55\x9f\x94\x62\x47\x1a\x44\xf1\x07\xad\xce\x1c\x30\x25\x3e\x7a\x9a\x57\x8a\x21\x48\x9c\xfd\xd5\xba\x7a\xdf\x64\x03\x7f\x14\xe5\x8a\xfd\xfe\x03\xf6\x5b\xa1\x00\xa5\xf4\x63\xac\xc8\x2c\xed\xa3\xe4\x91\x03\xfd\xc2\x9e\xf1\x3d\xe4\x28\x64\x0d\x1b\xf6\x97\x0c\xc0\x19\x91\x01\xab\x33\x81\x92\x77\x04\xac\xc9\xd4\xb0\x22\x17\xcd\x8d\x20\x95\xcb\x1a\x56\x5a\x8a\xac\x6e\x52\x09\x6d\x16\xa8\x44\x06\x19\xbf\x8d\xbc\xea\x41\x08\x84\x19\x6f\x06\xa3\x4e\xd7\xd6\xd3\x0e\xd1\xeb\x4f\xda\xb8\x02\x6e\xb5\x50\xee\xa5\x20\x9d\x1c\x09\x13\x93\xfe\x28\x09\xc9\x5f\x71\x3a\x48\x3b\x82\xf2\x30\x04\xeb\x2b\x34\x1c\x9b\x9f\x00\xf3\x68\x9c\x44\xfd\x41\x3a\x08\x8b\x60\x1e\x14\x33\x82\xb9\xd6\xb2\x01\x6d\x23\x6a\x0f\xc6\xc0\x09\x99\x70\x94\xb9\xca\xe0\xb6\x2e\xbd\x8d\x74\x7b\x2f\x7f\xbf\xab\x24\x3c\xe4\x83\xb0\xf1\x63\x17\xb8\xf6\xa2\x05\xe4\xba\x9a\xb9\x6d\xce\x2f\xec\x23\xaa\x18\xb8\x25\xb3\x6f\x8f\xd6\x41\x6c\x1d\x64\xa4\x5c\x65\x6a\x50\x44\xb9\x6d\xa9\x5b\x3c\x8a\x23\x9f\xaf\x35\xbb\xe6\xa1\x90\x86\x4d\xb5\x8c\x67\xe6\xdd\xa7\xf2\xb1\x01\x25\x3b\x49\x43\xdb\x48\x3f\xd7\xc6\x4b\xa7\x95\xf2\xf9\xad\x33\xb8\x66\xf2\x59\x68\xa7\x5b\x8d\x5f\x51\xa4\xe7\x11\x46\x4b\xc1\x2c\xb2\x92\xcd\x67\xeb\xd9\x37\xca\xdc\xa9\x12\x58\xdc\x1f\x87\x31\x77\x57\x39\x4e\x79\x5f\x51\x6b\x1d\x87\xea\xe4\x93\x70\x41\x27\x1e\x85\xa0\xf6\xb5\x10\x8e\x57\xfd\x29\xac\x0d\x46\xdd\x0a\xd8\xb5\x17\x30\x99\x76\x6f\x0a\x0d\x2b\x43\xf3\x56\x04\x58\x18\x54\x39\xb1\xa7\x52\x39\x38\x51\x36\xa1\xac\x89\xea\xad\x56\xba\x4f\x03\x7a\x3b\x18\xdc\xce\xaa\x51\xf5\x39\x63\x2b\xd1\x1c\x86\xf9\x07\xac\x35\xf0\xc3\x25\xd9\xad\xb6\x30\x63\xc2\xe9\xb4\x1f\xd1\x55\x46\x31\x08\x5c\x41\x51\x0b\xc1\x5e\x13\x86\x99\x70\x5e\x54\xd6\x79\x75\x35\x33\x35\x2f\x38\x64\x05\xcf\x5b\x0a\x32\xb6\xd7\x7c\x44\x2b\x98\x35\x91\x1b\xcd\x4c\x12\xac\x84\x94\xec\xc3\xf9\xea\xbd\x98\xcf\x79\x2b\x2c\xf8\x63\x7c\x88\x56\xa7\x43\x6b\x90\x84\xcb\x3d\x6d\x71\xe2\xf5\x24\xd6\xf1\x24\x14\x9c\x87\x1d\xd0\x0a\xe8\xf8\x17\x52\xd8\xd3\x83\x73\x9a\xc6\xd1\xa0\xab\xea\xfb\xd5\xd7\xd0\x1f\xd2\xcd\x3b\x87\xf5\x4a\x28\xc2\xd5\x4a\x12\xff\xd0\xf1\xba\xde\x56\xc0\x9a\x91\xf7\x3d\x06\xa5\xac\x1f\xe4\xaf\x43\xd5\xa2\x07\x99\x56\x6b\x52\xa2\x2d\xf2\x70\x18\x9d\x0b\x92\xf9\x01\xea\xae\x78\xa4\x26\x04\x06\x93\xd1\x4e\x55\xe2\xff\x36\x27\x1d\x4e\xbe\x53\x52\x17\x92\xac\xe2\x8e\x1a\x73\x7c\xa0\xcd\x7f\xe1\x6d\x05\xd7\x52\x1b\x52\xd9\x53\xf9\x67\x32\x1d\x47\xe3\x2e\xb8\xdc\xe9\x92\xc0\x62\xfd\x30\x10\x6c\xb6\xfc\x4c\x18\x57\x78\x6d\x62\x1b\xf2\x3e\x39\x94\xbc\xfb\x7f\x22\x85\xc2\x5a\x76\x4d\xe7\xcc\xc7\xda\x1a\x90\xd8\xa3\x85\xcc\x84\x98\xdc\x64\x9c\xad\x3e\xd6\x7f\xeb\x46\xa2\xf0\xc1\xe5\x61\xe0\xb3\x0b\x57\x90\x7a\x48\x22\x6f\x44\x9e\x4b\x82\x8b\x05\xd9\x33\xce\x58\xd9\x99\xe8\xca\x7b\x2d\x5f\x9c\xe2\xd8\xa5\x2c\x3a\x89\xca\x89\x0c\xe6\xd2\x93\xbd\x28\xf2\x0f\xef\x2f\x8b\xf7\x89\x9c\x14\x31\xeb\x6c\xd0\x85\x76\x0b\xf5\x05\x9d\x5a\x3d\x1c\x74\x10\xa9\x87\x86\x99\xd7\x43\xd0\x71\x97\x42\xda\x9f\x74\xa4\x05\xd3\x63\xcf\x73\x85\xc6\x6d\x1a\xb6\xdc\x44\x93\xe8\x54\x59\x23\x1d\xc4\xd1\x60\x32\x0a\xb7\x18\xfd\x89\xac\x93\x8f\xcd\x0b\xad\x3e\xf1\x73\x25\x88\x49\x00\x1a\x42\x7b\x2c\xa3\xb7\x66\xf1\x1a\xc7\x36\xe0\xd5\x8f\x40\x6a\xfc\x88\x50\x6b\xe1\x35\x6e\x9f\x30\x6e\x85\xf9\xcf\x68\x96\x76\x55\x99\xf6\x23\xf6\x3a\xc5\xbc\xcb\x6b\x5b\x01\xe7\xda\x00\xfd\x82\xe5\x4a\x52\xcf\x0b\xe9\x6d\xd9\xa5\x29\x2c\x81\x14\x6b\x5f\x29\xf2\x85\x97\xa6\xb0\xd4\x3e\xe8\x03\x5a\x6e\x70\x45\x8c\x5e\xe3\x50\xb4\x9d\x11\x99\x2e\xe7\xda\x38\xe4\x48\x67\x85\x4f\xb7\xdb\xef\x53\x84\x06\x10\x36\x68\x4a\x98\x0b\x43\x7e\x2f\x9d\xea\xb1\x3a\xc4\xd4\x17\x66\x51\x01\xc0\x4d\x43\x4d\x4d\xd3\x8e\x6e\xc3\x7e\x88\x45\xdd\xe0\x37\xb2\x4e\x64\xcf\x44\xb8\xce\xf4\xf3\x20\xc2\xed\x0f\x7a\x2c\xb0\xff\xee\x4a\xac\x45\xd6\xb6\x1a\xfc\x7e\x2f\xce\xed\xc5\xb0\xdd\xdb\x5a\x7d\x96\xed\xfb\xd0\xaa\xe0\x6f\xf6\xdd\xa5\xcd\x3d\x87\x7a\x2a\x23\xa4\x85\xfb\x5a\x64\x42\xb5\xf1\xea\xa3\x96\x39\x29\xf8\x88\x8e\xda\x96\xab\x53\xe9\xce\xb0\x1f\x8e\x5a\xdf\x01\x93\x1e\xa4\x81\xc4\x6d\x30\x1c\xf6\x47\x61\x1c\xc4\x87\x85\xbb\x26\xad\xbe\x13\x66\x3b\x63\xf8\x60\x70\xf9\x7c\x19\xcf\x57\x5e\xc6\x5d\x7d\x8d\xab\x95\xd1\x9c\x25\x9f\xf1\x8e\xf7\x03\x9e\x35\x5a\xf9\x7e\x63\xe3\x61\x61\xb6\xa3\x7c\x67\x57\xad\x78\x3a\xc7\x52\x57\x16\x2e\x89\xe6\x84\xbc\xf4\xb9\xd6\xa6\x24\xdf\x66\x3c\x6b\x7c\xa7\xa1\x9c\x13\x2b\xce\xb1\xec\x39\x7c\xa6\x05\xa9\xdc\xf7\x55\x89\x36\xe1\x63\x08\xf9\xc7\x50\xf5\xe0\x5e\x97\x70\xb7\x21\x52\x54\xf7\x1e\x64\x5f\x6d\x5d\x93\x8d\x2d\x0c\xae\x0a\xca\x61\x45\xc6\x36\xbd\xcc\xfb\x9c\x07\x4e\x4c\xfe\x87\x61\xf8\xbc\x7e\xcc\x9a\x4c\x27\x83\x60\xd4\xea\x72\x23\x87\x62\xbc\xdf\x72\x7f\xc4\x59\xb5\x14\x4f\x0a\xab\x83\x28\x19\x74\x38\x91\x1f\xca\x99\xc1\x6c\xab\xa4\xff\x33\xae\x50\x39\xbd\x51\x60\x2b\xe3\x1b\x5d\x85\x5a\xd8\xdd\x24\xab\x79\xdb\xd6\x58\x2b\x32\x73\xca\x1c\x64\x85\xef\x65\x11\x0f\x1c\x48\x6a\xed\x39\x0a\xc7\x99\xbf\x92\xf2\x4c\xe5\xe7\x4a\x48\x76\x06\x6d\x1b\x82\xcf\xf1\xf7\x5d\xd4\x82\x1c\x6e\xb0\x3e\x87\x4f\xce\xf2\x23\xb9\x68\x13\xf8\x56\x14\xb5\x9c\xdd\x33\x98\x54\x93\xff\xdb\x42\x7f\x13\xd0\xc8\x9f\xb6\xed\xc9\xf4\x9f\x40\x96\x60\xe1\x9b\xac\x9a\x80\x35\xd3\xca\xa2\x00\xb6\x7f\x73\xdf\x52\x0b\x98\x0b\x5b\xf4\xa0\x20\xb9\x62\xb7\x97\x19\xc6\xb4\x17\x91\xd5\x5a\x18\xad\xca\x6d\x17\xcf\x99\x85\x92\xb0\x71\x8e\xfe\x6e\xaf\x40\x35\x32\x2f\x66\x85\xa0\x35\x41\x89\xbf\x88\xb2\xf2\x15\x64\x62\x8f\x77\xaa\x6f\x1b\x75\x12\x2a\xfb\x62\x89\x5c\x7c\xdc\xf4\x39\x19\x0e\xd3\x80\x4a\x30\x98\xa4\xe3\x0e\xcf\x76\x20\xfe\xdf\xd4\x3e\x08\x3d\x5b\x2b\xec\xf7\xa3\xfe\xf0\xb4\xf8\xb6\x3b\x24\xcc\x6a\xb8\x2a\xd0\x48\x41\x70\x8b\xb2\x6c\x8b\x34\xbf\x3a\x6f\xfb\x68\x98\xde\x5f\xa1\x2b\x28\x37\x28\xbb\x13\xb5\x83\xb6\x88\xdf\xb0\x7c\xd8\xa6\xc7\x47\x41\xcf\xe9\x6c\xe9\x5e\xb0\xc1\xe0\xb8\x80\x38\x99\x06\x1b\x3b\xfb\xe3\x24\x4e\xc3\xe0\x38\x90\xe1\xf7\x82\xce\x69\x45\xe5\xf1\x78\x14\x8d\x87\x93\xae\x3e\xfe\x2d\x46\xba\x46\xfe\xb5\xb5\xe5\x13\xf2\xf9\x57\xae\x32\x27\x83\xef\x33\xa3\x1f\x0f\xa6\xa1\x9c\xbe\x9f\x8c\xa7\x61\x8c\x1c\x68\xda\x17\x39\x2a\xf7\x3c\x0f\x1a\x8d\x93\x68\x3a\xd8\xf2\xaa\xa7\x1d\xc8\xee\x90\x3d\x40\xb8\x8a\xee\x37\x7a\xfb\xe7\x8b\x78\x8f\xdf\x1a\x0d\xa3\xf8\x7b\xd5\x77\x26\xe3\x50\x2b\x78\x32\xea\x38\x05\x32\x18\x85\x89\xf2\x65\x95\x15\xa8\x50\x3d\x8b\x8b\xe9\x20\xee\x6e\x06\x3f\xc0\x45\x78\xf0\x5f\x9f\x40\x5d\x48\x2c\xf5\x2e\xa7\x7e\xca\xb0\x5d\xc0\x08\xb4\x23\xdd\xe2\x56\x84\xba\x16\xaa\x39\x36\x76\xf2\xe9\xaf\xef\x22\x85\x0a\x11\x8d\xd1\x38\x0d\x70\xe0\xa4\x3f\xe8\x68\xcc\x1d\x8c\x43\x1c\xf8\xde\x88\xa7\x4b\x5f\xc9\x34\xed\x26\x1a\xd7\xdb\xa3\x19\x0b\x2d\x73\xc8\x9a\x13\x57\x07\x67\x27\xdb\xba\xc0\xbc\x52\xcb\xba\xad\x80\x35\x4c\x72\xd6\x3c\x23\xb5\xb1\x20\x05\xc3\x40\x2d\xa0\x5a\xf9\xbb\x7d\x13\x40\x0f\x04\x53\x48\xa5\x99\x53\xaf\x8c\xb0\x6d\xb3\xef\xee\xcc\xbd\xd8\xb8\x7b\xde\x13\x61\x61\x74\xb5\x3a\xc2\x0b\xa0\x71\xc2\x3a\xdb\x88\x39\x94\xe9\xc7\xca\x69\xd3\x71\xe0\x8b\x6c\xae\xa0\x72\xdb\x20\x70\x76\x6f\x2a\xeb\x3c\x11\xbf\xad\x66\x52\x64\xf0\x19\x55\x7e\xe6\xe7\x7e\xf6\x23\xae\x56\x35\x07\x4a\x45\xae\x77\x06\x9b\x42\xc8\xa6\xb1\xe5\x8c\x49\xfb\x07\x52\xf6\xac\x29\xf5\x6e\xe5\x1d\x2f\x62\xb6\x0b\x84\xa2\x7c\xec\x08\xb5\x5a\xea\xb6\x96\x46\xe6\xe4\x06\xca\x8e\xe0\xf5\xea\x2d\x52\xfd\xa4\x1f\x6a\xa2\x1c\xa6\xfd\xb0\x26\x3d\x38\xd0\xa4\x9b\x2f\x78\x96\xfd\x0e\xe2\x68\x3c\x4c\x9f\x6b\x9f\xdc\x19\xec\xd7\xf2\x98\x50\x84\x7a\x5d\xe2\x32\x1e\x86\xdb\x9e\x2e\x2b\xfb\x72\x87\x04\x42\xdc\x76\x3c\x0e\xb5\xc7\x4d\x87\xd3\x0e\xde\x72\xa0\x25\xff\xa4\x73\xda\x60\xed\xfb\x2f\x76\xe3\xc5\x53\x30\x18\xa7\x51\x9a\x8c\xd3\x53\x62\x55\xd7\xf0\xff\x48\xb4\xda\x43\xc6\x01\xaf\x65\x8b\xfa\x23\x68\xbf\x06\x21\x4f\xa0\xa2\x09\x60\x6d\x27\xee\x69\x08\x99\x8c\xc2\x31\xeb\xb5\x0f\xb5\x0d\x26\xe3\x80\xf4\x3b\x88\xa7\x49\x18\x21\x49\x50\xfc\xbd\x14\xdf\x74\xf5\x64\xc2\x13\x47\x71\x97\x68\xf3\x97\x7f\x80\xac\xb6\x87\x89\x0f\x64\x62\xff\xfa\x9e\xef\xd6\xf8\xff\x92\xea\x3c\x36\x35\x7f\x67\xfd\xf6\x69\x3a\x1a\x86\x5c\xc6\xa8\xe3\xfc\x5a\x72\xa0\x02\xff\x50\xce\x7c\x81\xe7\x94\x5c\x27\x4e\xfb\x61\x3f\xb1\x37\x08\x5b\xe1\x99\xf0\xb0\x41\xe9\x05\xba\x5c\x58\xe7\x7b\x2b\x3c\x08\x2e\xac\x40\x7f\x9a\x70\xe7\x88\xe1\xf1\x3e\xee\xf0\x19\xff\x08\x97\xfd\xed\x72\x9d\x5b\x2d\x97\xaf\x89\x86\x41\x92\x84\x9a\x20\x26\x93\x51\x87\x7b\x38\x10\x75\x3f\xa2\x94\x64\x04\xc2\x2d\x9a\xe5\x49\xc7\x79\x92\xfe\xe8\xe9\x93\xd3\xfe\x10\xdf\xf6\x5f\xa4\xd0\xf3\x87\x53\xac\xad\x3d\x7c\x7d\x6e\xbf\x02\xc9\xc8\x09\x4c\xa4\xe9\xac\x59\xf8\x73\xcb\xad\xe4\xeb\xff\x11\x8b\x96\xf6\x59\xb0\xe4\x5b\xc5\x5a\xe1\x17\x15\xe0\x7c\xae\x4d\xee\x8b\x82\x07\xd2\xd9\x4e\x39\xf3\x49\xa8\x36\x4c\xb7\x79\x46\x1b\xb1\x10\xca\xbb\x3b\xbd\x22\xd5\x78\xbb\x78\x1a\xc7\xdb\x5e\xca\x66\x96\x4d\x56\xd3\xb4\x9f\x35\x07\x8c\x10\x0c\xf9\x83\xb8\x7a\xa3\xc8\xd8\x42\xac\x7c\xcb\xce\xa2\xf1\x6b\xdb\xd3\x6e\xed\x81\x36\x7f\x44\x5c\xb8\x6d\x8f\x63\xa6\x95\xa3\x72\xa5\x0d\xe7\xd0\x07\x5d\xc7\xcd\x59\xce\x96\x0f\x7b\x56\x6f\x68\x2e\x29\x73\xcd\x41\x27\x42\x23\xeb\x6d\x0b\x5b\x24\xb5\x5a\x6c\xab\xff\xa7\xfa\xbc\x69\xd8\xe7\xbd\x70\xe9\x2b\x0d\xf0\xa4\xd1\x34\xe4\xf4\x92\x34\xdd\x0a\xc4\x6f\xff\xe3\xed\xff\x06\x00\x00\xff\xff\x94\xd3\xad\xf5\x7f\x47\x00\x00")

func dataProfilesJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataProfilesJson,
		"data/profiles.json",
	)
}

func dataProfilesJson() (*asset, error) {
	bytes, err := dataProfilesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/profiles.json", size: 18303, mode: os.FileMode(436), modTime: time.Unix(1500835266, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataRatesJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8e\x41\xcb\x82\x30\x1c\xc6\xef\x7e\x8a\x3f\x3b\xab\x6c\x7a\xf0\xf5\xbd\x45\x41\x88\xd0\x41\xba\x45\x07\x75\xa3\x46\x6b\x7f\xd1\x19\x54\xf4\xdd\xc3\x49\xea\x82\x76\x78\x0e\xcf\xef\xb7\xed\x39\x78\x00\x00\x4f\x9b\xc3\x21\x67\x34\x42\x65\x9c\xfc\x03\x61\xc4\x9f\xfb\x1a\xb9\x18\xca\x62\xb5\xce\x97\xbd\xd4\x9b\xd2\x58\x12\x51\x96\x04\xf4\x2f\xa0\xe9\x92\x63\x6f\xbe\x85\xd8\x79\xb8\x45\xbc\xee\xef\xcd\x60\xcc\x33\x00\x48\x85\x78\x29\x2b\x25\x8a\xf1\x36\xa3\x69\x48\xa9\xbf\x34\x3e\x93\xf2\xdd\x96\x38\x80\x8b\xae\x6e\x65\x63\x24\x6a\xcb\xa5\x3e\x41\x27\x1f\x82\x43\x25\xb8\xab\x1a\x34\xa5\xfa\xfd\xc5\x84\x33\x5d\xab\xbe\x93\x37\xeb\x45\x71\xc8\x92\x49\x7b\x79\x63\x1e\xbd\x77\x00\x00\x00\xff\xff\xdc\x97\x83\xfb\x4d\x01\x00\x00")

func dataRatesJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataRatesJson,
		"data/rates.json",
	)
}

func dataRatesJson() (*asset, error) {
	bytes, err := dataRatesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/rates.json", size: 333, mode: os.FileMode(436), modTime: time.Unix(1501109263, 0)}
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
	"data/locations.json": dataLocationsJson,
	"data/profiles.json": dataProfilesJson,
	"data/rates.json": dataRatesJson,
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
	"data": &bintree{nil, map[string]*bintree{
		"locations.json": &bintree{dataLocationsJson, map[string]*bintree{}},
		"profiles.json": &bintree{dataProfilesJson, map[string]*bintree{}},
		"rates.json": &bintree{dataRatesJson, map[string]*bintree{}},
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

