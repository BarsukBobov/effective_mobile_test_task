package baseApi

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"strconv"
	"strings"
)

func GetPathID(c *gin.Context) (int, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.New("id должен быть числом!")
	}
	return id, err
}

func ValidatePhoto(photoFile *multipart.FileHeader) (string, error) {
	var errMessage string
	filename := photoFile.Filename
	// Проверка расширения файла
	if !strings.HasSuffix(filename, ".jpg") &&
		!strings.HasSuffix(filename, ".jpeg") &&
		!strings.HasSuffix(filename, ".png") {
		return "", errors.New("файл не является изображением")
	}
	// Попытка прочитать изображение
	file, err := photoFile.Open()
	if err != nil {
		errMessage = fmt.Sprintf("Ошибка при открытии файла:%s", err)
		err = errors.New(errMessage)
		logger.Error(err.Error())
		return "", err
	}
	defer file.Close()

	// Попытка декодировать изображение
	format := "jpeg"
	_, err = jpeg.Decode(file)
	if err != nil {
		file.Seek(0, 0)
		format = "png"
		_, err = png.DecodeConfig(file)
		if err != nil {
			errMessage = fmt.Sprintf("Ошибка при декодировании изображения:%s", err)
			err = errors.New(errMessage)
			logger.Error(err.Error())
			return "", err
		}
	}
	return format, nil
}
