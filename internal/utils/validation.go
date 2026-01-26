package utils

import (
	"mime/multipart"
	"strings"
	"time"
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

func CheckFileSize(file *multipart.FileHeader, maxSizeMB int64) bool {
	const bytesInMB = 1024 * 1024
	return file.Size > maxSizeMB*bytesInMB
}

func IsToday(ts int64) bool {
	activityDate := time.UnixMilli(ts)
	now := time.Now()

	return activityDate.Year() == now.Year() &&
		activityDate.Month() == now.Month() &&
		activityDate.Day() == now.Day()
}

func BeforeDate(ts int64) bool {
	return time.UnixMilli(ts).Before(time.Now())
}

func AfterDate(ts int64) bool {
	return time.UnixMilli(ts).After(time.Now())
}
