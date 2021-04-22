package property

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/mapstructure"
	"gorm.io/gorm"
	"jobor/internal/models"
	"jobor/internal/models/db"
	"jobor/internal/models/tbs"
	"jobor/internal/response"
	"jobor/pkg/convert"
	"jobor/pkg/utils"
)

type IProperty interface {
	Query(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	QueryLDAP(c *gin.Context)
	CreateOrUpdateLDAP(c *gin.Context)
	QueryEmail(c *gin.Context)
	CreateOrUpdateEmail(c *gin.Context)
	QueryAliYun(c *gin.Context)
	CreateOrUpdateAliYun(c *gin.Context)
	QueryGitlab(c *gin.Context)
	CreateOrUpdateGitlab(c *gin.Context)
	QueryJenkins(c *gin.Context)
	CreateOrUpdateJenkins(c *gin.Context)
}

type Property struct {
	DB *gorm.DB
}

func NewService(DB *gorm.DB) IProperty {
	return Property{DB: DB}
}

func (r Property) Query(c *gin.Context) {
	var obj []tbs.Property
	var pageData = response.InitPageData(c, &obj, false)
	type Param struct {
		Name string `form:"name"`
	}
	var param Param
	if err := c.ShouldBindQuery(&param);err!=nil{
		response.Error(c, err)
		return
	}
	var o models.Option
	o.Where = "name like ?"
	o.Value = append(o.Value, "%"+param.Name+"%")
	o.Select = "distinct id, name, k, v, ctime, mtime"
	o.Order = "ID DESC,name DESC"
	o.Scan = true

	err := models.Query(r.DB,&tbs.Property{}, o, &pageData)
	if err !=nil {
		response.Error(c, err)
		return
	}else {
		response.PageSuccess(c, pageData)
		return
	}
}

func (r Property) Create(c *gin.Context) {

}

func (r Property) Update(c *gin.Context) {

}

func (r Property) Delete(c *gin.Context) {
	var obj tbs.Property
	var rows map[string][]int
	if err := c.ShouldBindJSON(&rows); err!=nil{
		response.Error(c, err)
		return
	}
	if err:= r.DB.Where("id in (?)", rows["rows"]).Delete(&obj).Error;err!=nil{
		response.Error(c, err)
		return
	}
	response.DeleteSuccess(c)
}

// LDAP
func (r Property) QueryLDAP(c *gin.Context) {
	data,err := GetProperty("ldap")
	if err != nil{
		response.Error(c, err)
		return
	}
	if data["password"] != "" {
		data["password"] = utils.Base64Dec(convert.ToString(data["password"]))
	}
	response.Success(c, map[string]interface{}{"result": data})
	return
}

func (r Property) CreateOrUpdateLDAP(c *gin.Context) {
	type Param struct {
		//CheckEmail  string                 `json:"check_email"`
		Rows        map[string]interface{} `json:"row" binding:"required"`
	}
	var obj Param
	if err := c.ShouldBindQuery(&obj); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	obj.Rows["password"] = utils.Base64Enc([]byte(convert.ToString(obj.Rows["password"])))
	//fmt.Println("obj.Rows:", obj.Rows)
	if err:= CreateOrUpdate("ldap", obj.Rows); err!=nil{
		response.Error(c, err)
		return
	}
	response.SuccessMsg(c, "保存成功！", "")
}

// Email
type Email struct {
	Token  string `json:"token"`
	Host   string `json:"host"`
}
func (r Property) QueryEmail(c *gin.Context) {
	data,err := GetProperty("email")
	if err != nil{
		response.Error(c, err)
		return
	}
	if data["password"] != "" {
		data["password"] = utils.Base64Dec(convert.ToString(data["password"]))
	}
	response.Success(c, map[string]interface{}{"result": data})
}

func (r Property) CreateOrUpdateEmail(c *gin.Context) {
	type Param struct {
		//CheckEmail  string                 `json:"check_email"`
		Rows        map[string]interface{} `json:"row" binding:"required"`
	}
	var obj Param
	if err := c.ShouldBindQuery(&obj); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	obj.Rows["password"] = utils.Base64Enc([]byte(convert.ToString(obj.Rows["password"])))
	//fmt.Println("obj.Rows:", obj.Rows)
	if err:= CreateOrUpdate("email", obj.Rows); err!=nil{
		response.Error(c, err)
		return
	}
	response.SuccessMsg(c, "保存成功！", "")
}


// AliYun
type AliYun struct {
	Token  string `json:"token"`
	Host   string `json:"host"`
}
func (r Property) QueryAliYun(c *gin.Context) {
	data,err := GetProperty("aliyun")
	if err != nil{
		response.Error(c, err)
		return
	}
	if data["access_key_secret"] != "" {
		data["access_key_secret"] = utils.Base64Dec(convert.ToString(data["access_key_secret"]))
	}
	response.Success(c, map[string]interface{}{"result": data})
}

func (r Property) CreateOrUpdateAliYun(c *gin.Context) {
	type Param struct {
		CheckEmail  string                 `json:"check_email"`
		Rows        map[string]interface{} `json:"row" binding:"required"`
	}
	var obj Param
	if err := c.ShouldBindQuery(&obj); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	obj.Rows["access_key_secret"] = utils.Base64Enc([]byte(convert.ToString(obj.Rows["access_key_secret"])))
	//fmt.Println("obj.Rows:", obj.Rows)
	if err:= CreateOrUpdate("aliyun", obj.Rows); err!=nil{
		response.Error(c, err)
		return
	}
	response.SuccessMsg(c, "保存成功！", "")
}

// Gitlab
type Gitlab struct {
	Token  string `json:"token"`
	Host   string `json:"host"`
}
func (r Property) QueryGitlab(c *gin.Context) {
	data,err := GetProperty("gitlab")
	if err != nil{
		response.Error(c, err)
		return
	}
	if data["token"] != "" {
		data["token"] = utils.Base64Dec(convert.ToString(data["token"]))
	}
	response.Success(c, map[string]interface{}{"result": data})
}

func (r Property) CreateOrUpdateGitlab(c *gin.Context) {
	type Param struct {
		Rows        map[string]interface{} `json:"row" binding:"required"`
	}
	var obj Param
	if err := c.ShouldBindQuery(&obj); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	obj.Rows["token"] = utils.Base64Enc([]byte(convert.ToString(obj.Rows["token"])))
	//fmt.Println("obj.Rows:", obj.Rows)
	if err:= CreateOrUpdate("gitlab", obj.Rows); err!=nil{
		response.Error(c, err)
		return
	}
	response.SuccessMsg(c, "保存成功！", "")
}

// Jenkins
type Jenkins struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
}
func (r Property) QueryJenkins(c *gin.Context) {
	data,err := GetProperty("jenkins")
	if err != nil{
		response.Error(c, err)
		return
	}
	if data["password"] != "" {
		data["password"] = utils.Base64Dec(convert.ToString(data["password"]))
	}
	response.Success(c, map[string]interface{}{"result": data})
}

func (r Property) CreateOrUpdateJenkins(c *gin.Context) {
	type Param struct {
		Rows        map[string]interface{} `json:"row" binding:"required"`
	}
	var obj Param
	if err := c.ShouldBindQuery(&obj); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	obj.Rows["password"] = utils.Base64Enc([]byte(convert.ToString(obj.Rows["password"])))
	//fmt.Println("obj.Rows:", obj.Rows)
	if err:= CreateOrUpdate("jenkins", obj.Rows); err!=nil{
		response.Error(c, err)
		return
	}
	response.SuccessMsg(c, "保存成功！", "")
}

func GetJenkins() (Jenkins, error) {
	data,err := GetProperty("jenkins")
	var j Jenkins
	if err != nil{
		return j, err
	}
	if data["password"] != "" {
		data["password"] = utils.Base64Dec(convert.ToString(data["password"]))
	}
	err = mapstructure.Decode(data, &j)
	if err != nil {
		fmt.Println(err)
		return j, err
	}
	return j, err
}


func GetProperty(name string) (map[string]interface{}, error) {
	var o models.Option
	o.Where = "name = ?"
	o.Value = append(o.Value, name)
	o.Select = "distinct id, name, k, v, ctime, mtime"
	o.Order = "ID DESC"
	o.Scan = true
	var obj []tbs.Property
	var pageData = response.PageDataList{PageNumber: 1,PageSize:0,List:&obj}
	err := models.Query(db.DB,&tbs.Property{}, o, &pageData)
	if err!=nil{
		return nil, err
	}
	var data = map[string]interface{}{}
	for _, v := range obj {
		data[v.K] = v.V
		//fmt.Println("kv:",  v, v.K)
	}
	return data, nil
}

func CreateOrUpdate(name string,m map[string]interface{}) error {
	for k, v := range m {
		if err:= db.DB.Where(tbs.Property{Name: name, K: k}).Assign(
			tbs.Property{Name: name, K: k, V: convert.ToString(v)},
			).FirstOrCreate(&tbs.Property{}).Error; err!=nil{
				return err
		}
		//fmt.Println("update kv:", k, v)
	}
	return nil
}