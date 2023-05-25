package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/EdlanioJ/nlwpacetime/server/configs"
	"github.com/google/uuid"
	"github.com/uploadcare/uploadcare-go/ucare"
	"github.com/uploadcare/uploadcare-go/upload"
)

type UploadResponse struct {
	FileUrl string `json:"fileUrl"`
}

func Upload(file *multipart.FileHeader) (UploadResponse, error) {
	fileData, err := file.Open()
	if err != nil {
		return UploadResponse{}, fmt.Errorf("OpenFileError: %s", err.Error())
	}
	defer fileData.Close()

	ext := filepath.Ext(file.Filename)
	name := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	params := upload.FileParams{
		Data: fileData,
		Name: name,
	}

	cfg := configs.GetUploadcare()
	creds := ucare.APICreds{
		SecretKey: cfg.SecretKey,
		PublicKey: cfg.PublicKey,
	}

	conf := &ucare.Config{
		SignBasedAuthentication: true,
		APIVersion:              ucare.APIv06,
	}

	client, err := ucare.NewClient(creds, conf)
	if err != nil {
		return UploadResponse{}, fmt.Errorf("UploadClientError: %s", err.Error())
	}

	uploadSvc := upload.NewService(client)
	fID, err := uploadSvc.File(context.Background(), params)
	if err != nil {
		return UploadResponse{}, fmt.Errorf("UploadError: %s", err.Error())
	}
	fileUrl := fmt.Sprintf("https://ucarecdn.com/%s/", fID)
	return UploadResponse{
		FileUrl: fileUrl,
	}, nil
}
