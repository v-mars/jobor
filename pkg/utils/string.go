package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
)

// S 字符串类型转换
type S string

func (s S) String() string {
	return string(s)
}

// Bytes 转换为[]byte
func (s S) Bytes() []byte {
	return []byte(s)
}

// Bool 转换为bool
func (s S) Bool() (bool, error) {
	b, err := strconv.ParseBool(s.String())
	if err != nil {
		return false, err
	}
	return b, nil
}

// DefaultBool 转换为bool，如果出现错误则使用默认值
func (s S) DefaultBool(defaultVal bool) bool {
	b, err := s.Bool()
	if err != nil {
		return defaultVal
	}
	return b
}

// Int64 转换为int64
func (s S) Int64() (int64, error) {
	i, err := strconv.ParseInt(s.String(), 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// DefaultInt64 转换为int64，如果出现错误则使用默认值
func (s S) DefaultInt64(defaultVal int64) int64 {
	i, err := s.Int64()
	if err != nil {
		return defaultVal
	}
	return i
}

// Int 转换为int
func (s S) Int() (int, error) {
	i, err := s.Int64()
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

// DefaultInt 转换为int，如果出现错误则使用默认值
func (s S) DefaultInt(defaultVal int) int {
	i, err := s.Int()
	if err != nil {
		return defaultVal
	}
	return i
}

// Uint64 转换为uint64
func (s S) Uint64() (uint64, error) {
	i, err := strconv.ParseUint(s.String(), 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// DefaultUint64 转换为uint64，如果出现错误则使用默认值
func (s S) DefaultUint64(defaultVal uint64) uint64 {
	i, err := s.Uint64()
	if err != nil {
		return defaultVal
	}
	return i
}

// Uint 转换为uint
func (s S) Uint() (int, error) {
	i, err := s.Uint64()
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

// DefaultUint 转换为uint，如果出现错误则使用默认值
func (s S) DefaultUint(defaultVal int) int {
	i, err := s.Uint()
	if err != nil {
		return defaultVal
	}
	return int(i)
}

// Float64 转换为float64
func (s S) Float64() (float64, error) {
	f, err := strconv.ParseFloat(s.String(), 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// DefaultFloat64 转换为float64，如果出现错误则使用默认值
func (s S) DefaultFloat64(defaultVal float64) float64 {
	f, err := s.Float64()
	if err != nil {
		return defaultVal
	}
	return f
}

// Float32 转换为float32
func (s S) Float32() (float32, error) {
	f, err := s.Float64()
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

// DefaultFloat32 转换为float32，如果出现错误则使用默认值
func (s S) DefaultFloat32(defaultVal float32) float32 {
	f, err := s.Float32()
	if err != nil {
		return defaultVal
	}
	return f
}

func (s S) JsonToString(v interface{}) error {
	return json.Unmarshal(s.Bytes(), v)
}

// Unmarshal ToStruct 转换为Struct Json Unmarshal：将json字符串解码到相应的数据结构
func Unmarshal(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}

// Marshal ToJSON 转换为JSON   Json Marshal：将数据编码成json字符串
func Marshal(v interface{}) (string, error) {
	JsonByte, err := json.Marshal(v)
	if err != nil {
		//fmt.Println("生成json字符串错误")
		panic(any("生成json字符串错误"))
		//return "", errors.New("生成json字符串错误")
	}
	JsonStr := string(JsonByte)
	return JsonStr, nil
}

func ParseInt(s string) int {
	if s == "" {
		return 0
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func ParseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}

	return f
}

func ParseBool(s string) bool {
	if s == "\x01" || s == "true" {
		return true
	} else if s == "false" {
		return false
	}

	i := ParseInt(s)
	return i != 0
}

func BoolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

// CamelToSnakeCase This function transform camelcase in snakecase LoremIpsum in lorem_ipsum
func CamelToSnakeCase(camel string) string {
	var buf bytes.Buffer
	for _, c := range camel {
		if 'A' <= c && c <= 'Z' {
			// just convert [A-Z] to _[a-z]
			if buf.Len() > 0 {
				buf.WriteRune('_')
			}
			buf.WriteRune(c - 'A' + 'a')
			continue
		}
		buf.WriteRune(c)
	}
	return strings.ReplaceAll(buf.String(), " ", "")
}

func GetOwnerAndNameFromId(id string) (string, string) {
	tokens := strings.Split(id, "/")
	if len(tokens) != 2 {
		panic(errors.New("GetOwnerAndNameFromId() error, wrong token count for Id: " + id))
	}

	return tokens[0], tokens[1]
}

func GetOwnerFromId(id string) string {
	tokens := strings.Split(id, "/")
	if len(tokens) != 2 {
		panic(errors.New("GetOwnerAndNameFromId() error, wrong token count for Id: " + id))
	}

	return tokens[0]
}

func GetOwnerAndNameFromIdNoCheck(id string) (string, string) {
	tokens := strings.SplitN(id, "/", 2)
	return tokens[0], tokens[1]
}

func GetOwnerAndNameAndOtherFromId(id string) (string, string, string) {
	tokens := strings.Split(id, "/")
	if len(tokens) != 3 {
		panic(errors.New("GetOwnerAndNameAndOtherFromId() error, wrong token count for Id: " + id))
	}

	return tokens[0], tokens[1], tokens[2]
}

func GenerateId() string {
	return uuid.NewString()
}

func GenerateTimeId() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	t := tm.Format("20060102_150405")

	random := uuid.NewString()[0:7]

	res := fmt.Sprintf("%s_%s", t, random)
	return res
}

func GenerateSimpleTimeId() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	t := tm.Format("20060102150405")

	return t
}

func GetId(owner, name string) string {
	return fmt.Sprintf("%s/%s", owner, name)
}

func GetSessionId(owner, name, application string) string {
	return fmt.Sprintf("%s/%s/%s", owner, name, application)
}

func GetMd5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func IsStringsEmpty(strs ...string) bool {
	for _, str := range strs {
		if len(str) == 0 {
			return true
		}
	}
	return false
}

func GetMaxLenStr(strs ...string) string {
	m := 0
	i := 0
	for j, str := range strs {
		l := len(str)
		if l > m {
			m = l
			i = j
		}
	}
	return strs[i]
}

func GetMinLenStr(strs ...string) string {
	m := int(^uint(0) >> 1)
	i := 0
	for j, str := range strs {
		l := len(str)
		if l < m {
			m = l
			i = j
		}
	}
	return strs[i]
}

func ReadStringFromPath(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func WriteStringToPath(s string, path string) {
	err := os.WriteFile(path, []byte(s), 0o644)
	if err != nil {
		panic(err)
	}
}

// SnakeString transform XxYy to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	result := strings.ToLower(string(data[:]))
	return strings.ReplaceAll(result, " ", "")
}

func IsChinese(str string) bool {
	var flag bool
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			flag = true
			break
		}
	}
	return flag
}

var rePhone *regexp.Regexp

func GetMaskedPhone(phone string) string {
	return rePhone.ReplaceAllString(phone, "$1****$2")
}

func GetMaskedEmail(email string) string {
	if email == "" {
		return ""
	}

	tokens := strings.Split(email, "@")
	username := maskString(tokens[0])
	domain := tokens[1]
	domainTokens := strings.Split(domain, ".")
	domainTokens[len(domainTokens)-2] = maskString(domainTokens[len(domainTokens)-2])
	return fmt.Sprintf("%s@%s", username, strings.Join(domainTokens, "."))
}

func maskString(str string) string {
	if len(str) <= 2 {
		return str
	} else {
		return fmt.Sprintf("%c%s%c", str[0], strings.Repeat("*", len(str)-2), str[len(str)-1])
	}
}

// GetEndPoint remove scheme from url
func GetEndPoint(endpoint string) string {
	for _, prefix := range []string{"https://", "http://"} {
		endpoint = strings.TrimPrefix(endpoint, prefix)
	}
	return endpoint
}
