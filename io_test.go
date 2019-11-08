package nuts

import (
	"testing"
)

// go test -run TestImage_IO
func TestImage_IO(t *testing.T) {
	obj := NewImageReader("imgs/1.jpg")
	if img, err := obj.Read(); err != nil {
		t.Error(err)
	} else {
		if err = NewImageWriter(img, "imgs/1_1.jpg").Write(); err != nil {
			t.Error(err)
		}
	}
}
