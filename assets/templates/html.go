// Code generated by go-bindata.
// sources:
// templates/html/common_footer.html
// templates/html/common_header.html
// templates/html/confirm.html
// templates/html/confirm_update.html
// templates/html/email_confirmation.html
// templates/html/index.html
// templates/html/subscribe.html
// templates/html/subscribe_confirm.html
// templates/html/success.html
// templates/html/unsubscribe.html
// templates/html/unsubscribe_confirm.html
// DO NOT EDIT!

package templates

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

var _templatesHtmlCommon_footerHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\xce\xcf\xcd\xcd\xcf\x8b\x4f\xcb\xcf\x2f\x49\x2d\x52\xaa\xad\xe5\xe2\xe4\xb4\xd1\x4f\xc9\x2c\xb3\xe3\xe2\xb4\xd1\x4f\xca\x4f\xa9\xb4\xe3\xb2\xd1\xcf\x28\xc9\xcd\xb1\xe3\xaa\xae\x4e\xcd\x4b\xa9\xad\xe5\x02\x04\x00\x00\xff\xff\x51\x34\x73\xde\x3d\x00\x00\x00")

func templatesHtmlCommon_footerHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlCommon_footerHtml,
		"templates/html/common_footer.html",
	)
}

func templatesHtmlCommon_footerHtml() (*asset, error) {
	bytes, err := templatesHtmlCommon_footerHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/common_footer.html", size: 61, mode: os.FileMode(420), modTime: time.Unix(1568293556, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlCommon_headerHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\x31\x6f\xe3\x30\x0c\x85\x67\xfb\x57\xe8\x78\xeb\x09\xc6\x6d\x37\x48\x5e\xae\x5d\xdb\x0e\x5d\x3a\x15\xb4\xc4\xc4\x04\x6c\xca\x95\x19\x27\x85\xe1\xff\x5e\x24\x76\xdb\xa4\xd3\x93\xa8\xf7\xbe\x27\xce\x73\xa4\x1d\x0b\x19\x08\xa9\xef\x93\xbc\xb6\x84\x91\x32\x2c\x8b\xfb\x75\xf7\xf8\xff\xf9\xe5\xe9\xde\xb4\xda\x77\x75\xe9\x56\x29\xdc\xd9\x51\x97\x45\xe1\x7a\x52\x34\xa1\xc5\x3c\x92\x7a\x38\xe8\xce\xfe\x03\x53\x7d\x3f\xb5\xaa\x83\xa5\xb7\x03\x4f\x1e\x4e\xf6\x80\x36\xa4\x7e\x40\xe5\xa6\x23\x30\x21\x89\x92\xa8\x07\x26\x4f\x71\x4f\x37\x49\xc1\x9e\x3c\x4c\x4c\xc7\x21\x65\xbd\x32\x1f\x39\x6a\xeb\x23\x4d\x1c\xc8\x5e\x2e\x7f\x0c\x0b\x2b\x63\x67\xc7\x80\x1d\xf9\xbf\x9f\x20\x65\xed\xa8\x76\xd5\xaa\x3f\xd0\x91\xc6\x90\x79\x50\x4e\x72\x45\x5f\xa3\xae\xda\x36\x74\x4d\x8a\xef\x75\x59\x16\xc6\x18\xe3\x04\x27\x13\x3a\x1c\x47\x0f\x82\x53\x83\xd9\xac\x62\xe9\x34\xa0\x44\xdb\xed\xe1\xd2\x82\xb7\x2e\xdb\x64\x94\x08\xa6\xcd\xb4\xf3\xf0\x1b\xea\x87\xcb\xd8\x55\x58\x6f\xe0\x4a\x70\xda\xce\xdb\x24\xf2\x57\xd5\xf9\x6f\xc8\x42\x19\xea\x72\x9e\x49\xe2\xb2\x94\x1f\x01\x00\x00\xff\xff\x70\x5d\xaf\xec\xb4\x01\x00\x00")

func templatesHtmlCommon_headerHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlCommon_headerHtml,
		"templates/html/common_header.html",
	)
}

func templatesHtmlCommon_headerHtml() (*asset, error) {
	bytes, err := templatesHtmlCommon_headerHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/common_header.html", size: 436, mode: os.FileMode(420), modTime: time.Unix(1568294163, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlConfirmHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x41\x8b\xdc\x30\x0c\x85\xcf\xc9\xaf\x10\xba\x67\x02\x7b\x76\x72\x29\xbd\x15\x5a\xba\xed\xa1\xa7\xe2\xc4\xca\xc6\x60\x5b\x41\x71\x16\x06\x33\xff\xbd\xd8\x49\xa6\xb3\xc3\x5e\x8c\xd0\x93\xbe\x67\x9e\x9d\x12\x18\x9a\x6c\x20\xc0\x91\xc3\x64\xc5\xe3\xed\x56\xa7\x14\xc9\x2f\x4e\xc7\xd2\xf6\x9e\xc3\xdf\x99\xb4\x21\xc9\x62\xad\x8c\x7d\x87\xd1\xe9\x75\xed\x70\xd4\x62\xb0\xaf\x01\x00\xd4\xfc\xf2\xd8\x6d\x8e\x8d\xfe\xcb\xce\xbd\x5c\x2e\xaa\x9d\x5f\x8e\xd9\x27\x44\x33\xb0\xb9\x62\x5f\xd7\x55\x4a\x60\x27\xb8\x7c\x15\x61\x81\xec\x56\x3d\xce\x6a\x47\x12\xa1\x9c\x8d\xd1\xe1\x8d\x04\x41\xd8\xd1\xa1\x60\x5f\x57\x19\x9f\xd2\x03\xa1\x52\xad\xb1\xef\x7d\x41\x53\x30\xa5\x55\x57\x6a\x62\xf1\x27\x36\xd7\x08\x9e\xe2\xcc\xa6\xc3\x1f\xdf\x5f\x7f\x21\xe8\x31\x5a\x0e\x1d\x66\xd6\xef\x9f\xdf\xe0\x76\x3b\xe9\x1f\xaf\x94\x77\x9b\x37\xe1\x6d\x39\x75\xe5\xf4\x40\x0e\x26\x96\x0e\xb5\x31\x42\xeb\x7a\x4f\x41\x67\x28\x8c\x6c\x48\xb5\x65\xec\xdc\xb1\x61\xd9\x22\xc4\xeb\x42\x1d\x66\x19\x3f\xf0\x47\x0e\x51\xd8\x21\x04\xed\xef\x03\xd6\x9c\xd5\xe2\xf4\x48\x33\x3b\x43\xd2\xe1\x1f\xde\x04\xc6\x67\xbb\x12\xee\x99\x44\x5d\xa9\x61\x8b\x91\xc3\x61\xb8\x6e\x83\xb7\xf1\x6e\x39\xc4\x00\x43\x0c\xcd\x22\xd6\x6b\xb9\x62\xff\x5a\x74\xd5\xee\x4b\x3b\xa0\xcd\x17\x3b\x5e\x73\xc7\x96\x6c\x8e\xfa\xb3\x2f\x34\x31\xc7\xfd\x0b\xfd\x7f\x89\x7f\x01\x00\x00\xff\xff\x6a\xc4\x6d\xaa\x80\x02\x00\x00")

func templatesHtmlConfirmHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlConfirmHtml,
		"templates/html/confirm.html",
	)
}

func templatesHtmlConfirmHtml() (*asset, error) {
	bytes, err := templatesHtmlConfirmHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/confirm.html", size: 640, mode: os.FileMode(420), modTime: time.Unix(1568297772, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlConfirm_updateHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xd0\xc1\x6a\xc4\x20\x10\x06\xe0\x73\x7c\x8a\xc1\x7b\x12\xd8\xb3\xf5\x52\xfa\x1c\x8b\xd5\x49\x23\xa8\x23\xa3\xbb\x50\x64\xdf\xbd\x24\xa6\x6d\x28\x4b\x2f\x1e\xfc\x9d\xef\xc7\x69\x0d\x1c\x2e\x3e\x21\x48\x4b\x69\xf1\x1c\xaf\xb7\xec\x4c\x45\xf9\x78\x88\xd6\x2a\xc6\x1c\x4c\xdd\xd3\x18\x29\x5d\x57\x34\x0e\x79\x0b\x85\x72\xfe\x0e\x36\x98\x52\x5e\xa4\x35\xec\xa4\x16\x00\x00\x6a\xbd\x9c\x6f\xc7\x63\x42\xbf\x76\xde\x54\x4f\x09\x2c\xc5\x1c\xb0\xa2\x9a\xd7\xcb\x31\xf6\x47\x1b\xdf\xc9\x7d\x4a\x2d\xc4\xd0\x1a\xf8\x05\xa6\x37\x66\x62\xd8\x8a\x87\xf3\x5b\x13\x90\x2b\xec\xe7\xe8\x4c\xfa\x40\x96\xc0\x14\xf0\x48\xa4\x16\xc3\xc6\xb7\x76\x12\x06\x35\x3b\x7f\x3f\x6c\x0c\x05\xff\x67\xcb\xcd\x5a\x2c\xe5\xa9\xab\xb2\x9e\xa6\x49\xcd\x59\xff\xa8\x43\x67\x93\xdb\xab\xfa\xe7\x7a\xf2\x5d\xfb\x64\xaf\x0b\x51\xed\x7b\xfd\x9d\xfd\x0a\x00\x00\xff\xff\xdd\xd3\xf3\x2e\x9c\x01\x00\x00")

func templatesHtmlConfirm_updateHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlConfirm_updateHtml,
		"templates/html/confirm_update.html",
	)
}

func templatesHtmlConfirm_updateHtml() (*asset, error) {
	bytes, err := templatesHtmlConfirm_updateHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/confirm_update.html", size: 412, mode: os.FileMode(420), modTime: time.Unix(1568297803, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlEmail_confirmationHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x56\x48\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\xcd\x4d\xcc\xcc\x89\x4f\xce\xcf\x4b\xcb\x2c\xca\x4d\x2c\xc9\xcc\xcf\x53\x52\xa8\xad\xe5\xe2\xb2\x49\x54\xc8\x28\x4a\x4d\xb3\x55\x52\x56\xb2\xab\xae\x56\xd0\x73\xce\x4f\x49\x55\xa8\xad\xb5\xd1\x4f\xb4\xe3\xe2\xaa\xae\x56\x48\xcd\x4b\x01\xa9\x03\x04\x00\x00\xff\xff\x8d\x96\x16\x1b\x4a\x00\x00\x00")

func templatesHtmlEmail_confirmationHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlEmail_confirmationHtml,
		"templates/html/email_confirmation.html",
	)
}

func templatesHtmlEmail_confirmationHtml() (*asset, error) {
	bytes, err := templatesHtmlEmail_confirmationHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/email_confirmation.html", size: 74, mode: os.FileMode(420), modTime: time.Unix(1567952259, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\x5d\xaa\xc2\x30\x10\x46\xdf\xb3\x8a\x8f\x2c\xa0\xdd\xc0\xdc\xac\xe1\x82\xf8\x2c\x69\x33\xa5\x81\xfc\x48\x92\x82\x30\x64\xef\xd2\x22\xa2\x28\xbe\xcd\xc7\x70\x0e\x47\x04\x8e\x17\x9f\x18\xda\x27\xc7\x37\x8d\xde\x95\x08\x1a\xc7\x6b\xb0\x8d\xa1\xe7\x1c\x63\x4e\x97\x95\xad\xe3\x72\xbc\x15\x6d\xc1\x28\x05\x00\x14\xbc\x21\x8b\xb5\xf0\xf2\xa7\x45\x30\xfc\xdb\xb6\xd6\xe1\xb4\x4d\x75\x2e\x7e\x62\xf4\xae\xcd\x73\xd1\x68\x0d\x8d\xc1\x9b\x1d\xfd\xc5\x9f\x53\x7d\x33\xbc\xec\x4f\x87\xa2\xf1\xe8\xf9\x56\xbd\xe4\xdc\x1e\xd5\x22\xe0\xe4\xf6\xeb\x1e\x00\x00\xff\xff\x98\xed\x81\x03\xf4\x00\x00\x00")

func templatesHtmlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlIndexHtml,
		"templates/html/index.html",
	)
}

func templatesHtmlIndexHtml() (*asset, error) {
	bytes, err := templatesHtmlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/index.html", size: 244, mode: os.FileMode(420), modTime: time.Unix(1568293427, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlSubscribeHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x93\x4f\x6b\xdc\x3c\x10\xc6\xcf\xf6\xa7\x18\x74\xc9\xc9\x6b\xc8\x59\xd6\x2d\xf0\xbe\x50\x68\x68\x1a\x4a\x4f\x45\xb6\xc6\x6b\x51\xfd\x31\xa3\xf1\xb6\xc6\xec\x77\x2f\xf2\xda\x59\xa7\xcd\x45\x48\xa3\x99\x9f\x1e\x49\xcf\x2c\x0b\x18\xec\x6d\x40\x10\x69\x6a\x53\x47\xb6\x45\x71\xbd\x96\xcb\xc2\xe8\x47\xa7\x19\x41\x74\xd1\xfb\x18\x7e\x0c\xa8\x0d\x52\xde\x2c\xa5\xb1\x17\xe8\x9c\x4e\xa9\x11\x9d\x26\x23\x54\x09\x00\x20\x87\xc7\x63\xb4\xda\x2a\xd4\xcb\x4e\x96\xf5\xf0\xb8\xa5\xfe\x45\xa8\xda\x68\x66\xa1\xca\xb2\x58\x16\xb0\x3d\x9c\x9e\x88\x22\x41\x3e\xac\x38\xe6\x6a\x87\xc4\xb0\x8e\x95\xd1\xe1\x8c\x24\x80\xa2\xc3\x6d\x47\xa8\xb2\xc8\xf8\x65\x39\x10\x0a\x59\x1b\x7b\x51\x2b\x1a\x83\x59\x43\x65\x21\xfb\x48\x7e\xc7\xe6\xb9\x00\x8f\x3c\x44\xd3\x88\xe7\xcf\x2f\x5f\x05\xe8\x8e\x6d\x0c\x8d\xc8\xac\xd7\x2f\x9f\xe0\x7a\xdd\xe9\xef\x25\xe5\xda\xea\x4c\x71\x1a\xf7\x7d\xe9\x74\x8b\x0e\xfa\x48\x8d\xd0\xc6\x10\xa6\x24\xd4\x93\xd7\xd6\xc1\xb6\x94\xf5\x9a\xb2\xe7\xdb\x30\x4e\x0c\x3c\x8f\xd8\x08\xcc\x79\xe2\x1d\xbc\x8b\x81\x29\x3a\x01\x41\x7b\xbc\x23\xc1\x9a\xc3\x42\x93\xd5\x95\xc1\xdb\x43\x9b\x76\xde\x48\xff\xa1\x1b\x05\x8c\x4e\x77\x38\x44\x67\x90\x1a\xf1\x3d\x4e\x04\x78\x94\xf3\x26\x3c\x79\xed\xdc\xca\x3d\x14\x1f\xa5\x30\xfe\x66\xc8\x43\xe5\x27\x46\x23\xd4\x37\x7c\x70\x0e\x02\x5e\x90\x20\x0d\x9a\x10\xe6\x3b\xfe\x97\xe5\x01\x74\x98\x63\x40\x40\x97\xf0\x24\xeb\xf5\x04\x75\xff\x93\xb2\x90\xed\xc4\x1c\xc3\x76\xfd\x34\xb5\xde\xf2\xdb\xa1\x2d\x07\x68\x39\x54\x23\x59\xaf\x69\x5e\xbd\xe4\x2d\xcb\xfa\x56\x74\x03\xd4\x59\xda\xe6\xab\x1b\x76\xfd\xa5\x0f\x7d\xd6\xc7\xc8\xd9\x93\x65\x21\x47\xf5\x7f\x9f\xd5\x3e\x10\x82\x8b\xf1\xa7\x0d\x67\xe0\x08\x12\xbd\x9a\x42\xba\x7b\x16\xbd\xca\x71\x1e\x10\xf2\xad\x72\x9a\xb3\x89\x73\x20\x80\xd4\x30\x10\xf6\x37\x9b\x3c\x6b\x1e\xd2\xe9\xf5\x5e\xbc\x9a\x26\xb7\x01\xc4\x00\x31\xbf\xd1\x80\x84\xb2\xd6\xea\x24\xeb\xf1\x5f\xc9\xdb\xfc\xa3\xf6\xdb\x84\xaf\xbd\xb9\xdb\xf8\x4f\x00\x00\x00\xff\xff\xf7\x6c\x15\x4a\xbe\x03\x00\x00")

func templatesHtmlSubscribeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlSubscribeHtml,
		"templates/html/subscribe.html",
	)
}

func templatesHtmlSubscribeHtml() (*asset, error) {
	bytes, err := templatesHtmlSubscribeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/subscribe.html", size: 958, mode: os.FileMode(420), modTime: time.Unix(1568297123, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlSubscribe_confirmHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xc1\x6a\xf3\x30\x10\x84\xcf\xbf\x9f\x62\x7e\x5d\x7a\x4a\x02\x39\xbb\x7e\x89\x1e\x4a\x4f\x41\x96\xd6\x91\xc0\xd6\x9a\xdd\x75\x4a\x30\x7e\xf7\xe2\x44\x87\x50\x7a\x91\x84\x66\xe7\x63\x67\xd6\x15\x91\x86\x5c\x08\x4e\x97\x5e\x83\xe4\x9e\x2e\x81\xcb\x90\x65\xf2\x96\xb9\xb8\x6d\x6b\xd6\xd5\x68\x9a\x47\x6f\x04\x17\x78\x9a\xb8\x5c\x12\xf9\x48\xb2\x8b\x4d\x1b\xf3\x0d\x61\xf4\xaa\xef\x2e\x78\x89\xae\x6b\x00\xa0\x4d\xe7\xd7\xdf\x43\x75\x74\x1f\x4b\x08\xa4\xfa\xbf\x3d\xa5\x73\x9d\xfc\x05\x38\xf4\x1c\xef\xae\x6b\xfe\xbd\x0a\x7e\x24\x31\x3c\xce\x83\x3e\x11\x0e\xc2\x23\x55\x69\x9f\x7f\xc0\xe6\xee\x8b\x17\x01\x4d\x3e\x8f\xf0\x31\x0a\xa9\x22\x79\x45\x4f\x54\x20\x74\xcd\x6a\x24\x14\x8f\xf8\xa4\xb7\x1b\x41\xa9\x18\xee\xbc\xc0\xe3\x35\x77\x05\x58\xf2\x86\xef\x6c\x09\xb9\xa8\xc9\x12\x76\x4d\x31\xb0\x20\xf0\x34\x8f\x64\xb9\x5c\x61\x89\x50\xeb\x9b\x1f\xe6\x59\x78\xdf\xf0\xd8\x9e\xe6\x3d\xc7\x29\xe6\x5b\xcd\xfa\x7c\xd6\xeb\xaf\x66\x07\x66\x7b\x36\xbb\xae\xa0\x12\xb1\x6d\xcd\x4f\x00\x00\x00\xff\xff\xab\x7a\xf9\xe0\xa6\x01\x00\x00")

func templatesHtmlSubscribe_confirmHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlSubscribe_confirmHtml,
		"templates/html/subscribe_confirm.html",
	)
}

func templatesHtmlSubscribe_confirmHtml() (*asset, error) {
	bytes, err := templatesHtmlSubscribe_confirmHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/subscribe_confirm.html", size: 422, mode: os.FileMode(420), modTime: time.Unix(1568297793, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlSuccessHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xc9\x31\x0a\xc0\x20\x0c\x05\xd0\xdd\x53\x04\x2f\x26\xa2\x5f\x2a\x34\x49\xd1\x74\x0a\xb9\x7b\xa1\xd0\xb1\xf3\x73\xa7\x8e\x31\x05\x94\x9b\xca\x98\x8b\xcb\xbe\x5b\xc3\xde\x39\x22\xb9\x1b\xf8\x3a\xab\xbd\xcc\xac\x52\x0e\xd4\x8e\xf5\x83\x43\xd5\x3e\x24\x48\xa7\x88\xf4\x04\x00\x00\xff\xff\xe9\xd5\x22\x89\x62\x00\x00\x00")

func templatesHtmlSuccessHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlSuccessHtml,
		"templates/html/success.html",
	)
}

func templatesHtmlSuccessHtml() (*asset, error) {
	bytes, err := templatesHtmlSuccessHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/success.html", size: 98, mode: os.FileMode(420), modTime: time.Unix(1560917570, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlUnsubscribeHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x53\x41\x6f\xdc\x2c\x10\x3d\xdb\xbf\x62\xc4\xe5\x3b\x79\x2d\xe5\x8c\xb9\x45\xfa\x2a\x55\x6a\xd4\x24\x87\x9e\x2a\x6c\xc6\x01\x15\x18\x0b\xc6\x69\x57\xd6\xfe\xf7\x0a\xdb\x9b\x78\xab\xbd\x60\xc3\xbc\x79\xf3\x18\xde\x2c\x0b\x18\x1c\x5d\x44\x10\x73\xcc\x73\x9f\x87\xe4\x7a\x14\x97\x4b\xbd\x2c\x8c\x61\xf2\x9a\x11\xc4\x40\x21\x50\xfc\x69\x51\x1b\x4c\x25\x58\x4b\xe3\xde\x61\xf0\x3a\xe7\x4e\x0c\x3a\x19\xa1\x6a\x00\x00\x69\x1f\x8e\xa7\xcd\x9e\xa1\x5e\x3f\xb9\x65\x6b\x1f\x76\xf0\x3f\x1c\x4d\x4f\xe6\x2c\x54\x5d\x57\xcb\x02\x6e\x84\xd3\x63\x4a\x94\xa0\x94\xab\x8e\x58\xed\x31\x31\xac\x6b\x63\x74\x7c\xc3\x24\x20\x91\xc7\x3d\x22\x54\x5d\x15\xfa\x65\x39\x30\x54\xb2\x35\xee\x5d\xad\xd4\x18\x4d\x39\x2a\x98\xba\x92\x23\xa5\x70\x65\x2e\xff\x02\x02\xb2\x25\xd3\x89\xa7\x6f\xcf\x2f\x02\xf4\xc0\x8e\x62\x27\x0a\xdd\xeb\xf7\xaf\x70\xb9\x5c\x0b\x6c\xeb\x51\x59\xc9\x6f\xde\x12\xcd\x53\xc1\x54\xd2\xeb\x1e\x3d\x8c\x94\x3a\xa1\x8d\x49\x98\xb3\x50\x8f\x41\x3b\x0f\xfb\x56\xb6\x2b\x64\x05\xbb\x38\xcd\x0c\x7c\x9e\xb0\x13\x58\x40\xe2\x86\x76\xa0\xc8\x89\xbc\x80\xa8\x03\x7e\xf2\x81\x33\x87\x8d\x4e\x4e\x37\x06\xb7\x4e\x9b\xfe\xbc\x33\xfd\x8f\x7e\x12\x30\x79\x3d\xa0\x25\x6f\x30\x75\xe2\x07\xcd\x09\xf0\xa8\x65\x93\x9c\x83\xf6\x7e\x25\x3d\x64\x1e\x75\x30\xfe\x61\x28\x4b\x13\x66\x46\x23\xd4\x8b\x75\x19\x5c\x06\xb6\x78\x4b\x08\x67\x9a\xe1\xe3\xdd\x0d\x30\xad\x98\x02\x71\xf1\x0d\xbc\xcb\x0c\xbf\x1d\xdb\x93\x6c\xd7\xaa\x7b\x5b\xaf\x2f\x75\x68\x71\x3f\x33\x53\xdc\x7b\x93\xe7\x3e\x38\xfe\x10\xd5\x73\x84\x9e\x63\x33\x25\x17\x74\x3a\x0b\xf5\xbc\xc6\x65\xbb\x25\x7d\x10\xc9\xb6\xc8\xdf\x9d\xb7\x95\x58\x03\x77\x9d\x38\x12\x71\xf1\x6d\x5d\xc9\x49\x7d\x19\xcb\x4d\xfe\x4b\x08\x9e\xe8\x57\x91\xce\x04\x12\x83\x3a\x78\x1a\x83\xba\x7b\x3f\xb6\x18\x41\x6a\xb0\x09\xc7\xcd\x43\x4f\x9a\x6d\x3e\x3d\x5f\x53\x57\x3f\x95\x31\x01\x8a\x40\xef\x98\xc0\x62\x42\xd9\x6a\x75\x92\xed\x74\x23\x77\xff\xdc\x9b\xcc\x5d\xef\x3a\xb6\x57\x7f\xff\x0d\x00\x00\xff\xff\xd4\xce\xad\xed\xdb\x03\x00\x00")

func templatesHtmlUnsubscribeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlUnsubscribeHtml,
		"templates/html/unsubscribe.html",
	)
}

func templatesHtmlUnsubscribeHtml() (*asset, error) {
	bytes, err := templatesHtmlUnsubscribeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/unsubscribe.html", size: 987, mode: os.FileMode(420), modTime: time.Unix(1568295962, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlUnsubscribe_confirmHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcc\x31\x0e\xc3\x20\x0c\x46\xe1\x9d\x53\x58\x2c\xdd\xda\x0b\x44\xb9\x46\xc7\x88\xc0\x8f\x8a\x14\x6c\x04\x4e\x17\x8b\xbb\x57\x55\x97\x0e\x99\x9f\xde\x67\x46\x09\xb9\x30\xc8\x9f\x3c\xce\x7d\xc4\x5e\x76\x6c\x51\x38\x97\x5e\x83\x16\x61\x3f\xa7\x33\x53\xd4\x76\x04\x05\xf9\x28\xb5\x0a\x6f\x2f\x84\x84\xfe\x8d\x6e\x69\xeb\x13\xb7\x37\x68\x80\x95\x02\xfd\xdf\x84\x1a\xca\x71\x5f\x1e\x6d\x75\x57\x4c\x16\xd1\x1f\x63\x46\xe0\x44\x73\xba\x4f\x00\x00\x00\xff\xff\xb3\x49\xe6\xb3\x95\x00\x00\x00")

func templatesHtmlUnsubscribe_confirmHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlUnsubscribe_confirmHtml,
		"templates/html/unsubscribe_confirm.html",
	)
}

func templatesHtmlUnsubscribe_confirmHtml() (*asset, error) {
	bytes, err := templatesHtmlUnsubscribe_confirmHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/unsubscribe_confirm.html", size: 149, mode: os.FileMode(420), modTime: time.Unix(1567952223, 0)}
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
	"templates/html/common_footer.html":       templatesHtmlCommon_footerHtml,
	"templates/html/common_header.html":       templatesHtmlCommon_headerHtml,
	"templates/html/confirm.html":             templatesHtmlConfirmHtml,
	"templates/html/confirm_update.html":      templatesHtmlConfirm_updateHtml,
	"templates/html/email_confirmation.html":  templatesHtmlEmail_confirmationHtml,
	"templates/html/index.html":               templatesHtmlIndexHtml,
	"templates/html/subscribe.html":           templatesHtmlSubscribeHtml,
	"templates/html/subscribe_confirm.html":   templatesHtmlSubscribe_confirmHtml,
	"templates/html/success.html":             templatesHtmlSuccessHtml,
	"templates/html/unsubscribe.html":         templatesHtmlUnsubscribeHtml,
	"templates/html/unsubscribe_confirm.html": templatesHtmlUnsubscribe_confirmHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"html": &bintree{nil, map[string]*bintree{
			"common_footer.html":       &bintree{templatesHtmlCommon_footerHtml, map[string]*bintree{}},
			"common_header.html":       &bintree{templatesHtmlCommon_headerHtml, map[string]*bintree{}},
			"confirm.html":             &bintree{templatesHtmlConfirmHtml, map[string]*bintree{}},
			"confirm_update.html":      &bintree{templatesHtmlConfirm_updateHtml, map[string]*bintree{}},
			"email_confirmation.html":  &bintree{templatesHtmlEmail_confirmationHtml, map[string]*bintree{}},
			"index.html":               &bintree{templatesHtmlIndexHtml, map[string]*bintree{}},
			"subscribe.html":           &bintree{templatesHtmlSubscribeHtml, map[string]*bintree{}},
			"subscribe_confirm.html":   &bintree{templatesHtmlSubscribe_confirmHtml, map[string]*bintree{}},
			"success.html":             &bintree{templatesHtmlSuccessHtml, map[string]*bintree{}},
			"unsubscribe.html":         &bintree{templatesHtmlUnsubscribeHtml, map[string]*bintree{}},
			"unsubscribe_confirm.html": &bintree{templatesHtmlUnsubscribe_confirmHtml, map[string]*bintree{}},
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
