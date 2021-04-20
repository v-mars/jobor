package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"ops-go/pkg/convert"
)

// 填充最后一个分组的函数
// src - 原始数据
// blockSize - 每个分组的数据长度
func padding(src []byte, blockSize int) []byte {
	// 1、求出最后一个分组要填充多少个字节
	padding := blockSize - len(src)%blockSize
	// 2、创建新的切片，切片的字节数为padding,并初始化，每个字节的值为padding
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 3、将创建出的新切片和原始数据进行链接
	newText := append(src, padText...)
	// 4、返回新的字符串
	return newText
}

//删除末尾填充的字节
func unPadding(src []byte) []byte {
	//1、 求出切片的长度
	length := len(src)
	// 2、取出最后一个字符，得到整型值
	number := int(src[length-1])
	// 3、将切片末尾的number个字节删除
	netText := src[:length-number]
	return netText
}

//使用des进行对称加密 length 8
func EncryptDES(src []byte, private []byte) []byte {
	//1、创建并返回一个使用DES算法的cipher.Block接口
	block, _ := des.NewCipher(private)
	//2、对最后一个明文分组进行数据填充
	src = padding(src, block.BlockSize())
	//3、创建一个密码分组为链接模式的，底层使用DES加密的BlockMode接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//4、加密连续的数据块
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

//使用des进行解密
func DecryptDES(src, private []byte) []byte {
	block, err := des.NewCipher(private)
	if err != nil {
		print(err)
	}
	//2、创建一个密码分组位链接模式的，底层使用DES解密的BlockMode接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//3、数据库解密
	blockMode.CryptBlocks(src, src)
	//4、去掉最后一组的填充数据
	newText := unPadding(src)
	return newText
}

//使用des进行对称加密 length 24
func Encrypt3DES(src []byte, private []byte) []byte {
	//1、创建并返回一个使用DES算法的cipher.Block接口
	block, _ := des.NewTripleDESCipher(private)
	//2、对最后一个明文分组进行数据填充
	src = padding(src, block.BlockSize())
	//3、创建一个密码分组为链接模式的，底层使用DES加密的BlockMode接口
	blockMode := cipher.NewCBCEncrypter(block, private[:block.BlockSize()])
	//4、加密连续的数据块
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

//使用des进行解密
func Decrypt3DES(src, private []byte) []byte {
	block, err := des.NewTripleDESCipher(private)
	if err != nil {
		print(err)
	}
	//2、创建一个密码分组位链接模式的，底层使用DES解密的BlockMode接口
	blockMode := cipher.NewCBCDecrypter(block, private[:block.BlockSize()])
	//3、数据库解密
	blockMode.CryptBlocks(src, src)
	//4、去掉最后一组的填充数据
	newText := unPadding(src)
	return newText
}

//使用AES加密 length 16
func EncryptAES(src, private []byte) []byte {
	//1、创建并返回一个使用AES算法的cipher.Block接口
	block, err := aes.NewCipher(private)
	if err != nil {
		panic(err)
	}
	//2、数据填充
	src = padding(src, block.BlockSize())
	//3、创建一个密码分组位链接模式的，底层使用AES解密的BlockMode接口
	blockMode := cipher.NewCBCEncrypter(block, private)
	//4、数据加密
	blockMode.CryptBlocks(src, src)
	return src
}

//使用AES解密
func DecryptAES(src, private []byte) []byte {
	//1、创建并返回一个使用AES算法的cipher.Block接口
	block, err := aes.NewCipher(private)
	if err != nil {
		panic(err)
	}
	//2、创建一个密码分组位链接模式的，底层使用AES解密的BlockMode接口
	blockMode := cipher.NewCBCDecrypter(block, private)
	//3、数据解密
	blockMode.CryptBlocks(src, src)
	//4、去掉填充数据
	src = unPadding(src)
	return src
}


//测试DES加解密
func desTest() {
	fmt.Println("======== des 加解密 =========")
	src := []byte("今天天气好晴朗")
	key := []byte("12345670")
	enc := EncryptDES(src, key)
	fmt.Println("加密后明文：", enc, convert.ToString(enc))
	dec := DecryptDES(enc, key)
	fmt.Println("解密后明文：", dec, string(dec))
}

//测试3DES加解密
func tdesTest() {
	fmt.Println("======== 3des 加解密 =========")
	src := []byte("root")
	//key := []byte("123456788765432112345678")
	key := []byte("12345678876543211aaaaaaa")
	enc := Encrypt3DES(src, key)
	fmt.Println("加密后明文：", enc, convert.ToString(enc))
	dec := Decrypt3DES(enc, key)
	fmt.Println("解密后明文：", dec, string(dec))
}

//测试AES加解密
func aesTest() {
	fmt.Println("======== aes 加解密 =========")
	src := []byte("password")
	key := []byte("12345678876543aa")
	//key := []byte("1234567887654321")
	enc := EncryptAES(src, key)
	fmt.Println("加密后明文：", enc, string(enc))
	dec := DecryptAES(enc, key)
	fmt.Println("解密后明文：", dec, string(dec))
}

func TestAll()  {
	desTest()
	tdesTest()
	aesTest()
}