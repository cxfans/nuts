package nuts

import (
	"fmt"
	"testing"
)

// go test -run TestBase64Exporter
func TestBase64Exporter(t *testing.T) {
	img, err := ImageRead("crop/1.jpg")
	if err != nil {
		t.Error(err)
	}
	exporter := NewBase64Exporter(".jpg", img, true)
	if base64Str, err := exporter.Export(); err != nil {
		t.Error(err)
	} else {
		fmt.Println(base64Str)
	}
}

// go test -run TestImageExportBase64
func TestImageExportBase64(t *testing.T) {
	base64Str, err := ImageExportBase64("crop/2.png", false)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(base64Str)
}
