package utils

import (
	"math/rand"
	"time"
)

type createCourierRequest struct {
	Title string `json:"title"`
	Phone string `json:"phone"`
}

func CreateRandomCourierRequest() *createCourierRequest {
	rand.Seed(time.Now().UnixNano())

	title := RandString(5)
	phone := "8" + RandString(10)

	return &createCourierRequest{
		Title: title,
		Phone: phone,
	}
}

func RandString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)

}

func RandomModuloTen() int32 {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)
	return int32(randomNumber % 10)
}
