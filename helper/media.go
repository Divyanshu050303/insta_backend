package helper

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func UploadPost(userId string, c *fiber.Ctx) (url string, err error) {
	fmt.Printf("inside upload post function")
	cld, err := SetupCloudinary()
	if err != nil {
		ApiResponse(c, http.StatusBadRequest, "Connection is not setup with couldinary ", nil)
		return "", err
	}
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Printf("Error retrieving file: %v\n", err)
		ApiResponse(c, http.StatusBadRequest, "Failed to retrieve file", nil)
		return "", err
	}

	// Debugging log to confirm file retrieval
	fmt.Printf("File name: %s, Size: %d\n", file.Filename, file.Size)

	fileHeader, err := file.Open()
	if err != nil {
		ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		fmt.Printf("inside upload post function3")
		return "", err
	}

	postUrl, err := UploadMediaToCloudinary(userId, fileHeader, cld, c)
	if err != nil {
		ApiResponse(c, http.StatusBadRequest, "Bad Request", nil)
		return "", err
	}

	return postUrl, nil

}
func UploadMediaToCloudinary(userId string, file multipart.File, cld *cloudinary.Cloudinary, c *fiber.Ctx) (url string, err error) {
	err = godotenv.Load(".env")
	fmt.Printf("inside upload post function4")
	if err != nil {
		ApiResponse(c, http.StatusBadRequest, ".env is not loaded", nil)
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	bucketName := os.Getenv("CLOUD_BUCKET_NAME")

	uploadedFile, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: bucketName + "/" + userId, PublicID: userId})

	if err != nil {
		ApiResponse(c, http.StatusBadRequest, "file is not upoaded ", nil)
		fmt.Printf("inside upload post function7")
		return "", err
	}

	return uploadedFile.SecureURL, nil
}

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	godotenv.Load(".env")

	cldSecret := os.Getenv("CLOUD_SECRET_KEY")
	cldName := os.Getenv("CLOUD_NAME")
	cldKey := os.Getenv("CLOUD_API_KEY")

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		return nil, err
	}

	return cld, nil
}
