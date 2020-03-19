package api

import (
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"mime/multipart"
	"os"
)

func CompressAndSaveFile(fileHeader *multipart.FileHeader, dest string) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	resized := resize.Resize(600, 600, img, resize.Lanczos3)

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	err = jpeg.Encode(out, resized, nil)

	return err
}
