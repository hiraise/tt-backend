package request

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"task-trail/internal/usecase/dto"
)

type FileReq struct {
	Body     []byte
	Name     string
	MimeType string
}

func FileFromAPI(file *multipart.FileHeader) (*FileReq, error) {
	retVal := &FileReq{}
	retVal.MimeType = file.Header.Get("Content-Type")
	if retVal.MimeType == "" {
		return nil, fmt.Errorf("mime-type is undefinded")
	}
	retVal.Name = filepath.Base(file.Filename)

	buf := bytes.NewBuffer(nil)
	f, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("file reading failure: %w", err)
	}
	defer f.Close()

	if _, err = io.Copy(buf, f); err != nil {
		return nil, fmt.Errorf("file copying failure: %w", err)
	}
	if buf.Len() == 0 {
		return nil, fmt.Errorf("file corrupted, file length is 0")
	}
	retVal.Body = buf.Bytes()
	return retVal, nil

}

func (r *FileReq) ToDTO(userID int) *dto.UploadFile {
	return &dto.UploadFile{UserID: userID, File: &dto.File{Data: r.Body, Name: r.Name, MimeType: r.MimeType}}
}
