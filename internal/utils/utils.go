package utils

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
)

func LoadPic(path string) (*pixel.PictureData, error) {
	ifile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer ifile.Close()
	img, _, err := image.Decode(ifile)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
