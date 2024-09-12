package utils

import (
	"mime/multipart"
	"strings"

)


func CheckTypeFile(file *multipart.FileHeader, extensions []string) (status bool) {
	for _, item := range extensions {
		if fileExt := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, ".")+1:]); fileExt == item {
			return true
		}
	}

	return false
}

func CheckContainsInt(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
