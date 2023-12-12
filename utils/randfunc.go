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

	title := randString(5)
	phone := "8" + randString(10)

	return &createCourierRequest{
		Title: title,
		Phone: phone,
	}
}

func randString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)

}
