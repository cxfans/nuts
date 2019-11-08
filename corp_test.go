package nuts

import (
	"testing"
)

// go test corp_test.go corp.go io.go
func TestCrop(t *testing.T) {
	data, _ := ImageRead("imgs/1.jpg")
	img, err := Crop(data, 380, 640, 480, 240)
	if err != nil {
		t.Error(err)
	}
	if ImageWrite(img, "crop/1.jpg") != nil {
		t.Error(err)
	}
}

// go test -run TestCropImage
func TestCropImage(t *testing.T) {
	err := CropImage("imgs/2.jpg", "crop/2.png",
		380, 640, 480, 240)
	if err != nil {
		t.Error(err)
	}
}
