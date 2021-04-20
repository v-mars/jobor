package utils

func AnyToAny(src interface{}, des interface{}) error {
	bytesData,err:=json.Marshal(src)
	if err!=nil{return err}
	err = json.Unmarshal(bytesData, des)
	return err
}
