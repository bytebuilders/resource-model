// Package crds Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// cluster.bytebuilders.dev_clusterauthinfotemplates.yaml
// cluster.bytebuilders.dev_clusterinfos.yaml
// cluster.bytebuilders.dev_clusters.yaml
// cluster.bytebuilders.dev_clusteruserauths.yaml
package crds

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _clusterBytebuildersDev_clusterauthinfotemplatesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x4d\x8f\xdb\x36\x13\xbe\xfb\x57\x0c\xf0\x1e\x72\x89\x6d\x04\x6f\x51\x14\xbe\x2d\x36\x41\x60\x24\x41\x8d\xec\x26\x77\x8a\x1a\xc9\xac\x25\x92\x1d\x0e\xbd\x71\x8a\xfe\xf7\x82\xd4\x87\x65\x49\xf4\xba\xdd\xa6\xba\x91\x1c\x0e\x9f\x67\xbe\x38\x94\xb0\xea\x2b\x92\x53\x46\x6f\x40\x58\x85\xdf\x18\x75\x18\xb9\xd5\xe1\x17\xb7\x52\x66\x7d\x7c\xb3\x38\x28\x9d\x6f\xe0\xde\x3b\x36\xf5\x67\x74\xc6\x93\xc4\xb7\x58\x28\xad\x58\x19\xbd\xa8\x91\x45\x2e\x58\x6c\x16\x00\x92\x50\x84\xc9\x47\x55\xa3\x63\x51\xdb\x0d\x68\x5f\x55\x0b\x80\x4a\x64\x58\xb9\x20\x03\x20\xac\x5d\x1d\x7c\x86\xa4\x91\x31\x9e\xa2\x45\x8d\x1b\xc8\x4e\x8c\x99\x57\x55\x8e\xe4\x16\x00\xcd\xa4\xac\xbc\x63\x24\xe1\x79\xaf\x74\x61\x18\x6b\x5b\x89\xb0\xad\x5d\x58\x0d\x77\xad\x72\x3c\x2e\x9c\x45\x19\x0e\x2a\xc9\x78\xdb\x2b\x98\xca\x35\x27\xb4\x98\xa4\x60\x2c\x0d\xa9\x6e\xbc\x84\x33\xc2\x76\x82\x5a\xee\xcb\xda\xe4\x58\xb5\x93\xc2\x5a\x27\x4d\x8e\x71\xd8\x5a\xaa\x39\xf0\xce\xf3\x7e\xab\x0b\xf3\xd8\x22\x8e\x12\x95\x72\xfc\xe1\x9a\xd4\x47\xe5\x38\x4a\xda\xca\x93\xa8\xd2\xfc\xa3\x90\x53\xba\xf4\x95\xa0\xa4\xd8\x02\xc0\x49\x63\xb1\x3f\x6f\x01\x70\x6c\x1c\x1e\x89\x2e\x5b\x2b\x1f\xdf\x88\xca\xee\xc5\x9b\x46\xab\xdc\x63\x2d\x1a\x3b\x00\x18\x8b\xfa\x6e\xb7\xfd\xfa\xff\x87\x8b\x69\x00\x4b\xc6\x22\x71\x6f\xb2\xe6\x1b\x44\xd4\x60\x16\x20\x47\x27\x49\x59\x8e\xa1\xf6\x2a\x28\x6c\xa4\x20\x0f\xa1\x84\x0e\x78\x8f\x1d\x34\xcc\x5b\x0c\x60\x0a\xe0\xbd\x72\x40\x68\x09\x1d\x6a\x8e\xe1\x75\xa1\x18\x82\x90\xd0\x60\xb2\xdf\x50\xf2\x0a\x1e\x90\x82\x1a\x70\x7b\xe3\xab\x1c\xa4\xd1\x47\x24\x06\x42\x69\x4a\xad\xbe\xf7\xba\x1d\xb0\x89\x87\x46\x73\xf2\x48\xa7\xd2\x8c\xa4\x45\x05\x47\x51\x79\x7c\x0d\x42\xe7\x50\x8b\x13\x10\x86\x53\xc0\xeb\x81\xbe\x28\xe2\x56\xf0\xc9\x10\x42\xb0\xfe\x06\xf6\xcc\xd6\x6d\xd6\xeb\x52\x71\x97\x49\xd2\xd4\xb5\xd7\x8a\x4f\x6b\x69\x34\x93\xca\x3c\x1b\x72\xeb\x1c\x8f\x58\xad\x9d\x2a\x97\x82\xe4\x5e\x31\x4a\xf6\x84\x6b\x61\xd5\x32\x42\xd7\x1c\xd3\xb1\xce\xff\xd7\xc5\x9f\x7b\x75\x81\x95\x4f\xc1\xbd\x8e\x49\xe9\x72\xb0\x10\x63\xf1\x8a\x07\x42\x14\x82\x72\x20\xda\xad\x0d\x8b\xb3\xa1\xc3\x54\xb0\xce\xe7\x77\x0f\x8f\x7d\xe8\x47\x67\x8c\xad\x1f\xed\x7e\xde\xe8\xce\x2e\x08\x06\x53\xba\x40\x6a\x9c\x58\x90\xa9\xa3\x4e\xd4\xb9\x35\x4a\x73\x1c\xc8\x4a\xa1\x1e\x9b\xdf\xf9\xac\x56\x1c\xfc\xfe\xbb\x47\xc7\xc1\x57\x2b\xb8\x17\x5a\x1b\x86\x0c\xc1\xdb\x5c\x30\xe6\x2b\xd8\x6a\xb8\x17\x35\x56\xf7\xc2\xe1\x0f\x77\x40\xb0\xb4\x5b\x06\xc3\xde\xe6\x82\x61\x65\x1c\x0b\x37\x56\x1b\x2c\x74\x55\xeb\xfc\xcd\xe7\x57\xf8\x42\x96\xef\xc8\x1c\x55\x8e\x34\x5e\x1b\xf9\xf9\x6e\x20\x1a\xcf\x50\x85\xc2\xe0\x75\x19\xeb\x79\x54\x15\x28\xca\x98\x58\xa1\xea\x94\x6a\x9c\x5f\xe1\x2b\x0c\x45\x6f\x9d\xab\x62\x5f\x58\x27\xd2\x69\xe0\x11\xfc\x77\x4f\x78\x6f\x74\xa1\xca\xb9\xe5\xe7\xb6\x43\x53\x63\x5c\x0c\xb1\xa5\xca\x53\x32\x49\xaf\x0c\xbf\x26\xf6\x5e\xac\x25\x92\x89\xd7\xc2\x8b\xf4\xa0\x3e\x2a\x32\xba\x46\xcd\x2f\xd2\xc3\xa8\xc5\x4b\x59\x85\xc4\x53\x84\x09\x15\xcb\x0b\x17\x24\x44\x7a\xe3\x26\xd6\x07\x74\x13\x12\x3d\x91\xd9\xf5\x44\x22\x75\x5f\xbc\xd8\xae\x6c\x4c\x90\x4f\x13\x6f\xee\xca\xc9\xf4\x15\x18\x32\x84\x71\x11\x72\x0b\x43\x1e\x1a\x52\x7c\x7a\x3b\x29\x08\x30\xce\xd9\xfb\xc4\xb6\x10\x68\x2c\x94\x76\xb0\x7b\xf7\x69\x89\x3a\xf4\x1d\xf9\xf0\x90\xb9\x6c\xeb\x14\x0c\xe5\xdc\x34\x67\x0b\x43\xb5\xe0\xa6\x07\x4b\x50\x9c\x35\x58\xe3\xe4\x01\xe2\x5b\xf8\xcd\xed\x99\x27\x17\xca\x67\xbc\x37\xe6\xa8\xb5\x87\x47\x66\x50\xa8\x0a\x63\x95\x7a\xfc\xf8\xf0\x6f\xd3\xfb\x80\x37\xb9\x6d\x28\xfb\x0c\x9d\x1e\xfc\x0c\xaf\x03\x9e\x7e\x10\x1b\x55\x5b\x24\x67\xb4\xe0\x99\xcc\xb8\xe0\xb2\x3d\x4b\x86\x26\x21\x54\x7e\xef\x42\x43\x54\x63\x68\x9b\x84\xe4\xa5\x98\x89\xa2\xdb\xce\x7e\x1f\x1a\xf3\x99\xe2\x9e\x42\xd0\xc8\x77\x38\xca\x66\xc4\x66\xa8\x72\x0a\x45\x31\xd6\xb3\x17\xc8\xd5\xec\x6f\x16\x05\x91\x38\xa5\xf1\x7f\x71\x48\xef\xbe\x31\xcd\x04\x84\xc8\xf3\xf8\x2a\x12\xd5\xee\xea\x2d\x76\x41\x35\xea\xfa\x1a\xdb\xaf\x5a\xb8\x43\xdb\x0a\xc7\xb1\x33\xe1\x3a\x64\x93\xf9\x02\xa4\xd0\x50\xa2\x46\x9a\xcf\xf4\x24\xe1\x67\xab\x7d\x9a\x74\xda\x27\xbd\x0d\xce\x81\x7e\xe6\x1e\x7b\xb0\x10\xa2\xd3\x5e\x1d\xda\x5e\x62\x60\xce\x3c\x86\x56\x2a\x98\x66\x0b\x6b\xe8\x42\x52\x2d\xc4\x05\xe2\x0f\xbd\xe0\x19\x67\xa1\x02\xc4\xa0\xa2\xbd\xb9\xfb\xe6\x46\xf6\x6f\xa4\x9b\xed\x67\x9e\x34\xd2\xf6\xed\x14\x46\x97\xa3\x4a\xf3\xcf\x3f\x25\x34\x86\x47\x46\x39\x39\xcf\x0a\xe7\x9e\x0c\xcd\xdc\x40\x17\xd4\x76\xad\x58\x97\x17\xdd\xb6\xc8\x26\x13\x4e\xc9\x71\x73\xc7\x66\xce\xf9\x37\x36\x75\x57\x6c\xc0\xe6\x80\xfa\x19\xb4\x8f\x41\xa6\x83\x9a\xa1\xa0\xf0\x28\x88\x73\x01\xee\x04\xe8\x08\xd6\x0c\xee\x7f\x02\xd4\xcf\xb5\x44\xd7\xe4\xdb\x92\xf7\x0c\xb7\x2f\x5d\x65\x1c\x57\xca\xff\xd8\x13\xf3\xad\xcb\xb2\x8b\xd1\xd1\xac\x1f\xb5\x55\xa9\x77\x09\x0b\xf6\xee\xd6\x97\x89\xc9\x62\x57\x98\xbf\x6f\xea\xd4\xe4\x2f\xc0\xc4\x76\xbf\x4e\x36\x74\x56\xac\x8d\x8b\xaf\xf6\x70\xc3\x97\xe7\xd5\xee\x84\xe4\x1b\x25\xfe\x2d\x68\x9e\xab\x2b\xd8\x32\x48\x43\x84\xce\x1a\x9d\xf7\x2f\xfe\x6e\xfd\x95\x1b\x68\x7e\x3d\xa3\xf1\x69\xaf\xe4\x3e\x00\x6a\x1f\x9c\x60\x34\xd4\xbe\xf9\x09\x01\xd9\x29\x2a\xbb\xdb\x6d\xdb\x07\x6f\xfa\xa2\xfe\xbb\x45\x60\xd6\x17\x93\xc9\xc6\x10\x1b\x60\xf2\xcd\x85\xe0\xd8\x90\x28\x71\x38\xe3\xb3\xfe\xb7\x41\xe7\x89\xd6\xa3\xf0\xc7\x9f\x8b\xbf\x02\x00\x00\xff\xff\x3e\x4f\xaa\xa9\xff\x13\x00\x00")

func clusterBytebuildersDev_clusterauthinfotemplatesYamlBytes() ([]byte, error) {
	return bindataRead(
		_clusterBytebuildersDev_clusterauthinfotemplatesYaml,
		"cluster.bytebuilders.dev_clusterauthinfotemplates.yaml",
	)
}

func clusterBytebuildersDev_clusterauthinfotemplatesYaml() (*asset, error) {
	bytes, err := clusterBytebuildersDev_clusterauthinfotemplatesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cluster.bytebuilders.dev_clusterauthinfotemplates.yaml", size: 5119, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clusterBytebuildersDev_clusterinfosYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4d\x6f\xe3\x36\x10\xbd\xfb\x57\x0c\xd0\x43\x2e\x6b\x19\x41\x8b\xa2\xf0\x2d\xc8\x16\x85\xd1\xaf\x60\xb3\xd8\x3b\x45\x8e\x2d\x36\x14\x87\x9d\x19\x7a\x9b\x16\xfd\xef\x05\x29\xc9\x96\xed\x6c\x11\x03\x59\xde\x34\x1f\x8f\xc3\x37\x8f\x43\x99\xe4\x3f\x21\x8b\xa7\xb8\x06\x93\x3c\xfe\xa5\x18\xcb\x97\x34\x4f\x3f\x48\xe3\x69\xb5\xbf\x5d\x3c\xf9\xe8\xd6\x70\x9f\x45\xa9\xff\x80\x42\x99\x2d\xbe\xc7\xad\x8f\x5e\x3d\xc5\x45\x8f\x6a\x9c\x51\xb3\x5e\x00\x58\x46\x53\x8c\x1f\x7d\x8f\xa2\xa6\x4f\x6b\x88\x39\x84\x05\x40\x30\x2d\x06\x29\x31\x00\x26\xa5\xe6\x29\xb7\xc8\x11\x15\xeb\x2e\xd1\xf4\xb8\x86\xf6\x59\xb1\xcd\x3e\x38\x64\x59\x00\x0c\x46\x1b\xb2\x28\xb2\x8f\x5b\x92\x66\xfc\x68\xe6\x91\x8d\xc3\xfd\x42\x12\xda\x02\xbe\x63\xca\xe9\x90\x74\x19\x37\xa0\x8e\x75\x58\xa3\xb8\x23\xf6\xd3\xf7\x12\x8e\x55\x8d\x06\x1e\xcf\xbb\xec\xc9\x61\x18\x8d\x26\x25\xb1\xe4\xb0\x7e\x8e\xec\x0c\x1b\x6e\xe2\x96\xaa\x35\x78\xd1\x9f\xcf\x3d\xbf\x78\xd1\xea\x4d\x21\xb3\x09\xa7\x67\xab\x0e\xf1\x71\x97\x83\xe1\x13\xd7\x02\x40\x2c\x25\x3c\x60\x2d\x00\xf6\x43\xd3\x6a\xe1\xcb\x91\xa9\xfd\xad\x09\xa9\x33\xb7\x03\x92\xed\xb0\x37\xc3\xb9\x00\x28\x61\xbc\x7b\xd8\x7c\xfa\xf6\xf1\xc4\x0c\x90\x98\x12\xb2\x1e\x28\x18\xd6\x4c\x15\x33\x2b\x80\x43\xb1\xec\x93\x56\xb9\xdc\x14\xc0\x21\x0a\x5c\x91\x03\x0a\x68\x87\x53\x69\xe8\xc6\x1a\x80\xb6\xa0\x9d\x17\x60\x4c\x8c\x82\x51\xab\x44\x4e\x80\xa1\x04\x99\x08\xd4\xfe\x81\x56\x1b\x78\x44\x2e\x30\x20\x1d\xe5\xe0\xc0\x52\xdc\x23\x2b\x30\x5a\xda\x45\xff\xf7\x01\x5b\x40\xa9\x6e\x1a\x8c\xe2\x48\xee\x71\xf9\xa8\xc8\xd1\x04\xd8\x9b\x90\xf1\x1d\x98\xe8\xa0\x37\xcf\xc0\x58\x76\x81\x1c\x67\x78\x35\x44\x1a\xf8\x95\x18\xa1\xb0\xbe\x86\x4e\x35\xc9\x7a\xb5\xda\x79\x9d\x6e\x83\xa5\xbe\xcf\xd1\xeb\xf3\xca\x52\x54\xf6\x6d\x56\x62\x59\x39\xdc\x63\x58\x89\xdf\x2d\x0d\xdb\xce\x2b\x5a\xcd\x8c\x2b\x93\xfc\xb2\x96\x1e\xb5\x5e\xa9\xde\x7d\x33\xe9\x49\x6e\x4e\x6a\xd5\xe7\xd2\x5e\x51\xf6\x71\x37\x73\x54\x6d\xfd\x4f\x07\x8a\xc2\xc0\x0b\x98\x31\x75\x38\xc5\x91\xe8\x62\x2a\xec\x7c\xf8\xf1\xf1\xe3\x41\xca\xb5\x19\xe7\xec\x57\xde\x8f\x89\x72\x6c\x41\x21\xcc\xc7\x2d\xf2\xd0\xc4\x2d\x53\x5f\x31\x31\xba\x44\x3e\x6a\xfd\xb0\xc1\x63\x3c\xa7\x5f\x72\xdb\x7b\x2d\x7d\xff\x33\xa3\x68\xe9\x55\x03\xf7\x26\x46\x52\x68\x11\x72\x72\x46\xd1\x35\xb0\x89\x70\x6f\x7a\x0c\xf7\x46\xf0\xab\x37\xa0\x30\x2d\xcb\x42\xec\xeb\x5a\x30\x9f\x6e\xe7\xc1\x03\x6b\x33\xc7\x34\x85\x8e\xeb\xe5\xfb\x55\x56\x1d\x95\xe8\xee\xf4\xdc\x01\xb0\x25\xee\x8d\xae\x8b\x7e\xbf\xff\xee\xc2\x3b\xec\x5d\xb4\xbd\xab\xa3\x60\xbe\xa6\xa6\x5c\x62\x7e\xe1\x78\x65\xf9\x3e\x11\xbf\x79\x29\x81\x6c\xbd\xe7\x57\x95\x52\x07\xd9\x35\x09\xf4\x39\x22\x6f\xde\xbf\x65\xe1\x15\xf2\xb7\x6b\x0b\x49\x4c\x45\x0c\x57\xe5\x64\xef\xae\x88\x2f\xd7\xc8\x33\x9e\xa5\x0c\xc3\xff\xcc\x34\xb2\x72\x66\xcd\xde\xbd\x4a\xc4\x6a\x34\xcb\x6b\x65\x4c\xad\x94\x41\xe1\x7e\xc2\x88\xfc\x85\x86\x9f\x0c\xad\xdf\x2f\x12\xca\x04\x2b\x43\xa4\x27\xa9\x23\x1e\xa3\xc2\xee\xe8\x9d\x76\xb8\x80\xad\x4d\x9e\x9e\x96\x61\xb6\x35\xb0\x51\xb0\xc4\x8c\x92\x28\xba\xc3\xf3\x30\xf9\x6f\x64\x86\xfc\xee\x05\xc4\xcf\x9d\xb7\x5d\x29\x68\x9c\x4e\x40\x11\xfa\x3c\xbc\x58\xd0\x3e\x57\xb0\xbb\x87\xcd\x38\x1d\x9b\x37\x94\x5d\xea\x8c\xbc\x20\xb9\x13\xee\x1e\x4a\xcc\x6c\x4a\x83\xcd\xcc\x85\xae\xa1\x67\xc3\x4b\x8b\xd3\xbf\xc3\xeb\xa5\xf5\xa2\x12\x2e\x8c\x43\x1b\xd6\xa0\x9c\x07\xc1\x89\x12\x9b\x1d\xce\x2d\xb9\x3d\xbc\x70\xd3\x59\x46\x3d\xc1\x3f\xff\x2e\xfe\x0b\x00\x00\xff\xff\x2e\xa2\x0d\x15\x6e\x0a\x00\x00")

func clusterBytebuildersDev_clusterinfosYamlBytes() ([]byte, error) {
	return bindataRead(
		_clusterBytebuildersDev_clusterinfosYaml,
		"cluster.bytebuilders.dev_clusterinfos.yaml",
	)
}

func clusterBytebuildersDev_clusterinfosYaml() (*asset, error) {
	bytes, err := clusterBytebuildersDev_clusterinfosYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cluster.bytebuilders.dev_clusterinfos.yaml", size: 2670, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clusterBytebuildersDev_clustersYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x4d\x8f\xe3\x36\x0c\xbd\xe7\x57\x10\xe8\x61\x2e\x1b\x07\x83\x16\x45\xe1\xdb\x20\x2d\x8a\x41\xbf\x06\x3b\x8b\xbd\xd3\x12\x63\x6b\x47\x96\x54\x92\x4a\x9b\x16\xfd\xef\x85\x64\x3b\xf1\x24\x8b\x69\x2f\x9b\x53\xf4\x48\x3d\xd1\xef\x91\xc4\xe4\x3e\x12\x8b\x8b\xa1\x05\x4c\x8e\xfe\x54\x0a\xe5\x24\xcd\xcb\x77\xd2\xb8\xb8\x3b\xde\x6f\x5e\x5c\xb0\x2d\xec\xb3\x68\x1c\xdf\x93\xc4\xcc\x86\xbe\xa7\x83\x0b\x4e\x5d\x0c\x9b\x91\x14\x2d\x2a\xb6\x1b\x00\xc3\x84\x05\xfc\xe0\x46\x12\xc5\x31\xb5\x10\xb2\xf7\x1b\x00\x8f\x1d\x79\x29\x39\x00\x98\x52\xf3\x92\x3b\xe2\x40\x4a\xf5\x95\x80\x23\xb5\xd0\x9d\x94\xba\xec\xbc\x25\x96\x0d\xc0\x04\x1a\x9f\x45\x89\xa5\x99\xff\x34\xeb\xac\xc6\xd2\x71\x23\x89\x4c\x21\xee\x39\xe6\x74\xbe\x70\x9b\x37\x31\xce\x35\x18\x54\xea\x23\xbb\xe5\xbc\x85\x4b\x45\x33\x80\x29\x89\x89\x96\xea\x71\x16\x61\xe2\xae\x88\x77\xa2\x3f\xad\xd1\x9f\x9d\x68\x8d\x24\x9f\x19\xfd\xa5\xf4\x0a\x8a\x0b\x7d\xf6\xc8\x67\x78\x03\x20\x26\x26\x6a\xe1\xd7\x52\x56\x42\x43\x76\x03\x70\x9c\xec\xa8\x65\x6d\x01\xad\xad\x2a\xa3\x7f\x62\x17\x94\x78\x1f\x7d\x1e\xc3\xb9\xe8\x4f\x12\xc3\x13\xea\xd0\x42\xb3\xf8\xd0\xdc\x98\x50\x73\x17\x3d\x1f\x7a\x9a\xcf\x7a\x2a\x8f\x5b\xd4\x09\x98\xc2\xc7\x7b\xf4\x69\xc0\xfb\xa9\x66\x33\xd0\x88\xed\x9c\x1f\x13\x85\x87\xa7\xc7\x8f\x5f\x3f\xbf\x82\x01\x12\xc7\x44\xac\x67\x2d\xa7\xdf\xaa\xb5\x56\x28\x80\x25\x31\xec\x92\xd6\x9e\xbb\x2b\x84\x53\x16\xd8\xd2\x53\x24\xa0\x03\x2d\x2a\x90\x9d\x6b\x80\x78\x00\x1d\x9c\x00\x53\x62\x12\x0a\x5a\x3f\xf1\x15\x31\x94\x24\x0c\x10\xbb\x4f\x64\xb4\x81\x67\xe2\x42\x03\x32\xc4\xec\x2d\x98\x18\x8e\xc4\x0a\x4c\x26\xf6\xc1\xfd\x75\xe6\x16\xd0\x58\x1f\xf5\xa8\x34\x5b\x78\xf9\x55\xd5\x03\x7a\x38\xa2\xcf\xf4\x0e\x30\x58\x18\xf1\x04\x4c\xe5\x15\xc8\x61\xc5\x57\x53\xa4\x81\x5f\x22\x13\xb8\x70\x88\x2d\x0c\xaa\x49\xda\xdd\xae\x77\xba\x8c\x94\x89\xe3\x98\x83\xd3\xd3\xce\xc4\xa0\xec\xba\xac\x91\x65\x67\xe9\x48\x7e\x27\xae\xdf\x22\x9b\xc1\x29\x19\xcd\x4c\x3b\x4c\x6e\x5b\x4b\x0f\x5a\xe7\x72\xb4\x5f\xf1\x3c\x84\x72\xf7\xaa\xd6\xc9\x4f\x51\x76\xa1\x5f\x05\x6a\xe7\xbe\xe1\x40\xe9\x61\x70\x02\x38\x5f\x9d\xbe\xe2\x22\x74\x81\x8a\x3a\xef\x7f\x78\xfe\x00\xcb\xd3\xd5\x8c\x6b\xf5\xab\xee\x97\x8b\x72\xb1\xa0\x08\xe6\xc2\x81\x78\x32\xf1\xc0\x71\xac\x9c\x14\x6c\x8a\x2e\x68\x3d\x18\xef\x28\x5c\xcb\x2f\xb9\x1b\x9d\x16\xdf\x7f\xcf\x24\x5a\xbc\x6a\x60\x8f\x21\x44\x85\x8e\x20\xa7\xd2\xbf\xb6\x81\xc7\x00\x7b\x1c\xc9\xef\x51\xe8\x8b\x1b\x50\x94\x96\x6d\x11\xf6\xff\x59\xb0\x5e\x91\xd7\xc9\x93\x6a\xab\xc0\xb2\xce\xfe\x3b\x51\x51\xb3\xbc\x4e\xfd\xfc\x28\x4e\xee\x48\x31\xc3\xfe\x48\x81\xb8\x0e\xcf\x75\xc6\x55\x63\xfc\x76\x73\xa1\x74\x49\x31\x6a\x8c\x52\xc7\x88\x82\x42\x7f\x89\x2e\x2f\xdc\xd0\x02\x1c\xe2\xec\xfc\xbc\x2a\x1b\x78\x54\x30\x91\x99\x24\xc5\x60\xcf\x13\x38\x87\xef\x64\xc5\xfb\x0e\xfe\x18\x9c\x19\x3e\xc3\xea\x64\xf1\x1f\x62\x80\x31\x4f\x3b\x01\xba\x53\xe5\x7a\x78\x7a\x9c\xfb\xaf\xb9\xb9\x7b\x88\x3c\xa2\xb6\x65\xb8\xbf\xfd\xe6\x26\x3a\xe9\x5d\x06\xbf\x9f\x77\xfd\x9b\x4e\xdc\x80\x93\x0c\x2d\x28\xe7\x69\xb5\x8a\x46\xc6\x9e\xd6\x48\xee\xce\x53\xbc\xf8\x30\xfb\x09\x7f\xff\xb3\xf9\x37\x00\x00\xff\xff\xd6\xd2\x64\xd4\x97\x07\x00\x00")

func clusterBytebuildersDev_clustersYamlBytes() ([]byte, error) {
	return bindataRead(
		_clusterBytebuildersDev_clustersYaml,
		"cluster.bytebuilders.dev_clusters.yaml",
	)
}

func clusterBytebuildersDev_clustersYaml() (*asset, error) {
	bytes, err := clusterBytebuildersDev_clustersYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cluster.bytebuilders.dev_clusters.yaml", size: 1943, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clusterBytebuildersDev_clusteruserauthsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x18\x4d\x6f\xdb\x36\xf4\xee\x5f\xf1\x80\x1d\x7a\xa9\x1d\x14\x1b\x86\xc1\xb7\x20\x0d\x8a\x20\xed\x16\x34\x69\x77\xa6\xa8\x27\x99\x8b\x44\x6a\x7c\x8f\x6e\xdc\x61\xff\x7d\x20\x29\xc9\xb2\x44\xd9\x4e\x82\x4e\x37\xf1\x7d\xf0\x7d\x7f\x50\x34\xea\x2b\x5a\x52\x46\xaf\x41\x34\x0a\x9f\x18\xb5\xff\xa3\xd5\xe3\x6f\xb4\x52\xe6\x62\xfb\x6e\xf1\xa8\x74\xbe\x86\x2b\x47\x6c\xea\xcf\x48\xc6\x59\x89\xef\xb1\x50\x5a\xb1\x32\x7a\x51\x23\x8b\x5c\xb0\x58\x2f\x00\xa4\x45\xe1\x0f\x1f\x54\x8d\xc4\xa2\x6e\xd6\xa0\x5d\x55\x2d\x00\x2a\x91\x61\x45\x1e\x07\x40\x34\xcd\xea\xd1\x65\x68\x35\x32\x86\x5b\xb4\xa8\x71\x0d\xd9\x8e\x31\x73\xaa\xca\xd1\xd2\x02\x20\x1e\xca\xca\x11\xa3\x75\x84\x56\x38\xde\xd0\xaa\x3d\x58\x0d\xb1\x57\x39\x6e\x17\xd4\xa0\xf4\x17\x94\xd6\xb8\xa6\x27\x9c\xe2\x45\xce\xad\x2c\x52\x30\x96\xc6\xaa\xee\x7f\x09\x7b\xc9\xda\x03\xdb\xea\xbc\xac\x4d\x8e\x55\x7b\x28\x9a\x86\xa4\xc9\x31\xfc\xb6\x16\x8a\x17\x7e\x21\xb4\x97\x8e\x37\x01\x52\x29\xe2\xdb\x14\xf4\xa3\x22\x0e\x18\x4d\xe5\xac\xa8\xa6\x7a\x06\x20\x29\x5d\xba\x4a\xd8\x09\x78\x01\x40\xd2\x34\xd8\xf3\x5d\x00\x6c\xa3\x23\x83\x22\xcb\xd6\x7a\xdb\x77\xa2\x6a\x36\xe2\x5d\xe4\x26\x37\x58\x8b\xa8\x27\x80\x69\x50\x5f\xde\xdd\x7c\xfd\xf9\xfe\xe0\x18\xa0\xb1\xa6\x41\xcb\xbd\x49\xe2\x37\x88\x94\xc1\x29\x40\x8e\x24\xad\x6a\x38\x84\xd0\x1b\xcf\x30\x62\x41\xee\x43\x04\x09\x78\x83\x9d\x68\x98\xb7\x32\x80\x29\x80\x37\x8a\xc0\x62\x63\x91\x50\x73\x08\x9b\x03\xc6\xe0\x91\x84\x06\x93\xfd\x85\x92\x57\x70\x8f\xd6\xb3\x01\xda\x18\x57\xe5\x20\x8d\xde\xa2\x65\xb0\x28\x4d\xa9\xd5\xf7\x9e\x37\x01\x9b\x70\x69\x25\x18\x5b\x23\xef\x3f\xa5\x19\xad\x16\x15\x6c\x45\xe5\xf0\x2d\x08\x9d\x43\x2d\x76\x60\xd1\xdf\x02\x4e\x0f\xf8\x05\x14\x5a\xc1\x27\x63\x11\x94\x2e\xcc\x1a\x36\xcc\x0d\xad\x2f\x2e\x4a\xc5\x5d\x86\x48\x53\xd7\x4e\x2b\xde\x5d\x48\xa3\xd9\xaa\xcc\xb1\xb1\x74\x91\xe3\x16\xab\x0b\x52\xe5\x52\x58\xb9\x51\x8c\x92\x9d\xc5\x0b\xd1\xa8\x65\x10\x5d\x73\x48\xb3\x3a\xff\xa9\x8b\x2f\x7a\x73\x20\x2b\xef\xbc\x7b\x89\xad\xd2\xe5\x00\x10\x62\xed\x88\x07\x7c\xb4\x81\x22\x10\x2d\x69\xd4\x62\x6f\x68\x7f\xe4\xad\xf3\xf9\xfa\xfe\xa1\x0f\xed\xe0\x8c\xb1\xf5\x83\xdd\xf7\x84\xb4\x77\x81\x37\x98\xd2\x05\xda\xe8\xc4\xc2\x9a\x3a\xf0\x44\x9d\x37\x46\x69\x0e\x3f\xb2\x52\xa8\xc7\xe6\x27\x97\xd5\x8a\xbd\xdf\xff\x76\x48\xec\x7d\xb5\x82\x2b\xa1\xb5\x61\xc8\x10\x5c\x93\x0b\xc6\x7c\x05\x37\x1a\xae\x44\x8d\xd5\x95\x20\xfc\xe1\x0e\xf0\x96\xa6\xa5\x37\xec\x79\x2e\x18\x56\xbc\x31\x72\xb4\xda\x00\xd0\x55\xa5\xfd\x97\xce\x2f\xff\x09\x29\x91\xe8\xc1\x3c\xa2\x1e\x83\x46\x6e\xbe\x13\x56\xd4\xc8\xde\x19\x85\xb1\xf0\xe1\xf6\x1a\xee\xac\xd9\xaa\x3c\x54\x82\xc3\x6f\x46\x8b\x70\x21\x91\xab\xf1\xb3\xa9\xf0\xd2\x26\xae\x3c\x8b\xf2\xfa\x29\x66\xd4\x4d\xfe\x3c\x06\x8e\x37\x9d\xc8\x27\x94\xbd\x1c\xa0\x06\x7b\xaa\x42\xa1\x8f\x70\x19\x7a\x52\x60\xe5\xdd\x29\x43\x11\xf1\x15\xb5\x54\xe3\x5a\xe2\x3f\x6f\x29\x1f\x99\xfb\x0a\xdf\x37\x89\x09\xf6\xbc\x93\x82\xf0\xdf\x9d\xc5\x2b\xa3\x0b\x55\xa6\xc0\xa7\xc8\x21\xd6\x53\x0a\xe9\xb4\x54\x09\xc3\x9d\x61\xc0\xee\x8b\x79\xf6\x6a\x2e\x41\x99\xd0\xe2\x5e\xc5\x07\xf5\x56\x59\xa3\x6b\xd4\xfc\x2a\x3e\x8c\x5a\xbc\x56\x2b\x5f\x64\x94\xc5\x19\x16\xcb\x03\x17\xcc\xa0\xf4\xc6\x9d\x81\x0f\xd4\x9d\xc1\xe8\x15\x49\xc2\x67\x8a\x46\xf7\x85\x26\x7e\x84\x70\x46\xf9\x79\xc5\xe3\x5c\x30\x93\xa7\x49\x31\xa4\x90\x9b\x84\x0c\x91\x22\x33\xa6\x42\x31\xce\xb5\x68\xb4\x2b\x1f\xff\x85\x4f\x4a\x7c\x3f\x29\x96\x30\xce\xf1\xab\x14\x8d\x8f\x4a\x16\x4a\x13\xdc\x5d\x7f\x5a\xa2\xf6\x03\x57\x0e\xbe\xf4\x86\x9e\x93\x4a\xcc\xf6\x72\x90\x7e\x36\x28\x54\x85\x21\xeb\x1f\x3e\xde\x4f\x53\xbc\x30\xb6\x16\x1c\xc7\xce\xe7\x54\xae\x78\xc3\x2d\xee\xce\x56\xab\xc5\x3d\xa1\x4e\x2f\x7c\x42\xaf\x47\xdc\xfd\x30\x6d\x42\x05\x7c\x66\xf5\x6e\xa9\xbe\xdc\xbc\x7f\x16\x19\x3e\x35\xca\xee\xa6\x24\x9d\xf0\x4a\xf3\xaf\xbf\xcc\x30\xf4\x93\x5b\x39\x69\x6f\x85\xb1\xdf\x84\xcd\xef\x91\xfc\x80\xf9\x7b\x32\x5f\x8e\xc5\xaa\xaa\x1b\xb4\x64\xb4\xe0\x04\xe1\x81\x2b\x6f\xf6\x98\x7e\xbe\xf2\x8d\xc4\x4f\xe3\x3e\x9f\xfc\xc4\x29\x24\x2f\x05\x4d\xfd\x72\xc4\x1a\x83\xbb\x3f\xf8\x9d\x25\xd1\x2b\xe6\x24\x88\xf8\x9d\x1c\x65\xfc\x63\x33\x64\x39\x15\x45\x31\xd6\xc9\x7e\x74\xb4\x98\x44\xa0\xb0\x56\xec\xe6\xe5\xf7\x6b\xcd\xf5\x13\xdb\x44\x3e\x88\x3c\x0f\x8b\xa2\xa8\xee\x8e\x36\xc5\x03\x55\x03\xaf\xaf\x61\x72\xad\x05\x3d\xb6\x5b\x44\xf8\x27\xe3\xbb\x2b\x9b\xcc\x15\x20\x85\x86\x12\x35\x5a\x91\x08\xfa\x23\x0a\x9f\x6c\x1e\xf3\x4a\xcf\xfb\xa4\xb7\xc1\x3e\xcf\xf7\xba\x87\xf1\xd5\x07\xf9\x74\xcd\x81\x76\x34\x19\x98\x33\x0f\xa1\x35\x17\x4c\xc9\x3a\xad\xea\xc6\x58\xc6\xfc\x32\xd1\x77\x5f\x9e\x5e\x7e\x54\x9a\x9b\x73\x0e\xec\x70\xdb\x23\xee\xb5\x2f\x94\x57\xdc\xb3\x68\xc7\x8b\x7e\x02\x93\xfd\xd2\x9a\x12\x25\xe9\x95\x46\x10\x7d\x33\x36\x51\xa6\x46\x83\x71\x44\xeb\x52\xa3\x23\x0b\x57\x67\x82\x94\x1c\x8f\x8b\x6c\x52\xfe\x3f\x73\x4c\x3c\x26\xf0\xec\x70\x7b\x84\xc8\x62\x61\x91\x36\x33\x2b\xc0\x51\xc2\x72\xb2\x9a\x27\x8c\x73\xb0\x35\x5c\xfe\x79\xff\xa2\xad\x81\x4e\x17\xdb\x24\x1d\x9f\xb1\xd8\x04\xcd\x3b\xe7\x65\x28\xac\x5f\x32\xc3\x99\x97\x78\xe2\xba\x91\xa3\x12\x9e\x7c\x89\xeb\x7c\xf2\xa5\x3a\xdb\xcb\xf3\xa8\xeb\x14\x27\xb4\xff\xd2\x35\x94\x71\x83\xf9\x9f\xa3\x37\x3d\x40\x2e\x07\x6d\x7f\x04\xd8\x57\x9f\x11\x20\x9a\xf2\xac\x45\x99\x05\x3b\x3a\x77\x55\x36\x59\x18\xdd\xf3\x0f\xb1\xfa\x9f\x8e\xfd\x3f\x26\x04\x9d\x91\x6b\x43\xe1\x19\xc9\x8f\x8d\xe5\x1e\xda\xdd\x30\xbb\x48\x86\xe7\xab\xf8\x7e\xb2\x82\x1b\x06\x69\xac\x45\x6a\x8c\xce\xfb\x27\xa8\x0e\xfe\x86\x06\x9c\xdf\x26\x38\x7e\xdb\x28\xb9\xf1\x02\xb5\x2f\x20\x60\x34\xd4\x2e\xbe\x8a\x41\xb6\x0b\xcc\x2e\xef\x6e\xda\x17\x98\xf9\xe9\xef\xb9\x91\x99\xf4\xc5\xe4\x30\x1a\x62\x0d\x6c\x5d\x6c\xb3\xc4\xc6\x8a\x12\x87\x27\x2e\xeb\xdf\xb1\x3a\x4f\xb4\x1e\x85\x7f\xfe\x5d\xfc\x17\x00\x00\xff\xff\x68\x18\x48\x54\x68\x16\x00\x00")

func clusterBytebuildersDev_clusteruserauthsYamlBytes() ([]byte, error) {
	return bindataRead(
		_clusterBytebuildersDev_clusteruserauthsYaml,
		"cluster.bytebuilders.dev_clusteruserauths.yaml",
	)
}

func clusterBytebuildersDev_clusteruserauthsYaml() (*asset, error) {
	bytes, err := clusterBytebuildersDev_clusteruserauthsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cluster.bytebuilders.dev_clusteruserauths.yaml", size: 5736, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"cluster.bytebuilders.dev_clusterauthinfotemplates.yaml": clusterBytebuildersDev_clusterauthinfotemplatesYaml,
	"cluster.bytebuilders.dev_clusterinfos.yaml":             clusterBytebuildersDev_clusterinfosYaml,
	"cluster.bytebuilders.dev_clusters.yaml":                 clusterBytebuildersDev_clustersYaml,
	"cluster.bytebuilders.dev_clusteruserauths.yaml":         clusterBytebuildersDev_clusteruserauthsYaml,
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
	"cluster.bytebuilders.dev_clusterauthinfotemplates.yaml": {clusterBytebuildersDev_clusterauthinfotemplatesYaml, map[string]*bintree{}},
	"cluster.bytebuilders.dev_clusterinfos.yaml":             {clusterBytebuildersDev_clusterinfosYaml, map[string]*bintree{}},
	"cluster.bytebuilders.dev_clusters.yaml":                 {clusterBytebuildersDev_clustersYaml, map[string]*bintree{}},
	"cluster.bytebuilders.dev_clusteruserauths.yaml":         {clusterBytebuildersDev_clusteruserauthsYaml, map[string]*bintree{}},
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
