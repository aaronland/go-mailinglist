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

var _templatesHtmlConfirmHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x56\x48\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\xce\xcf\x4b\xcb\x2c\xca\x55\xaa\xad\xe5\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x01\x0b\xe7\xe6\xe6\xe7\xc5\x67\xa4\x26\xa6\xa4\x16\xe1\x90\x4c\xcb\xcf\x2f\x81\x49\x2a\xa4\xe6\xa5\x28\xd4\xd6\x72\x01\x02\x00\x00\xff\xff\x69\x90\xa5\x70\x5a\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/confirm.html", size: 90, mode: os.FileMode(420), modTime: time.Unix(1560917555, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlConfirm_updateHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xc9\x31\x0a\x80\x30\x0c\x05\xd0\xbd\xa7\x08\xbd\x58\x29\xe6\x17\x0b\x26\x29\x25\x4e\x21\x77\x17\x04\x47\xe7\x17\x41\x8c\x31\x15\x54\x0f\xd3\x31\xb7\xb4\x7b\x71\x77\xd4\xcc\x12\xe1\x90\x75\x75\x7f\x55\xc4\xb4\x9d\xe8\x8c\xfd\x83\xc3\xcc\x3f\x24\x28\x53\x66\x79\x02\x00\x00\xff\xff\x82\xf5\x03\x9a\x61\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/confirm_update.html", size: 97, mode: os.FileMode(420), modTime: time.Unix(1567952396, 0)}
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

var _templatesHtmlSubscribeHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x3f\x8b\xdc\x30\x10\xc5\xeb\xd5\xa7\x18\xa6\x49\xe5\x35\x5c\x6d\xbb\x0b\xa4\xcb\xc1\x05\x42\xaa\x20\x59\xb3\xb1\x40\xd2\x18\x69\xbc\x89\x11\xfe\xee\xc1\x7f\x6e\xcf\x0b\xd7\x18\x6c\xbd\xf7\x7b\xcf\x9a\x29\x05\x2c\xdd\x5c\x24\xc0\x3c\x99\xdc\x27\x67\x08\x97\x45\x95\x22\x14\x46\xaf\x85\x00\x7b\x0e\x81\xe3\xef\x81\xb4\xa5\xb4\x1e\xaa\xc6\xba\x3b\xf4\x5e\xe7\xdc\x62\xaf\x93\xc5\x4e\x01\x00\x34\xc3\xcb\xf9\x6b\x75\x38\xba\xb7\x77\x72\x53\x0f\x2f\x87\xf4\xc6\x29\x3c\x89\x0d\xdb\x19\x21\x90\x0c\x6c\x5b\x7c\xfd\xfe\xf6\x03\x41\xf7\xe2\x38\xb6\x58\x0a\x5c\x5f\xb5\x0c\xf9\xfa\x40\xc1\xb2\x60\xa7\x2e\xea\x72\xee\xb2\x42\xab\x3f\x89\xa7\x71\x3d\xdb\x72\xbc\x36\xe4\xe1\xc6\xa9\x45\x6d\x6d\xa2\x9c\xb1\xfb\x1a\xb4\xf3\x70\xbc\x36\xf5\x26\x79\xd7\xbb\x38\x4e\x02\x32\x8f\xd4\x22\xad\x3a\x7c\x82\xf7\x1c\x25\xb1\x47\x70\xf6\x03\x08\x3a\x39\x5d\x59\xda\x9b\x59\x33\x1f\xd6\x6f\xe4\x47\x84\xd1\xeb\x9e\x06\xf6\x96\x52\x8b\xbf\x78\x4a\x40\xe7\xfc\x47\xd3\x1c\xb4\xf7\x1b\xf7\x64\x3e\x67\x0b\xfd\x13\x58\x1f\x55\x98\x84\x2c\x76\x3f\xe9\x8b\xf7\x10\xe9\x4e\x09\xf2\xa0\x13\xc1\xfc\x81\xff\xeb\x64\x00\x1d\x67\x8e\x04\xe4\x33\x5d\x9b\x7a\x4b\xe8\xd4\xa5\xa9\xad\xbb\xef\x97\x67\x26\x11\x8e\xc7\xff\xe6\xc9\x04\x27\x8f\x50\x23\x11\x8c\xc4\x6a\x4c\x2e\xe8\x34\x6f\x73\x0c\x4e\x9a\x7a\x37\x75\x6a\x9f\x64\xbd\x96\xeb\xd4\x01\xfd\x6c\x73\x6e\xcc\xb2\x6f\x4e\x29\x40\xd1\xc2\xb2\xa8\xff\x01\x00\x00\xff\xff\xdd\x16\xb7\x70\x79\x02\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/subscribe.html", size: 633, mode: os.FileMode(420), modTime: time.Unix(1568293981, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlSubscribe_confirmHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcc\x31\x0e\xc2\x30\x0c\x46\xe1\x3d\xa7\xb0\xba\xb0\xc1\x05\xaa\x5e\x83\xb1\x4a\x93\x3f\x22\x52\x6d\x47\x89\x61\xb1\x72\x77\x84\x58\x18\x98\x9f\xde\xe7\x4e\x19\xa5\x0a\x68\x19\xcf\x63\xa4\x5e\x0f\xec\x49\xa5\xd4\xce\xd1\xaa\xca\x32\x67\x70\x37\x70\x3b\xa3\x81\x96\xa4\xcc\x2a\xfb\x03\x31\xa3\x7f\x62\x58\xdb\x76\xc7\xe5\x05\x1a\x10\xa3\x48\xbf\x37\x81\x63\x3d\xaf\xeb\xad\x6d\xe1\x1f\x53\x54\xed\xcb\xb8\x13\x24\xd3\x9c\xe1\x1d\x00\x00\xff\xff\xd3\xb5\x2d\xd4\x93\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/subscribe_confirm.html", size: 147, mode: os.FileMode(420), modTime: time.Unix(1567952005, 0)}
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

var _templatesHtmlUnsubscribeHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x41\x6a\xc4\x30\x0c\x45\xf7\x3e\x85\xd0\x01\x9a\x0b\x38\xde\x75\x57\x68\xe9\xb4\xeb\xe2\xc4\x4a\x63\x88\x25\x63\xcb\xd0\x60\x7c\xf7\x32\xcc\x50\xba\x98\xa5\x78\x0f\xf4\x5f\xef\x10\x68\x8b\x4c\x80\x8d\x6b\x5b\xea\x5a\xe2\x42\x38\x86\xe9\x5d\x29\xe5\xc3\x2b\x01\xae\x92\x92\xf0\xd7\x4e\x3e\x50\xb9\x42\x63\x17\x09\xa7\x33\x76\x93\x92\x20\x91\xee\x12\x66\x7c\x7b\xbd\x7c\x20\xf8\x55\xa3\xf0\x8c\xbd\xc3\xd3\xe7\xfb\x0b\x8c\x81\xce\xd8\xc8\xb9\x29\xe8\x99\x69\x46\xa5\x1f\x45\x60\x9f\x68\x46\x1f\x42\xa1\x5a\x11\x62\xf8\x77\xe4\xc3\xaf\xb4\xcb\x11\xa8\xcc\xf8\xcc\x4a\x05\x4e\x69\x05\x28\xf9\x78\xc0\x5d\x83\x3f\x7d\x72\xc6\x2e\x4d\x55\xf8\xfe\xa1\xb6\x25\x45\x45\x77\x89\xdf\x0c\x2d\xdb\xe9\x46\x9d\xb1\xd3\x75\xb1\x33\x8f\xf2\x36\x11\xbd\xe5\xf5\x0e\xc4\x01\xc6\x30\xbf\x01\x00\x00\xff\xff\x1f\x63\x7d\x95\x20\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/unsubscribe.html", size: 288, mode: os.FileMode(420), modTime: time.Unix(1567952174, 0)}
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
