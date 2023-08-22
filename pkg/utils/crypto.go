package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"errors"
	"fmt"
	"jobor/pkg/convert"

	"golang.org/x/crypto/bcrypt"
)

// PKCS7Padding PKCS7 填充模式
// 填充最后一个分组的函数
// ciphertext - 原始数据
// blockSize - 每个分组的数据长度
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	p := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padText := bytes.Repeat([]byte{byte(p)}, p)
	return append(ciphertext, padText...)
}

// PKCS7UnPadding 填充的反向操作，删除填充字符串
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}

//// 填充最后一个分组的函数
//// src - 原始数据
//// blockSize - 每个分组的数据长度
//func padding(src []byte, blockSize int) []byte {
//	// 1、求出最后一个分组要填充多少个字节
//	p := blockSize - len(src)%blockSize
//	// 2、创建新的切片，切片的字节数为padding,并初始化，每个字节的值为padding
//	padText := bytes.Repeat([]byte{byte(p)}, p)
//	// 3、将创建出的新切片和原始数据进行链接
//	newText := append(src, padText...)
//	// 4、返回新的字符串
//	return newText
//}
//
//// 删除末尾填充的字节
//func unPadding(src []byte) []byte {
//	//1、 求出切片的长度
//	length := len(src)
//	// 2、取出最后一个字符，得到整型值
//	number := int(src[length-1])
//	// 3、将切片末尾的number个字节删除
//	netText := src[:length-number]
//	return netText
//}

// EncryptDES 使用des进行对称加密 length 8
func EncryptDES(src []byte, private []byte) []byte {
	//1、创建并返回一个使用DES算法的cipher.Block接口
	block, _ := des.NewCipher(private)
	//2、对最后一个明文分组进行数据填充
	src = PKCS7Padding(src, block.BlockSize())
	//3、创建一个密码分组为链接模式的，底层使用DES加密的BlockMode接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//4、加密连续的数据块
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

// DecryptDES 使用des进行解密
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
	newText, _ := PKCS7UnPadding(src)
	return newText
}

// Encrypt3DES 使用des进行对称加密 length 24
func Encrypt3DES(src []byte, private []byte) []byte {
	//1、创建并返回一个使用DES算法的cipher.Block接口
	block, _ := des.NewTripleDESCipher(private)
	//2、对最后一个明文分组进行数据填充
	src = PKCS7Padding(src, block.BlockSize())
	//3、创建一个密码分组为链接模式的，底层使用DES加密的BlockMode接口
	blockMode := cipher.NewCBCEncrypter(block, private[:block.BlockSize()])
	//4、加密连续的数据块
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

// Decrypt3DES 使用des进行解密
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
	newText, _ := PKCS7UnPadding(src)
	return newText
}

//// EncryptAES 使用AES加密 length 16
//func EncryptAES(src, private []byte) []byte {
//	//1、创建并返回一个使用AES算法的cipher.Block接口
//	block, err := aes.NewCipher(private)
//	if err != nil {
//		panic(any(err))
//	}
//	//2、数据填充
//	src = PKCS7Padding(src, block.BlockSize())
//	//3、创建一个密码分组位链接模式的，底层使用AES解密的BlockMode接口
//	blockMode := cipher.NewCBCEncrypter(block, private)
//	//4、数据加密
//	blockMode.CryptBlocks(src, src)
//	return src
//}
//
//// DecryptAES 使用AES解密
//func DecryptAES(src, private []byte) []byte {
//	//1、创建并返回一个使用AES算法的cipher.Block接口
//	block, err := aes.NewCipher(private)
//	if err != nil {
//		panic(any(err))
//	}
//	//2、创建一个密码分组位链接模式的，底层使用AES解密的BlockMode接口
//	blockMode := cipher.NewCBCDecrypter(block, private)
//	//3、数据解密
//	blockMode.CryptBlocks(src, src)
//	//4、去掉填充数据
//	src, _ = PKCS7UnPadding(src)
//	return src
//}

// EncBase64ByAes Aes加密 后 base64 再加密
func EncBase64ByAes(data string, keyText string) string {
	//res := EncryptAES([]byte(data), []byte(keyText))
	res, _ := AesEncrypt([]byte(data), []byte(keyText))
	return Base64Enc(res)
}

// DecBase64ByAes Aes + base64 解密
func DecBase64ByAes(data string, keyText string) []byte {
	//return DecryptAES([]byte(Base64Dec(data)), []byte(keyText))
	decrypt, _ := AesDecrypt([]byte(Base64Dec(data)), []byte(keyText))
	return []byte(decrypt)
}

func AesDemo() {
	//plain := "The text need to be encrypt."
	//// AES 规定有3种长度的key: 16, 24, 32分别对应AES-128, AES-192, or AES-256
	//key := "AY3b5Z78806543Za"
	////key := "abcdefgehjhijkmlkjjwwoew"
	//// 加密
	//cipherByte := EncryptAES([]byte(plain), []byte(key))
	//
	//fmt.Printf("%s ==> %x\n", plain, cipherByte)
	//fmt.Printf("%s ==> %s\n", plain, Base64Enc(cipherByte))
	//// 解密
	//plainText := EncBase64ByAes(Base64Dec(Base64Enc(cipherByte)), key)
	//
	//fmt.Printf("%x ==> %s\n", cipherByte, plainText)
}

func AesDemo2() {
	//plain := "The text need to be encrypt."
	// AES 规定有3种长度的key: 16, 24, 32分别对应AES-128, AES-192, or AES-256
	//key := "385f33cb91484b04a177828829081ab7"
	var encStr = "63bc5fd2ee8ff296355efa56152e2e71"
	//key := "abcdefgehjhijkmlkjjwwoew"
	// 加密
	//cipherByte := EncryptAES([]byte(plain), []byte(key))
	//
	//fmt.Printf("%s ==> %x\n", plain, cipherByte)
	fmt.Printf(" %s\n", DeTxtByAes(encStr, "lInkbook1qazEWSP"))
	// 解密
	//plainText := DecBase64ByAes(encStr, key)
	//
	//fmt.Printf("%s\n", plainText)
}

// 测试DES加解密
func desTest() {
	fmt.Println("======== des 加解密 =========")
	src := []byte("今天天气好晴朗")
	key := []byte("12345670")
	enc := EncryptDES(src, key)
	fmt.Println("加密后明文：", enc, convert.ToString(enc))
	dec := DecryptDES(enc, key)
	fmt.Println("解密后明文：", dec, string(dec))
}

// 测试3DES加解密
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
//func aesTest() {
//	fmt.Println("======== aes 加解密 =========")
//	//src := []byte("password")
//	src := []byte("sPuDScvgKxoWV0jBx6OR5bYmTlhqf9")
//	key := []byte("12345678876543aa")
//	//key := []byte("1234567887654321")
//	enc := EncryptAES(src, key)
//	fmt.Println("加密后明文：", string(enc))
//	//dec := DecryptAES([]byte("æ,A�+����g���K1��/�Y�D�8�"), key)
//	dec := DecryptAES(enc, key)
//	fmt.Println("解密后明文：", dec, string(dec))
//}

func TestAll() {
	desTest()
	//tdesTest()
	//aesTest()
}

// HashAndSalt Hash & Salt 用户的密码
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		panic(any(err))
	}
	// 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	// 说明：每次运行，计算的密码值都不同。因此使用GoLang golang.org/x/crypto/bcrypt 模块对密码进行处理，可以避免字典攻击。
	return string(hash)
}

// ValidateSaltPasswords 验证密码
// 参数 hashedPwd: Hash & Salt 用户的密码，plainPwd: 明文密码
func ValidateSaltPasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

// ################# golang and js AES enc/dec ###################
//参考文档
//http://www.topgoer.com/%E5%85%B6%E4%BB%96/%E5%8A%A0%E5%AF%86%E8%A7%A3%E5%AF%86/%E5%8A%A0%E5%AF%86%E8%A7%A3%E5%AF%86.html
//高级加密标准（Adevanced Encryption Standard ,AES）

// DefaultTxtKey
// 16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法
// key不能泄露
// var PwdKey = []byte("DIS**#KKKDJJSKDI")
var DefaultTxtKey = "AirkJobor1qaz*xx"

// AesEncrypt 实现加密
func AesEncrypt(origData []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = PKCS7Padding(origData, blockSize)
	//采用AES加密方法中CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	//执行加密
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AesDecrypt 实现解密
func AesDecrypt(cypted []byte, key []byte) (string, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	//blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cypted))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(origData, cypted)
	//去除填充字符串
	origData, err = PKCS7UnPadding(origData)
	if err != nil {
		return "", err
	}
	return string(origData), err
}

// EnTxtByAes 加密base64
func EnTxtByAes(pwdStr, PwdKey string) string {
	if len(PwdKey) == 0 {
		PwdKey = DefaultTxtKey
	}
	pwd := []byte(pwdStr)
	result, err := AesEncrypt(pwd, []byte(PwdKey))
	//fmt.Println("EnTxtByAes:", err)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(result)
}

// DeTxtByAes 解密base64 aes
func DeTxtByAes(pwd, PwdKey string) string {
	if len(PwdKey) == 0 {
		PwdKey = DefaultTxtKey
	}
	temp, _ := hex.DecodeString(pwd)
	//执行AES解密
	res, _ := AesDecrypt(temp, []byte(PwdKey))
	return res
}

/* JS
import CryptoJS from "crypto-js";
const defaultSettings = require('../settings.js')
// var key1 = "1234567887654321";
let key1 = "lInkbook1qazEWSP";
// var plaintText = '"name"="lisi",age=18'; // 明文
// let str = {
//   name: "abcname",
//   site: "http://www.abc.com"
// }
// let plaintText = JSON.stringify(str)
// console.log(plaintText)
// let endata = encodeAes(plaintText)
//加密
export function encodeAes(plaintTextStr) {
  let key = CryptoJS.enc.Utf8.parse(Key);

  let encryptedData = CryptoJS.AES.encrypt(plaintTextStr, key, {
    iv: key,
    mode: CryptoJS.mode.CBC,
    padding: CryptoJS.pad.Pkcs7
  });

  // console.log("加密前：" + plaintText);
  // console.log("加密后：" + encryptedData);    //Pkcs7:   WoCzvm6eZiM4/bx5o/CzGw==

  // console.log("加密后 base64：" + encryptedData.ciphertext.toString(CryptoJS.enc.Base64));
  encryptedData = encryptedData.ciphertext.toString();
  // encryptedData = encryptedData.ciphertext.toString(CryptoJS.enc.Base64);
  // console.log("加密后-no-hex：" + encryptedData);
  return encryptedData
}

// 解密
// endata1 = "46ce4f5bb33896c4c75a24a46c6f16c32991228f40831003b98acffe41fee255f892d68283b8a1b07a4dfd66622b6c50685854e918ac059d5d8e969b3b105c6b";
// decodeAes(endata1)
export function decodeAes(encryptedDataStr) {
  let key = CryptoJS.enc.Utf8.parse(Key);
  let encryptedHexStr = CryptoJS.enc.Hex.parse(encryptedDataStr);
  // console.log("解密前hex：" + encryptedHexStr);
  let encryptedBase64Str = CryptoJS.enc.Base64.stringify(encryptedHexStr);
  // console.log("解密前：" + encryptedBase64Str);

  let decryptedData = CryptoJS.AES.decrypt(encryptedBase64Str, key, {
    iv: key,
    mode: CryptoJS.mode.CBC,
    padding: CryptoJS.pad.Pkcs7
  });

  let decryptedStr = decryptedData.toString(CryptoJS.enc.Utf8);
  // console.log("解密后:" + decryptedStr);
  return decryptedStr
}

*/
