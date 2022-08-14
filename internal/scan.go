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
	Path   string
	Images []*ImageInfo
}

type ImageInfo struct {
	Name   string
	Width  int
	Height int
}

func isImageFileName(fileName string) bool {
	reg := regexp.MustCompile(".(jpg|jpeg|png|webp|gif)$")
	return reg.Match([]byte(fileName))
}

func Scan(path string) (*ScanResult, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	images := []*ImageInfo{}

	fileInfoList, err := ioutil.ReadDir(path)
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
		imageFilePath := filepath.Join(path, imageName)
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
		Path:   path,
		Images: images,
	}

	return res, nil
}
