package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"image/png"
	"jobor/internal/app/response"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func VerifyCode(c *gin.Context)  {
	//var param configJsonBody
	//generateCaptchaHandler(param)
	a,b:=GetCaptcha()
	fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(Verify(a,""))
	//fmt.Println(code)
	c.Writer.Header().Set("Content-TypeV1", "image/png; charset=utf-8")
	pngI, _ := png.Decode(bytes.NewReader([]byte(b)))
	pngI.At(80,240)
	if err := png.Encode(c.Writer, pngI);err!=nil{
		response.Error(c, err)
		return
	}
	//_, _ = c.Writer.WriteString(b)
	//response.Success(c, "")
}


func GenValidateCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
//fmt.Println( GenValidateCode(6) )


//configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// base64Captcha create http handler
func generateCaptchaHandler(param configJsonBody) {
	//parse request parameters
	//decoder := json.NewDecoder(r.Body)
	//var param configJsonBody
	//err := decoder.Decode(&param)
	//if err != nil {
	//	log.Println(err)
	//}
	//defer r.Body.Close()
	var driver base64Captcha.Driver
	param.CaptchaType = "string"
	//create base64 encoding captcha
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, store)
	fmt.Printf("body3: %+v\n", c)
	id, b64s, err := c.Generate()
	fmt.Println("body5:", err)
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	fmt.Println("body:", body)
	//w.Header().Set("Content-TypeV1", "application/json; charset=utf-8")
	//_ = json.NewEncoder(w).Encode(body)
}

//  获取验证码
func GetCaptcha()(string, string){
	// 生成默认数字
	//driver := base64Captcha.DefaultDriverDigit
	//base64Captcha.NewDriverDigit(height , width , length , maxSkew , dotCount)
	// 生成base64图片
	//c := base64Captcha.NewCaptcha(driver, store)


	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:          30,
		Width:           60,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890", // qwertyuioplkjhgfdsazxcvbnm
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	// 自定义配置，如果不需要自定义配置，则上面的结构体和下面这行代码不用写
	driverString = captchaConfig
	driver = driverString.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, store)
	//id, b64s, err := captcha.Generate()

	// 获取
	id, b64s, err := c.Generate()
	if err != nil {
		fmt.Println("Register GetCaptchaPhoto get base64Captcha has err:", err)
		return "",""
	}
	return id, b64s
}

// 验证验证码
func Verify(id string, val string) bool {
	if id == ""{
		return false
	}
	// 同时在内存清理掉这个图片
	return store.Verify(id, val, true)
}


// base64Captcha verify http handler
func captchaVerifyHandle(w http.ResponseWriter, r *http.Request) {

	//parse request json body
	decoder := json.NewDecoder(r.Body)
	var param configJsonBody
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	//verify the captcha
	body := map[string]interface{}{"code": 0, "msg": "failed"}
	if store.Verify(param.Id, param.VerifyValue, true) {
		body = map[string]interface{}{"code": 1, "msg": "ok"}
	}
	//set json response
	w.Header().Set("Content-TypeV1", "application/json; charset=utf-8")

	_ = json.NewEncoder(w).Encode(body)
}



