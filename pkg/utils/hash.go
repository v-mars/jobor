package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

// MD5Hash MD5哈希值
func MD5Hash(b []byte) string {
	h := md5.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// MD5HashString MD5哈希值
func MD5HashString(s string) string {
	return MD5Hash([]byte(s))
}

// SHA1Hash SHA1哈希值
func SHA1Hash(b []byte) string {
	h := sha1.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA1HashString SHA1哈希值
func SHA1HashString(s string) string {
	return SHA1Hash([]byte(s))
}

// SHA1Hash SHA256哈希值
func SHA256Hash(b []byte) string {
	h := sha256.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA1HashString SHA1哈希值
func SHA256HashString(s string) string {
	return SHA256Hash([]byte(s))
}

type IHash interface {
	MD5HashString(s string) string
	SHA1HashString(s string) string
	SHA256HashString(s string) string
}

type Hash struct {}

func (h *Hash)SHA256HashString(s string) string {
	return SHA256Hash([]byte(s))
}

func (h *Hash)SHA1HashString(s string) string {
	return SHA1Hash([]byte(s))
}

// MD5HashString MD5哈希值
func (h *Hash)MD5HashString(s string) string {
	return MD5Hash([]byte(s))
}



