/*
	Refer to github.com/noelyahan/impexp
	Implements functions to export base64 format image data
*/

package nuts

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"path/filepath"
	"strings"
)

type Base64Exporter struct {
	ext string
	img image.Image
	inc bool
}

func NewBase64Exporter(ext string, img image.Image, inc bool) Base64Exporter {
	return Base64Exporter{ext, img, inc}
}

func (o Base64Exporter) Export() (base64Str string, err error) {
	img, ext := o.img, o.ext
	if img == nil {
		return "", errors.New("null data")
	}
	buf := bytes.NewBuffer(make([]byte, 0))
	defer buf.Reset()
	if ext == ".jpg" || ext == ".jpeg" {
		err = jpeg.Encode(buf, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	} else if ext == ".png" {
		err = png.Encode(buf, img)
	} else {
		return "", fmt.Errorf("[%s] format is not supported", ext)
	}
	if err != nil {
		return "", fmt.Errorf("encode failed: [%w]", err)
	}
	base64Str = base64.StdEncoding.EncodeToString(buf.Bytes())
	if o.inc == true {
		base64Str = fmt.Sprintf("data:image/%s;base64,%s",
			strings.TrimLeft(ext, "."), base64Str)
	}
	return base64Str, nil
}

// Export base64 format image data form the image path directly
func ImageExportBase64(path string, inc bool) (base64Str string, err error) {
	img, err := ImageRead(path)
	if err != nil {
		return "", err
	}
	ex := NewBase64Exporter(filepath.Ext(path), img, inc)
	if base64Str, err = ex.Export(); err == nil {
		return base64Str, nil
	} else {
		return "", err
	}
}
