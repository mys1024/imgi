package internal

import (
	"bufio"
	"image"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"
)

type ScanResult struct {
	Dir    string       `json:"dir"`
	Images []*ImageInfo `json:"images"`
}

type ImageInfo struct {
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func isImageFileName(fileName string) bool {
	reg := regexp.MustCompile(".(jpg|jpeg|png|webp|gif)$")
	return reg.Match([]byte(fileName))
}

func Scan(dir string) (*ScanResult, error) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	images := []*ImageInfo{}

	fileInfoList, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range fileInfoList {
		if fileInfo.IsDir() {
			continue
		}
		if !isImageFileName(fileInfo.Name()) {
			continue
		}

		imageName := fileInfo.Name()
		imageFilePath := filepath.Join(dir, imageName)
		imageFile, err := os.OpenFile(imageFilePath, os.O_RDONLY, fileInfo.Mode().Perm())
		if err != nil {
			return nil, err
		}

		imageConfig, _, err := image.DecodeConfig(bufio.NewReader(imageFile))
		if err != nil {
			return nil, err
		}

		images = append(images, &ImageInfo{
			Name:   imageName,
			Width:  imageConfig.Width,
			Height: imageConfig.Height,
		})
	}

	res := &ScanResult{
		Dir:    dir,
		Images: images,
	}

	return res, nil
}
