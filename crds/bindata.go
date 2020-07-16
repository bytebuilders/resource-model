// Package crds Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// cluster.bytebuilders.dev_cloudcredentials.yaml
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

var _clusterBytebuildersDev_cloudcredentialsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x51\x8f\xdb\xb6\x0f\x7f\xcf\xa7\x20\xf0\x7f\xe8\x4b\x93\xa0\xf8\x0f\xc3\x90\xb7\xc3\x75\x18\x0e\xdd\xd0\xa2\x77\xe8\xbb\x2c\x31\x8e\x76\xb2\xa4\x91\x54\xae\xe9\xb0\xef\x3e\x48\xb6\x13\x5f\xe2\xa4\x3d\xdf\x15\xdb\xfc\x66\x91\xa2\xc8\xdf\x8f\xa4\x45\xab\x68\x3f\x21\xb1\x0d\x7e\x05\x2a\x5a\xfc\x2c\xe8\xf3\x1b\x2f\xee\x7f\xe2\x85\x0d\xcb\xed\x9b\xd9\xbd\xf5\x66\x05\xd7\x89\x25\x34\x1f\x91\x43\x22\x8d\x6f\x71\x6d\xbd\x15\x1b\xfc\xac\x41\x51\x46\x89\x5a\xcd\x00\x34\xa1\xca\x8b\x77\xb6\x41\x16\xd5\xc4\x15\xf8\xe4\xdc\x0c\xc0\xa9\x0a\x1d\x67\x1d\x00\x15\xe3\xe2\x3e\x55\x48\x1e\x05\xcb\x29\x5e\x35\xb8\x82\x6a\x27\x58\x25\xeb\x0c\x12\xcf\x00\xda\x45\xed\x42\x32\x9a\xd0\xa0\x17\xab\x1c\x2f\xb4\x4b\x2c\x48\x8b\xa1\xf6\xc2\xe0\x76\xc6\x11\x75\x3e\xa0\xa6\x90\x62\xde\x78\x46\xaf\xb5\xdc\xf9\xa2\x95\x60\x1d\xc8\xf6\xef\x73\x38\x78\xd6\x2d\x50\x17\xf3\xbc\x09\x06\x5d\xb7\xa8\x62\x64\x1d\x0c\x96\xd7\x0e\xa1\xec\xe9\xf5\xde\xd3\x22\x71\x96\xe5\xdd\x98\xf4\x57\xcb\x52\x34\xa2\x4b\xa4\xdc\x69\x9c\x45\xc8\xd6\xd7\xc9\x29\x3a\x11\xcf\x00\x58\x87\x88\xd9\x6e\x09\x73\x06\xb0\x6d\x89\x2c\x81\xcc\x3b\xf4\xb6\x6f\x94\x8b\x1b\xf5\xa6\xb5\xa6\x37\xd8\xa8\x36\x4e\x80\x10\xd1\x5f\x7d\xb8\xf9\xf4\xff\xdb\x47\xcb\x00\x91\x42\x44\x92\x3d\x24\xed\x33\xc8\x94\xc1\x2a\x80\x41\xd6\x64\xa3\x94\x14\x7a\x95\x0d\xb6\x5a\x60\x72\x8a\x20\x83\x6c\xb0\x77\x0d\x4d\xe7\x03\x84\x35\xc8\xc6\x32\x10\x46\x42\x46\x2f\x25\x6d\x1e\x19\x86\xac\xa4\x3c\x84\xea\x77\xd4\xb2\x80\x5b\xa4\x6c\x06\x78\x13\x92\x33\xa0\x83\xdf\x22\x09\x10\xea\x50\x7b\xfb\x65\x6f\x9b\x41\x42\x39\xd4\x29\xc1\x0e\xe4\xc3\x63\xbd\x20\x79\xe5\x60\xab\x5c\xc2\xd7\xa0\xbc\x81\x46\xed\x80\x30\x9f\x02\xc9\x0f\xec\x15\x15\x5e\xc0\x6f\x81\x10\xac\x5f\x87\x15\x6c\x44\x22\xaf\x96\xcb\xda\x4a\x5f\x21\x3a\x34\x4d\xf2\x56\x76\x4b\x1d\xbc\x90\xad\x92\x04\xe2\xa5\xc1\x2d\xba\x25\xdb\x7a\xae\x48\x6f\xac\xa0\x96\x44\xb8\x54\xd1\xce\x8b\xeb\x5e\x4a\x99\x35\xe6\x7f\x7d\x7e\xf1\xab\x47\xbe\xca\x2e\xd3\xcb\x42\xd6\xd7\x03\x41\xc9\xb5\x0b\x0c\xe4\x6c\x03\xcb\xa0\xba\xad\x6d\x14\x07\xa0\xf3\x52\x46\xe7\xe3\xcf\xb7\x77\xfb\xd4\x2e\x64\x1c\xa3\x5f\x70\x3f\x6c\xe4\x03\x05\x19\x30\xeb\xd7\x48\x2d\x89\x6b\x0a\x4d\xb1\x89\xde\xc4\x60\xbd\x94\x17\xed\x2c\xfa\x63\xf8\x39\x55\x8d\x95\xcc\xfb\x1f\x09\x59\x32\x57\x0b\xb8\x56\xde\x07\x81\x0a\x21\x45\xa3\x04\xcd\x02\x6e\x3c\x5c\xab\x06\xdd\xb5\x62\xfc\xee\x04\x64\xa4\x79\x9e\x81\xfd\x36\x0a\x86\x1d\xef\x58\xb9\x45\x6d\x20\xe8\xbb\xd2\xe1\x19\xaf\xaf\xfc\xa8\x07\x3e\x74\x88\x63\xe1\xa5\x8d\x65\xb3\xd6\xc8\xfc\x0e\x77\x37\x6f\xc7\xc4\x67\xa3\xd9\x3b\x8a\x9a\x50\xae\x7a\x2b\x13\x6c\x64\x52\x2d\xa1\x39\xdd\x3a\x1f\x7a\x37\x22\x3d\x3a\xfb\x44\xe3\x0c\xb6\x25\xec\x2f\x89\x70\x3a\x6a\x6d\x96\x4e\x84\xac\xdd\x7c\x5b\x9c\x9f\x86\x79\xaa\xf6\xb5\x3b\xd1\x07\x41\xaf\x26\x06\x70\x89\xaf\x1e\x97\xb3\xa2\x36\xea\x31\x2e\x1f\xc5\x34\xa2\xd0\x7b\xfc\x14\x96\x8d\xad\xad\x28\xf7\x5e\xa3\xf2\xd3\xc9\x96\x70\x8f\xfe\x85\x81\x2a\x36\x9f\x12\x4a\xad\x9f\x91\xae\x91\x42\x36\x39\xb9\xc4\x69\x6b\x35\x5e\x69\x1d\x92\x9f\x92\xb1\x97\x80\xd8\xbb\x36\x5a\xdf\xc3\x83\x9f\x84\x56\x08\xb5\x7b\x06\x60\x6d\xdf\xb9\x9b\x48\xfc\x3f\xdd\x1f\xf0\x73\xb4\x74\xa6\x15\xaf\x03\x35\x4a\x56\xf9\x46\xf3\xe3\x0f\x17\x8c\xe7\x1b\x4f\x5d\x2e\x88\xc7\x0f\xe1\x9a\x90\x37\xd3\xc1\x29\x57\xd0\x51\xd8\x01\xac\x60\x73\x46\xf4\x55\xbb\xbd\x82\x22\x52\xa7\x1f\x83\xaf\x7f\x66\xee\x46\x6b\x72\x7a\x53\xbb\x90\x9f\xce\xfa\x60\x9e\x91\x9f\xff\x8a\x96\x54\x66\x85\x33\x1b\x46\xcf\x0f\x0f\x1e\x69\xac\x28\x2e\x27\xe5\xa5\x84\x8c\x4a\xdf\xa3\x3c\xa3\xd0\xa3\x9d\x76\x6b\x79\x5e\x53\xbd\x98\x8c\xc5\xa5\x27\xb5\xca\x0b\x24\x45\x0a\x5b\x6b\x90\x9e\x44\x14\x6b\xe5\xf0\x41\xed\xa6\xe3\x1a\xa8\x56\xde\x7e\x29\xf3\xd9\xb4\x0b\xca\x77\xc8\xf0\xa1\x53\x2f\x51\x00\xdb\xe4\x84\xfe\xd3\x45\x3c\x6e\xaa\xfd\x0d\x70\xb4\xd4\x15\xef\xd1\x6a\x9f\x5e\xdf\x34\xd3\x88\x92\xc4\xdf\x3a\xd5\x84\x2a\x7f\xfe\xd1\xfc\x82\x1e\xe9\x4c\x1e\x3d\x9a\x61\xdf\x9f\x6c\xc8\x03\x6d\x9e\x29\x9b\xc0\x65\xe2\x47\x2f\x50\x1f\xa4\xfd\x09\x23\x0c\xac\x03\xf5\x7f\x1a\xda\x51\x77\x01\x37\x02\x3a\x10\x21\xc7\xe0\xcd\xfe\x6f\x41\x2f\x7f\xc5\x03\xcb\xaf\x47\x2c\x3e\x6c\xac\xde\x64\x87\xba\x61\x15\x82\x87\x26\xb5\x3f\x30\xa0\xda\x15\x63\x57\x1f\x6e\xba\x61\x79\xf1\x62\x2d\x72\x94\x8b\x93\xc5\x16\x88\x15\x08\xa5\x96\x77\x96\x40\xaa\xc6\xe1\x4a\xaa\xf6\xbf\x1c\x7a\x26\x3a\x46\xe1\xcf\xbf\x66\x7f\x07\x00\x00\xff\xff\x23\x25\xaf\x6e\x13\x14\x00\x00")

func clusterBytebuildersDev_cloudcredentialsYamlBytes() ([]byte, error) {
	return bindataRead(
		_clusterBytebuildersDev_cloudcredentialsYaml,
		"cluster.bytebuilders.dev_cloudcredentials.yaml",
	)
}

func clusterBytebuildersDev_cloudcredentialsYaml() (*asset, error) {
	bytes, err := clusterBytebuildersDev_cloudcredentialsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cluster.bytebuilders.dev_cloudcredentials.yaml", size: 5139, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clusterBytebuildersDev_clusterauthinfotemplatesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\xcd\x6e\x1b\x37\x10\xbe\xeb\x29\x06\xe8\x21\x97\x48\x42\xd0\xa2\x28\x74\x33\x9c\x20\x10\x92\xa0\x46\xec\xe4\xce\x25\x47\x2b\xd6\xbb\x24\xcb\x19\xca\x51\x8b\xbe\x7b\x31\xdc\x5d\x69\xbd\x3f\xb2\x5a\x34\xe5\x8d\xe4\x70\x38\xdf\x37\x7f\xa4\x0a\xf6\x2b\x46\xb2\xde\x6d\x40\x05\x8b\xdf\x18\x9d\xcc\x68\xf5\xf8\x0b\xad\xac\x5f\x1f\xde\x2c\x1e\xad\x33\x1b\xb8\x4d\xc4\xbe\xfe\x8c\xe4\x53\xd4\xf8\x16\x77\xd6\x59\xb6\xde\x2d\x6a\x64\x65\x14\xab\xcd\x02\x40\x47\x54\xb2\xf8\x60\x6b\x24\x56\x75\xd8\x80\x4b\x55\xb5\x00\xa8\x54\x81\x15\x89\x0c\x80\x0a\x61\xf5\x98\x0a\x8c\x0e\x19\xf3\x2d\x4e\xd5\xb8\x81\xe2\xc8\x58\x24\x5b\x19\x8c\xb4\x00\x68\x16\x75\x95\x88\x31\xaa\xc4\x7b\xeb\x76\x9e\xb1\x0e\x95\x92\x63\xed\xc6\xaa\x7f\x6a\x65\xf0\xb0\xa0\x80\x5a\x2e\x2a\xa3\x4f\xe1\xa4\x60\x2c\xd7\xdc\xd0\xda\xa4\x15\x63\xe9\xa3\xed\xe6\x4b\x38\x5b\xd8\x2e\xc4\x16\xfb\xb2\xf6\x06\xab\x76\x51\x85\x40\xda\x1b\xcc\xd3\x96\xa9\xe6\xc2\x9b\xc4\xfb\xad\xdb\xf9\x87\xd6\xe2\x2c\x51\x59\xe2\x0f\x97\xa4\x3e\x5a\xe2\x2c\x19\xaa\x14\x55\x35\x8f\x3f\x0b\x91\x75\x65\xaa\x54\x9c\x15\x5b\x00\x90\xf6\x01\x4f\xf7\x2d\x00\x0e\x8d\xc3\x33\xd0\x65\xcb\xf2\xe1\x8d\xaa\xc2\x5e\xbd\x69\xb4\xea\x3d\xd6\xaa\xe1\x01\xc0\x07\x74\x37\x77\xdb\xaf\x3f\xde\x3f\x5b\x06\x08\xd1\x07\x8c\x7c\xa2\xac\x19\xbd\x88\xea\xad\x02\x18\x24\x1d\x6d\xe0\x1c\x6a\xaf\x44\x61\x23\x05\x46\x42\x09\x09\x78\x8f\x9d\x69\x68\x5a\x1b\xc0\xef\x80\xf7\x96\x20\x62\x88\x48\xe8\x38\x87\xd7\x33\xc5\x20\x42\xca\x81\x2f\x7e\x43\xcd\x2b\xb8\xc7\x28\x6a\x80\xf6\x3e\x55\x06\xb4\x77\x07\x8c\x0c\x11\xb5\x2f\x9d\xfd\xe3\xa4\x9b\x80\x7d\xbe\x34\xd3\xc9\x03\x9d\xd6\x31\x46\xa7\x2a\x38\xa8\x2a\xe1\x6b\x50\xce\x40\xad\x8e\x10\x51\x6e\x81\xe4\x7a\xfa\xb2\x08\xad\xe0\x93\x8f\x08\xc2\xfe\x06\xf6\xcc\x81\x36\xeb\x75\x69\xb9\xcb\x24\xed\xeb\x3a\x39\xcb\xc7\xb5\xf6\x8e\xa3\x2d\x12\xfb\x48\x6b\x83\x07\xac\xd6\x64\xcb\xa5\x8a\x7a\x6f\x19\x35\xa7\x88\x6b\x15\xec\x32\x9b\xee\x38\xa7\x63\x6d\x7e\xe8\xe2\x8f\x5e\x3d\xb3\x95\x8f\xe2\x5e\xe2\x68\x5d\xd9\xdb\xc8\xb1\x78\xc1\x03\x12\x85\x60\x09\x54\x7b\xb4\x41\x71\x26\x5a\x96\x84\x9d\xcf\xef\xee\x1f\x4e\xa1\x9f\x9d\x31\x64\x3f\xf3\x7e\x3e\x48\x67\x17\x08\x61\xd6\xed\x30\x36\x4e\xdc\x45\x5f\x67\x9d\xe8\x4c\xf0\xd6\x71\x9e\xe8\xca\xa2\x1b\xd2\x4f\xa9\xa8\x2d\x8b\xdf\x7f\x4f\x48\x2c\xbe\x5a\xc1\xad\x72\xce\x33\x14\x08\x29\x18\xc5\x68\x56\xb0\x75\x70\xab\x6a\xac\x6e\x15\xe1\x77\x77\x80\x30\x4d\x4b\x21\xf6\x3a\x17\xf4\x2b\xe3\x50\xb8\x61\xad\xb7\xd1\x55\xad\xf3\x98\xce\x2f\x19\x92\xe5\x77\xd1\x1f\xac\xc1\x38\xdc\x1b\xf8\xf9\xa6\x27\x9a\xef\xb0\x3b\x8b\xe2\x75\x9d\xeb\x79\x56\x25\x10\x75\x4e\x2c\xa9\x3a\xa5\x1d\xe6\x97\x8c\x9d\x8f\xd9\x5b\xe7\xaa\x78\x2a\xac\x23\xe9\x79\xc3\x65\x68\xef\x76\xb6\x9c\xda\x01\x50\xc6\xe4\xa6\xa2\xaa\xbb\x8b\x3a\x60\x9e\xf4\xa1\xc0\x88\xe8\x6e\xe4\xc2\x77\xe1\xe0\x8c\x66\x89\x48\x1b\xd1\x8c\x8f\x36\xb5\x74\xb4\x7c\xc1\x0c\x2d\x08\x77\xc2\x3d\x8a\x9f\x7c\xb4\x7c\x7c\x3b\x0a\x18\x18\xfa\xf4\x76\xe6\x98\x50\xcb\xca\x3a\x82\xbb\x77\x9f\x96\xe8\xa4\x2f\x99\xfe\x25\x13\x68\x55\xa7\xa0\x2f\x47\x63\x9f\xee\x7c\xac\x15\x37\x3d\x7a\x06\xe2\x24\x61\x4d\x76\xf7\x2c\xbe\x06\xdf\xd4\x99\x69\x70\x92\x5e\xb9\xae\x4c\x41\x6b\x2f\xcf\xc8\x60\x67\x2b\xcc\x51\xfc\xf0\xf1\xfe\xbf\x86\xf7\x01\xaf\x72\x5b\x5f\xf6\x05\x38\x27\xe3\x27\x70\x3d\xe2\xf1\x3b\xa1\xb1\x75\xc0\x48\xde\x29\x9e\xc8\x8c\x67\x58\xb6\x67\x49\x69\x22\x52\x19\x12\x49\xc3\xac\x51\xda\xaa\xd2\xbc\x54\x13\x51\x74\xdd\xdd\xef\xe5\xe1\x36\x91\xf7\x73\x16\x34\xf2\x9d\x1d\x65\x33\x63\xdf\x57\x39\x36\xc5\x32\xd6\x93\xb5\xe5\x62\xf6\x37\x9b\x2a\x46\x75\x9c\xb7\xff\x0b\x61\x7c\xf7\x8d\xe3\x44\x40\x5c\x5b\xe0\x9e\x41\xcd\xba\xbe\xe6\xf6\x5c\x2b\x7a\x6c\x9f\x4a\x79\x4e\x5e\xaa\x2d\xfb\x22\xed\x40\x2b\x07\x25\x3a\x8c\xd3\x99\x3e\x0b\xf8\xc5\x52\x3a\x0f\x7a\xde\x27\x27\x0e\xce\x81\x7e\xc6\x9e\x7b\xb4\x84\xe8\xf8\x2d\x07\x6d\xaf\xe9\xd1\x69\x72\x68\xcd\x05\xd3\x64\x61\xf5\x4f\x0e\xe3\xf6\xed\x18\x6d\x97\x19\xd6\xf1\xcf\x3f\xcd\x68\x94\xa7\x5f\x99\x5f\xca\xfd\x11\x14\xd1\x93\x8f\x13\x75\xff\x19\x05\x77\xad\x58\x17\x8d\xdd\xb1\x0c\xaa\x50\x64\xf5\xb0\xe5\xb2\x9f\xa2\xfc\xca\x56\x7b\xc1\x73\xec\x1f\xd1\xbd\x60\xed\x83\xc8\x74\xa6\x16\xa8\xa2\x3c\xd5\xf2\x9a\x98\x3b\x32\x74\x60\xd6\x54\x73\xff\x17\x86\x26\x3b\x41\xea\x25\xf9\xb6\xd0\xbc\x80\xed\x4b\x57\x8f\x86\xf5\xe9\x7f\xf6\xc4\xf4\x83\x61\xd9\xc5\xe8\x60\x35\x59\x73\xd5\x6b\x91\x15\x27\xba\xf6\xbd\xe8\x0b\x92\x17\xb9\x79\xdf\x54\x87\xd1\xdf\x6c\xc4\xdd\xaf\xa3\x03\x1d\x8b\xb5\xa7\xfc\x97\x92\xbe\x5a\x9e\x77\xbb\x1b\x66\x5f\x8e\xf9\x0f\xd7\x7c\x22\x56\xb0\x65\xd0\x3e\x46\xa4\xe0\x9d\x39\xfd\xc3\xba\xfd\x57\xd4\xd3\xfc\x7a\x42\xe3\xd3\xde\xea\xbd\x18\xd4\x7e\x03\xc0\x3b\xa8\x53\xf3\x35\x84\xe2\x98\x95\xdd\xdc\x6d\xdb\x6f\xc8\x7c\x7b\xfc\xa7\x45\x60\xd2\x17\xa3\xc5\x86\x88\x0d\x70\x4c\x4d\x19\x26\xf6\x51\x95\xd8\x5f\x49\xc5\xe9\x33\xd7\x79\xa2\xf5\x28\xfc\xf9\xd7\xe2\xef\x00\x00\x00\xff\xff\xb4\xb0\x0d\xe8\x95\x11\x00\x00")

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

	info := bindataFileInfo{name: "cluster.bytebuilders.dev_clusterauthinfotemplates.yaml", size: 4501, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clusterBytebuildersDev_clusterinfosYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4d\x8f\xdb\x36\x10\xbd\xfb\x57\x0c\xd0\x43\x2e\xb1\x8c\x6d\x8a\xa2\xf0\x6d\xb1\x29\x0a\x23\x6d\xba\xc8\x06\xb9\x53\xe4\x58\x62\x97\x22\xd9\x99\xa1\x93\x6d\xd1\xff\x5e\x90\x94\x6c\xf9\x23\x85\x5d\xb4\xba\x71\x3e\x1e\x87\x6f\x1e\x87\x52\xd1\x7e\x42\x62\x1b\xfc\x1a\x54\xb4\xf8\x45\xd0\xe7\x15\x37\xcf\x3f\x70\x63\xc3\x6a\x77\xb7\x78\xb6\xde\xac\xe1\x21\xb1\x84\xe1\x03\x72\x48\xa4\xf1\x2d\x6e\xad\xb7\x62\x83\x5f\x0c\x28\xca\x28\x51\xeb\x05\x80\x26\x54\xd9\xf8\xd1\x0e\xc8\xa2\x86\xb8\x06\x9f\x9c\x5b\x00\x38\xd5\xa2\xe3\x1c\x03\xa0\x62\x6c\x9e\x53\x8b\xe4\x51\xb0\xec\xe2\xd5\x80\x6b\x68\x5f\x04\xdb\x64\x9d\x41\xe2\x05\x40\x35\x6a\x97\x58\x90\xac\xdf\x06\x6e\xc6\x45\x33\x8f\x6c\x0c\xee\x16\x1c\x51\x67\xf0\x8e\x42\x8a\xfb\xa4\xf3\xb8\x8a\x3a\xd6\xa1\x95\x60\x17\xc8\x4e\xeb\x25\x1c\xaa\x1a\x0d\x34\x9e\x77\x39\x04\x83\x6e\x34\xaa\x18\x59\x07\x83\x65\x39\xb2\x53\x37\xdc\xf8\x6d\x28\x56\x67\x59\xde\x9d\x7a\x7e\xb6\x2c\xc5\x1b\x5d\x22\xe5\x8e\xcf\x56\x1c\x6c\x7d\x97\x9c\xa2\x23\xd7\x02\x80\x75\x88\xb8\xc7\x5a\x00\xec\x6a\xd3\x4a\xe1\xcb\x91\xa9\xdd\x9d\x72\xb1\x57\x77\x15\x49\xf7\x38\xa8\x7a\x2e\x80\x10\xd1\xdf\x3f\x6e\x3e\xbd\x79\x3a\x32\x03\x44\x0a\x11\x49\xf6\x14\xd4\x6f\xa6\x8a\x99\x15\xc0\x20\x6b\xb2\x51\x8a\x5c\x5e\x65\xc0\x1a\x05\x26\xcb\x01\x19\xa4\xc7\xa9\x34\x34\x63\x0d\x10\xb6\x20\xbd\x65\x20\x8c\x84\x8c\x5e\x8a\x44\x8e\x80\x21\x07\x29\x0f\xa1\xfd\x0d\xb5\x34\xf0\x84\x94\x61\x80\xfb\x90\x9c\x01\x1d\xfc\x0e\x49\x80\x50\x87\xce\xdb\x3f\xf6\xd8\x0c\x12\xca\xa6\x4e\x09\x8e\xe4\x1e\x3e\xeb\x05\xc9\x2b\x07\x3b\xe5\x12\xbe\x06\xe5\x0d\x0c\xea\x05\x08\xf3\x2e\x90\xfc\x0c\xaf\x84\x70\x03\xbf\x04\x42\xc8\xac\xaf\xa1\x17\x89\xbc\x5e\xad\x3a\x2b\xd3\x6d\xd0\x61\x18\x92\xb7\xf2\xb2\xd2\xc1\x0b\xd9\x36\x49\x20\x5e\x19\xdc\xa1\x5b\xb1\xed\x96\x8a\x74\x6f\x05\xb5\x24\xc2\x95\x8a\x76\x59\x4a\xf7\x52\xae\xd4\x60\xbe\x99\xf4\xc4\xaf\x8e\x6a\x95\x97\xdc\x5e\x16\xb2\xbe\x9b\x39\x8a\xb6\xfe\xa1\x03\x59\x61\x60\x19\xd4\x98\x5a\x4f\x71\x20\x3a\x9b\x32\x3b\x1f\x7e\x7c\xfa\xb8\x97\x72\x69\xc6\x29\xfb\x85\xf7\x43\x22\x1f\x5a\x90\x09\xb3\x7e\x8b\x54\x9b\xb8\xa5\x30\x14\x4c\xf4\x26\x06\xeb\xa5\x2c\xb4\xb3\xe8\x4f\xe9\xe7\xd4\x0e\x56\x72\xdf\x7f\x4f\xc8\x92\x7b\xd5\xc0\x83\xf2\x3e\x08\xb4\x08\x29\x1a\x25\x68\x1a\xd8\x78\x78\x50\x03\xba\x07\xc5\xf8\xbf\x37\x20\x33\xcd\xcb\x4c\xec\x75\x2d\x98\x4f\xb7\xd3\xe0\xca\xda\xcc\x31\x4d\xa1\xc3\x77\xf9\x7e\xe5\xaf\x8c\x4a\x34\xf7\x72\xea\x00\xd8\x06\x1a\x94\xac\xb3\x7e\xbf\xff\xee\xcc\x5b\xf7\xce\xda\xee\xca\x28\x98\x7f\x53\x53\xce\x31\xbf\x72\xbc\x92\xf4\xa5\x5e\x93\xcd\xdb\x9b\xd2\xca\xfc\x57\xee\xdd\x7e\x64\x5e\x1c\x19\xd7\x81\xbc\x0f\x06\x1f\x42\xba\x54\xf8\x8c\x8c\x37\xdf\xde\x44\x86\x0b\xba\x4c\x9a\x9b\x0a\x2a\xa3\xf4\x96\x84\xf0\xd9\x23\x5d\xa2\xee\xdf\x77\xb1\x40\xbe\xbf\xb5\x90\x48\x21\xcb\xf1\xd6\x9c\x9d\x35\x48\x37\x25\x25\x6b\x6e\x88\xcf\xb7\xdf\x12\x9e\xa4\xd4\x37\xeb\xc4\x34\x52\x79\x62\x4d\xd6\x5c\x75\xf7\x44\x49\xe2\x6b\x6f\x5f\x68\x39\xcf\x37\xf3\x13\x7a\xa4\xaf\xa8\xe4\x68\xd6\xfe\x7a\x96\x90\x07\x6f\x9e\x7d\x43\xe0\xf2\x32\xa1\x17\xe8\x0e\xde\x69\x87\x33\xd8\xa2\x8c\xe9\x45\xac\x23\xb9\x81\x8d\x80\x0e\x44\xc8\x31\x78\xb3\x7f\xd5\x26\xff\x2b\x9e\x21\xbf\xbe\x80\xf8\xb9\xb7\xba\xcf\x05\x8d\x43\x15\x82\x87\x21\xd5\x87\x16\xda\x97\x02\x76\xff\xb8\x19\x87\x7a\xf3\x1f\x6a\x35\xf6\x8a\x2f\xe8\xf4\x88\xbb\xc7\x1c\x33\x7b\x5c\x40\x27\xa2\x4c\x57\xed\x59\xfd\x41\xc0\xe9\x97\xe7\x7a\x69\x5d\x54\xc2\x99\xb1\xb6\x61\x0d\x42\xa9\x0a\x8e\x25\x90\xea\x70\x6e\x49\xed\xfe\x61\x9e\xce\x32\xea\x09\xfe\xfc\x6b\xf1\x77\x00\x00\x00\xff\xff\x85\x33\xfe\xf8\x25\x0b\x00\x00")

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

	info := bindataFileInfo{name: "cluster.bytebuilders.dev_clusterinfos.yaml", size: 2853, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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

var _clusterBytebuildersDev_clusteruserauthsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x5f\x8f\xdb\x36\x0c\x7f\xcf\xa7\x20\xb0\x87\xbe\x34\x39\x14\x1b\x86\x21\x6f\x87\xeb\xa1\x38\xb4\xdd\x0e\xbd\x6b\xdf\x65\x99\x76\xb4\xd8\x92\x47\xd2\xb9\xcb\x86\x7d\xf7\x41\x92\x9d\x38\xfe\x77\xb9\x14\x9d\xdf\x2c\x51\x14\xf9\x23\xf9\x13\x25\x55\x99\x6f\x48\x6c\x9c\x5d\x83\xaa\x0c\x3e\x0b\x5a\xff\xc7\xab\xed\x6f\xbc\x32\xee\x6a\xf7\x6e\xb1\x35\x36\x5d\xc3\x4d\xcd\xe2\xca\x2f\xc8\xae\x26\x8d\xef\x31\x33\xd6\x88\x71\x76\x51\xa2\xa8\x54\x89\x5a\x2f\x00\x34\xa1\xf2\x83\x8f\xa6\x44\x16\x55\x56\x6b\xb0\x75\x51\x2c\x00\x0a\x95\x60\xc1\x5e\x06\x40\x55\xd5\x6a\x5b\x27\x48\x16\x05\xc3\x2e\x56\x95\xb8\x86\x64\x2f\x98\xd4\xa6\x48\x91\x78\x01\x10\x07\x75\x51\xb3\x20\xd5\x8c\xa4\x6a\xd9\xf0\xaa\x19\x58\x75\xa5\x57\x29\xee\x16\x5c\xa1\xf6\x1b\xe4\xe4\xea\xea\xb0\x70\x28\x17\x35\x37\xb6\x68\x25\x98\x3b\x32\xed\xff\x12\x8e\x96\x35\x03\xd4\xf8\xbc\x2c\x5d\x8a\x45\x33\xa8\xaa\x8a\xb5\x4b\x31\xfc\x36\x08\xc5\x0d\xbf\x32\xd2\x75\x2d\x9b\x30\x53\x18\x96\x8f\x63\xb3\x9f\x0c\x4b\x90\xa8\x8a\x9a\x54\x31\xf4\x33\x4c\xb2\xb1\x79\x5d\x28\x1a\x4c\x2f\x00\x58\xbb\x0a\x0f\x7a\x17\x00\xbb\x18\xc8\xe0\xc8\xb2\x41\x6f\xf7\x4e\x15\xd5\x46\xbd\x8b\xda\xf4\x06\x4b\x15\xfd\x04\x70\x15\xda\xeb\xfb\xbb\x6f\x3f\x3f\x9c\x0c\x03\x54\xe4\x2a\x24\x39\x40\x12\xbf\x4e\xa6\x74\x46\x01\x52\x64\x4d\xa6\x92\x90\x42\x6f\xbc\xc2\x28\x05\xa9\x4f\x11\x64\x90\x0d\xb6\xa6\x61\xda\xd8\x00\x2e\x03\xd9\x18\x06\xc2\x8a\x90\xd1\x4a\x48\x9b\x13\xc5\xe0\x85\x94\x05\x97\xfc\x89\x5a\x56\xf0\x80\xe4\xd5\x00\x6f\x5c\x5d\xa4\xa0\x9d\xdd\x21\x09\x10\x6a\x97\x5b\xf3\xf7\x41\x37\x83\xb8\xb0\x69\xa1\x04\x1b\x90\x8f\x9f\xb1\x82\x64\x55\x01\x3b\x55\xd4\xf8\x16\x94\x4d\xa1\x54\x7b\x20\xf4\xbb\x40\x6d\x3b\xfa\x82\x08\xaf\xe0\xb3\x23\x04\x63\x33\xb7\x86\x8d\x48\xc5\xeb\xab\xab\xdc\x48\x5b\x21\xda\x95\x65\x6d\x8d\xec\xaf\xb4\xb3\x42\x26\xa9\xc5\x11\x5f\xa5\xb8\xc3\xe2\x8a\x4d\xbe\x54\xa4\x37\x46\x50\x4b\x4d\x78\xa5\x2a\xb3\x0c\xa6\x5b\x09\x65\x56\xa6\x3f\xb5\xf9\xc5\x6f\x4e\x6c\x95\xbd\x0f\x2f\x0b\x19\x9b\x77\x26\x42\xae\xcd\x44\xc0\x67\x1b\x18\x06\xd5\x2c\x8d\x5e\x1c\x81\xf6\x43\x1e\x9d\x2f\xb7\x0f\x8f\x87\xd4\x0e\xc1\xe8\xa3\x1f\x70\x3f\x2e\xe4\x63\x08\x3c\x60\xc6\x66\x48\x31\x88\x19\xb9\x32\xe8\x44\x9b\x56\xce\x58\x09\x3f\xba\x30\x68\xfb\xf0\x73\x9d\x94\x46\x7c\xdc\xff\xaa\x91\xc5\xc7\x6a\x05\x37\xca\x5a\x27\x90\x20\xd4\x55\xaa\x04\xd3\x15\xdc\x59\xb8\x51\x25\x16\x37\x8a\xf1\x87\x07\xc0\x23\xcd\x4b\x0f\xec\x79\x21\xe8\x32\x5e\x5f\x38\xa2\xd6\x99\x68\x59\xe9\xf8\x8d\xd7\x97\xff\x7c\x65\xdf\x93\xdb\x99\x14\xa9\x3f\xd7\x8b\xf3\x75\x47\x34\xec\x61\x32\x83\x3e\xea\x3a\xf0\x74\x50\xe5\x5d\xd4\xa1\xb0\x3c\xcb\xe4\xa6\x5f\x5f\xfe\xcb\x1c\x85\x68\x1d\x59\xef\x40\x9c\x03\xe9\x69\xc3\xfd\xa7\x9d\xcd\x4c\x3e\x36\x03\xa0\xd2\x34\x1c\x16\xaa\xb8\x9f\xd5\x01\xd3\xa0\xf7\x05\x06\x40\xb7\x5f\x20\xbe\x99\x85\x13\x9a\x7d\x46\x1a\xc2\x74\xb8\x34\x72\xe9\x60\x78\xc6\x0c\xf5\xc4\xd3\x61\x9c\x47\x51\x69\x8d\xcc\x1f\x71\x7f\xf7\xfe\x02\x27\x00\x14\x73\x5d\xe2\x17\x57\xe0\x35\xd9\xef\xd4\x70\xfb\x1c\xb9\xf2\x42\x53\xb4\xd2\x9b\xd9\x48\x24\xce\x15\xa8\xc6\xb2\xb2\x49\xc1\x0b\x37\xce\x1c\x3d\x29\x4a\x1f\x90\xfd\xa1\xf3\xfb\x0b\xf9\x30\x6d\x05\x61\x3e\x38\xee\xce\x34\x81\x51\x13\xca\x75\x1b\xcc\x0b\x75\x9c\x65\xff\x05\xf9\xdc\x49\xb2\x91\xd9\x9e\xed\xaf\x49\xfc\xc8\xf9\x37\x3e\xbd\x33\x4f\x3d\xf8\x7e\x40\x93\xd0\x67\xb2\x9b\xb1\x35\x9e\x4d\x44\x19\xcb\x70\x7f\xfb\x79\x89\xd6\xb7\x5a\x29\x78\xd2\x0d\xa7\xcd\x58\xda\x36\x9b\x83\xf6\x5d\x41\x66\x0a\x0c\xdc\xf6\xf8\xe9\x61\x48\x64\x99\xa3\x52\x49\x6c\x38\x27\xdc\x1b\x45\x35\xee\xf0\x11\xf7\x67\xbb\xd5\xc8\xbe\xe0\xce\xc1\xf8\x11\xbf\xb6\xb8\xff\x61\xde\xc4\x8e\x74\xac\xca\x66\x96\xe5\x5b\xfc\x3e\x72\x7b\x74\x5b\xbc\xac\xaa\x22\x48\x97\xd2\x51\x58\xfc\x10\x92\xfb\x22\x05\xf8\x5c\x19\x9a\x28\xe5\x36\x06\xc6\xca\xaf\xbf\xcc\x28\xf7\xed\x67\x1e\xba\xf5\xfe\x47\x98\x11\xf2\xe6\x52\x70\x5e\x2e\xf7\xa0\x79\x64\xb6\x05\x75\x72\x2a\x42\xf6\x1a\x1a\x30\x65\x85\xc4\xce\x2a\x19\x61\xae\x93\x2a\xb9\x3b\x4a\xfa\xa6\xd5\x77\x22\xfe\x8a\xe3\x0f\x5c\xdf\xc6\x2b\x2d\x4b\xc5\xc3\x94\x9f\x81\xa2\xb3\xf7\x07\x7f\x11\x1c\x49\xc4\x29\x0b\xa2\x7c\x6b\x47\x1e\xff\xc4\x75\x55\x0e\x4d\x31\x82\xe5\x68\xb2\xcf\x86\x2b\x4e\x2a\x22\xd5\xe7\xd7\xce\x66\xfe\xae\x78\xfb\x2c\x34\x42\x35\xe7\x36\x54\x27\xae\x06\x5d\xdf\xc2\x75\xa0\x54\xbc\x6d\xae\x66\xe1\x9f\x9d\x2f\x5d\x71\x49\x9d\x81\x56\x16\x72\xb4\x48\x6a\x84\x4f\x66\x1c\x7e\xb1\x80\xa6\x9d\x9e\x8e\xc9\x01\x83\x23\x85\x1e\x7d\x0f\x77\x02\x5f\x78\xc3\xbb\x23\x34\xbd\x6d\x07\xce\x34\xa4\xd6\x54\x32\x8d\x26\x72\xa5\x98\x9f\x1c\x8d\x54\xd5\x89\xbd\xf7\x8d\x58\x9b\x3a\xed\xb2\x60\x41\xa2\xd8\xe8\x7e\x3f\x2e\x6e\x0c\x9f\x33\xfb\xf0\x19\x98\xab\xf3\x6e\x0f\x87\x9b\x43\x3c\xe5\x21\x70\x03\x54\x8a\x54\xd9\xbf\x06\xce\x6e\x27\xe3\x6c\x75\xb2\x57\xd4\xdd\x20\x93\xa0\x22\x7f\x6d\x0c\x63\x1e\x9d\x01\x2e\x3d\x14\xa6\x9b\xc3\x57\xe1\xe2\x23\x3f\x76\x6e\xcc\xf3\xf6\x1c\x67\xb7\x34\xf5\x82\xf7\x5f\x5b\x36\xeb\xb3\xdb\xff\x9c\x1a\xe3\xe7\xc3\xb2\xd3\x04\xf4\x26\x22\x62\x67\xdd\x70\x45\x49\xcd\xe7\xde\x71\x5d\xc2\x48\x3b\x4c\x3f\x44\x86\x19\x6d\xb0\x4f\x10\xfc\x63\xb0\xa0\xc5\xb2\x74\x1c\xde\x7f\x7c\xd7\x97\x1f\x67\xdb\x1d\x26\x6f\xbb\xe1\xdd\x29\x3e\x7c\xac\xe0\x4e\x40\x3b\x22\xe4\xca\xd9\xf4\xf0\x76\xd4\xce\xbf\xe1\x8e\xe6\xb7\x23\x1a\x9f\x36\x46\x6f\xbc\x41\xcd\xd3\x05\x38\x0b\x65\x1d\x9f\xb3\x20\xd9\x07\x65\xd7\xf7\x77\xcd\xd3\xc9\x74\xf3\xf6\xda\x04\x1c\x8d\xc5\x60\x30\x02\xb1\x06\xa1\x3a\x52\x39\x8b\x23\x95\x63\x77\xa4\x4e\x0e\x0f\x50\x6d\x24\x9a\x88\xc2\x3f\xff\x2e\xfe\x0b\x00\x00\xff\xff\x96\x77\xd1\x6e\x21\x16\x00\x00")

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

	info := bindataFileInfo{name: "cluster.bytebuilders.dev_clusteruserauths.yaml", size: 5665, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"cluster.bytebuilders.dev_cloudcredentials.yaml":         clusterBytebuildersDev_cloudcredentialsYaml,
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
	"cluster.bytebuilders.dev_cloudcredentials.yaml":         {clusterBytebuildersDev_cloudcredentialsYaml, map[string]*bintree{}},
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
