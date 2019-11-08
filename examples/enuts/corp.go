package main

import (
	"fmt"
	"github.com/cxfans/nuts"
	"os"
)

const (
	ImagesDir = "imgs/"
	CropDir   = "crop/"
)

func main() {
	d, err := os.Open(ImagesDir)
	if err != nil {
		panic(err)
	}
	items, err := d.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		src := ImagesDir + item.Name()
		dst := CropDir + item.Name()
		err := nuts.CropImage(src, dst, 400, 650, 470, 205) // 1440x3040
		//err := nuts.CropImage(src, dst, 400, 500, 470, 205) // 1440x3040
		//err := nuts.CropImage(src, dst, 300, 430, 360, 175) // 1080x1920
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("finished!")
}
