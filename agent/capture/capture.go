package capture

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
	"time"
)

func CaptureScreen(hostname string) []string {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		_ = os.Mkdir("data", 0755)
	}

	n := screenshot.NumActiveDisplays()
	var filenames []string

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}

		currentTimestamp := time.Now().Format("2006-01-02_15-04-05")
		fileName := fmt.Sprintf("data/%s_%s_%dx%d.png", hostname, currentTimestamp, bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)
		filenames = append(filenames, fileName)
	}
	return filenames
}

func DeleteScreenshot(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}
