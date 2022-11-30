package queries

import (
	"context"

	"os"
	"time"

	"stray-dogs/app/models"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Media struct{}

func (*Media) FileUpload(file models.File) (string, error) {

	uploadUrl, err := ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func ImageUploadHelper(input interface{}) (string, error) {
	CLOUDINARY_API_KEY := os.Getenv("CLOUDINARY_API_KEY")
	CLOUDINARY_API_SECRET := os.Getenv("CLOUDINARY_API_SECRET")
	CLOUDINARY_NAME := os.Getenv("CLOUDINARY_NAME")
	CLOUDINARY_UPLOAD_FOLDER := os.Getenv("CLOUDINARY_UPLOAD_FOLDER")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(CLOUDINARY_NAME, CLOUDINARY_API_KEY, CLOUDINARY_API_SECRET)
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: CLOUDINARY_UPLOAD_FOLDER})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}
