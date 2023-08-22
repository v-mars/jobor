package utils

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//fmt.Println(MD5HashString16("127.0.0.1:2889"))
	//fmt.Println(MD5HashString16("127.0.0.1:2888"))
	//fmt.Println(MD5HashString16("127.0.0.1:2887"))
	//fmt.Println(len(MD5HashString16("1.1.1.1")))
	// admin $2a$10$lWjGqds.mVNxjuDHqsyJSeBjmEu75Hws8yDgQwGmwMFxHY8QuCkVi
	fmt.Println(HashAndSalt([]byte("admin")))
	fmt.Println(HashAndSalt([]byte("admin")))
	fmt.Println(ValidateSaltPasswords("$2a$10$O9GeHneK.st.qXkrZhX5auwWUxbJ7hWOJSN7zQsD.nhlQSECYYYQG",
		[]byte("admin")))
	// $2a$10$LBwJ2MIa8erQq7k6dkqvc.Wv1vLH9FsjYLu4VkMibcse8J6rvS4s6
	// $2a$10$iHz9anq5x2d8XikLCnMlWu5HJjEax2OUGdjdzRPwFLIXdiqPgqwe6
	// $2a$10$TB/VtGfygXrNUX0uszIjBurGZRNcQFjj2lyn4o.2NzcxEUy0f6cTm
}
