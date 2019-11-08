// Refer to github.com/noelyahan/mergi

package nuts

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
)

// Crop uses Go standard image.Image, the starting X, Y position as Go standard
// image.Point crop width and height as image.Point returns the crop image output

// x, y: the top left starting point
// w, h: the final size width * height
func Crop(img image.Image, x, y, w, h int) (image.Image, error) {
	if img == nil {
		return nil, errors.New("null data")
	}
	if x < 0 || y < 0 || w < 0 || h < 0 {
		return nil, errors.New("please input more than 0 value for bounds")
	}
	r := image.Rect(0, 0, w, h)
	resImg := image.NewRGBA(r)
	draw.Draw(resImg, r, img, image.Pt(x, y), draw.Src)
	return resImg, nil
}

// Simplify the process
func CropImage(src, dst string, x, y, w, h int) error {
	data, err := ImageRead(src)
	if err != nil {
		return fmt.Errorf("crop failed: %w", err)
	}
	img, err := Crop(data, x, y, w, h)
	if err != nil {
		return fmt.Errorf("crop failed: %w", err)
	}
	return ImageWrite(img, dst)
}
