package formatter

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func (fm *Formatter) AddPhotoFromStorage(path string) {
	fm.Message.Photo = path
	fm.kindofmedia = fromStorage
	fm.mediatype = "photo"
}

func (fm *Formatter) AddPhotoFromTG(path string) {
	fm.Message.Photo = path
	fm.kindofmedia = fromTelegram
	fm.mediatype = "photo"
}

func (fm *Formatter) AddPhotoFromInternet(path string) {
	fm.Message.Photo = path
	fm.kindofmedia = fromInternet
	fm.mediatype = "photo"
}

func (fm *Formatter) AddVideoFromStorage(path string) {
	fm.Message.Video = path
	fm.kindofmedia = fromStorage
	fm.mediatype = "video"
}

func (fm *Formatter) AddVideoFromTG(path string) {
	fm.Message.Video = path
	fm.kindofmedia = fromTelegram
	fm.mediatype = "video"
}

func (fm *Formatter) AddVideoFromInternet(path string) {
	fm.Message.Video = path
	fm.kindofmedia = fromInternet
	fm.mediatype = "video"
}

func (fm *Formatter) PrepareMedia(buf *bytes.Buffer) (string, error) {
	var (
		file   *os.File
		part   io.Writer
		writer *multipart.Writer
		err    error
		media  string
	)

	writer = multipart.NewWriter(buf)
	if fm.mediatype == "photo" {
		media = fm.Message.Photo
	} else if fm.mediatype == "video" {
		media = fm.Message.Video
	}

	file, err = os.Open(media)
	if err == nil {
		part, err = writer.CreateFormFile(fm.mediatype, media)
	}
	if err == nil {
		_, err = io.Copy(part, file)
	}
	if err == nil {
		err = writer.WriteField("chat_id", fmt.Sprintf("%d", fm.Message.ChatID))
	}
	if err == nil {
		err = writer.WriteField("caption", fm.Message.Text)
	}
	if err == nil {
		err = writer.Close()
	}
	file.Close()

	return writer.FormDataContentType(), err

}
