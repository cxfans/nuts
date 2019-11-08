/*
	Refer to github.com/noelyahan/impexp
	Implements ImageReader, ImageWriter interfaces for images(jpg, png) inputs and outputs.

	Loaders pkg supports

		ImageReader
		ImageRead(path string) (image.Image, error)
		ImageWriter
		ImageWrite(img image.Image, path string) error
*/

package nuts

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type ImageReader interface {
	Read() (image.Image, error)
}

type Image struct {
	path string
	img  image.Image
}

func NewImageReader(path string) ImageReader {
	return Image{path, nil}
}

func (o Image) Read() (image.Image, error) {
	ext := filepath.Ext(o.path) // return .ext
	f, err := os.Open(o.path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	var img image.Image
	if ext == ".jpg" || ext == ".jpeg" {
		img, err = jpeg.Decode(f)
	} else if ext == ".png" {
		img, err = png.Decode(f)
	} else {
		return nil, fmt.Errorf("[%s] format is not supported", ext)
	}
	if err != nil {
		return nil, fmt.Errorf("decode failed: [%w]", err)
	}
	return img, nil
}

// Read the image directly from the path and return the data
func ImageRead(path string) (image.Image, error) {
	reader := NewImageReader(path)
	return reader.Read()
}

type ImageWriter interface {
	Write() error
}

func NewImageWriter(img image.Image, path string) ImageWriter {
	return Image{path, img}
}

func (o Image) Write() error {
	img := o.img
	if img == nil {
		return errors.New("empty file")
	}
	ext := filepath.Ext(o.path)
	f, err := os.Create(o.path)
	defer f.Close()
	if err != nil {
		return err
	}
	if ext == ".jpg" || ext == ".jpeg" {
		err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	} else if ext == ".png" {
		err = png.Encode(f, img)
	} else {
		_ = os.Remove(o.path)
		return fmt.Errorf("[%s] format is not supported", ext)
	}
	if err != nil {
		return fmt.Errorf("encode failed: [%w]", err)
	}
	return nil
}

// Write the image directly from the data and save to the path
func ImageWrite(img image.Image, path string) error {
	writer := NewImageWriter(img, path)
	return writer.Write()
}
