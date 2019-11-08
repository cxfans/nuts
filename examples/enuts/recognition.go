package main

import (
	"fmt"
	"github.com/cxfans/nuts"
	"os"
	"time"
)

const (
	apiKey    = "vIvU0EiFywrceahSSXXQ1Lta"
	secretKey = "r4sFrU2G7UiToYZPiuTx5b20aQFYItk8"
	apiUrl    = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"

	CropDir = "crop/"
	DoneDir = "examples/enuts/done/"
	SaveFile = "examples/enuts/nuts.txt"
)

func main() {
	d, err := os.Open(CropDir)
	if err != nil {
		panic(err)
	}
	items, err := d.Readdir(-1)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(SaveFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	client := nuts.NewClient(apiKey, secretKey, apiUrl)
	for _, item := range items {
		src := CropDir + item.Name()
		dst := DoneDir + item.Name()

		if words, err := client.GetWordsFromImage(src); err == nil {
			_, _ = file.WriteString(item.Name() + "\n")
			for _, word := range words {
				_, _ = file.WriteString(word)
				_, _ = file.WriteString("\n")
			}
			_, _ = file.WriteString("\n")
			fmt.Println("finished: ", item.Name())
			// Move images that finished successfully
			err := os.Rename(src, dst)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}

		// Since Baidu does not guarantee concurrency, the request
		// will be lost if too fast. So wait 1 second each request.
		time.Sleep(time.Second)
	}
}
