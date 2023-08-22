package utils

//
//import (
//	"encoding/json"
//
//	"github.com/asaskevich/govalidator"
//)
//
//// ValidateStruct struct covert validate struct and validate
//func ValidateStruct(source interface{}, out interface{}) (bool, error) {
//	err := StructToStruct(source, out)
//	if err != nil {
//		return false, err
//	}
//	result, err := govalidator.ValidateStruct(out)
//	if err != nil {
//		return result, err
//	}
//	return true, err
//}
//
//// StructToStruct struct covert validate struct
//func StructToStruct(source interface{}, out interface{}) error {
//	byteJson, err := json.Marshal(source)
//	if err != nil {
//		return err
//	}
//	err = json.Unmarshal(byteJson, out)
//	if err != nil {
//		return err
//	}
//	return err
//}
//
//func StructValid() {
//	//govalidator.SetFieldsRequiredByDefault(true)
//	type User struct {
//		Name  string      `valid:"type(string)"`
//		Age   int         `valid:"type(int)"`
//		Meta  interface{} `valid:"type(string),required"`
//		Name1 string      `valid:"-"`
//		Email string      `valid:"email,optional"`
//	}
//	result, err := govalidator.ValidateStruct(User{Name: "Bob", Age: 20, Meta: 123})
//	if err != nil {
//		println("error: " + err.Error())
//	}
//	println(result)
//}
//
//func MapValid() {
//	//govalidator.SetFieldsRequiredByDefault(true)
//	var mapTemplate = map[string]interface{}{
//		"name":       "required,alpha",
//		"family":     "required,alpha",
//		"email":      "required,email",
//		"cell-phone": "numeric",
//		"string1":    "type(string),required",
//		"address": map[string]interface{}{
//			"line1":       "required,alphanum",
//			"line2":       "alphanum",
//			"postal-code": "numeric",
//		},
//	}
//
//	var inputMap = map[string]interface{}{
//		"name":    "Bob",
//		"family":  "Smith",
//		"email":   "foo@bar.baz",
//		"string1": "12134",
//		"address": map[string]interface{}{
//			"line1":       "123",
//			"line2":       "",
//			"postal-code": "",
//		},
//	}
//
//	result, err := govalidator.ValidateMap(inputMap, mapTemplate)
//	if err != nil {
//		println("error: " + err.Error())
//	}
//	println(result)
//}
//
//// https://github.com/asaskevich/govalidator
//func aa() {
//	// before
//	//govalidator.CustomTypeTagMap["customByteArrayValidator"] = func(i interface{}, o interface{}) bool {
//	//	// ...
//	//	return true
//	//}
//	//
//	//// after
//	//govalidator.CustomTypeTagMap.Set("customByteArrayValidator", func(i interface{}, o interface{}) bool {
//	//	// ...
//	//	return true
//	//})
//}
//
//func bb() {
//	data := []interface{}{1, 2, 3, 4, 5}
//	var fn govalidator.Iterator = func(value interface{}, index int) {
//		println(value.(int))
//	}
//	govalidator.Each(data, fn)
//
//	data1 := []interface{}{1, 2, 3, 4, 5}
//	var fn1 govalidator.ResultIterator = func(value interface{}, index int) interface{} {
//		return value.(int) * 3
//	}
//	_ = govalidator.Map(data1, fn1) // result = []interface{}{1, 6, 9, 12, 15}
//
//	data2 := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	var fn2 govalidator.ConditionIterator = func(value interface{}, index int) bool {
//		return value.(int)%2 == 0
//	}
//	_ = govalidator.Filter(data2, fn2) // result = []interface{}{2, 4, 6, 8, 10}
//	_ = govalidator.Count(data2, fn2)  // result = 5
//}
//
//// Here is a list of available validators for struct fields (validator - used function):
///*
//"email":              IsEmail,
//"url":                IsURL,
//"dialstring":         IsDialString,
//"requrl":             IsRequestURL,
//"requri":             IsRequestURI,
//"alpha":              IsAlpha,
//"utfletter":          IsUTFLetter,
//"alphanum":           IsAlphanumeric,
//"utfletternum":       IsUTFLetterNumeric,
//"numeric":            IsNumeric,
//"utfnumeric":         IsUTFNumeric,
//"utfdigit":           IsUTFDigit,
//"hexadecimal":        IsHexadecimal,
//"hexcolor":           IsHexcolor,
//"rgbcolor":           IsRGBcolor,
//"lowercase":          IsLowerCase,
//"uppercase":          IsUpperCase,
//"int":                IsInt,
//"float":              IsFloat,
//"null":               IsNull,
//"uuid":               IsUUID,
//"uuidv3":             IsUUIDv3,
//"uuidv4":             IsUUIDv4,
//"uuidv5":             IsUUIDv5,
//"creditcard":         IsCreditCard,
//"isbn10":             IsISBN10,
//"isbn13":             IsISBN13,
//"json":               IsJSON,
//"multibyte":          IsMultibyte,
//"ascii":              IsASCII,
//"printableascii":     IsPrintableASCII,
//"fullwidth":          IsFullWidth,
//"halfwidth":          IsHalfWidth,
//"variablewidth":      IsVariableWidth,
//"base64":             IsBase64,
//"datauri":            IsDataURI,
//"ip":                 IsIP,
//"port":               IsPort,
//"ipv4":               IsIPv4,
//"ipv6":               IsIPv6,
//"dns":                IsDNSName,
//"host":               IsHost,
//"mac":                IsMAC,
//"latitude":           IsLatitude,
//"longitude":          IsLongitude,
//"ssn":                IsSSN,
//"semver":             IsSemver,
//"rfc3339":            IsRFC3339,
//"rfc3339WithoutZone": IsRFC3339WithoutZone,
//"ISO3166Alpha2":      IsISO3166Alpha2,
//"ISO3166Alpha3":      IsISO3166Alpha3,
//*/
//
////Validators with parameters
///*
//"range(min|max)": Range,
//"length(min|max)": ByteLength,
//"runelength(min|max)": RuneLength,
//"stringlength(min|max)": StringLength,
//"matches(pattern)": StringMatches,
//"in(string1|string2|...|stringN)": IsIn,
//"rsapub(keylength)" : IsRsaPub,
//"minstringlength(int): MinStringLength,
//"maxstringlength(int): MaxStringLength,
//*/
