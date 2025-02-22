package filesystem

import (
	"os"
	"path/filepath"
)

func SaveScreenshot() error {
	return nil
}

//func LoadScreenshot() {
//
//}

func ListScreenshots() ([]string, error) {
	var files []string

	err := filepath.Walk("data/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
