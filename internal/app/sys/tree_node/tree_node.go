package tree_node

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jobor/internal/app"
	"jobor/internal/models"
	"jobor/internal/models/tbs"
	"jobor/internal/response"
	"jobor/pkg/convert"
)

type ITreeNode interface {
	app.CommonInterfaces
	Rename(c *gin.Context)
	GetAllMark(c *gin.Context)
}

type TreeNode struct {
	DB *gorm.DB
}

func NewService(DB *gorm.DB) ITreeNode {
	return TreeNode{DB: DB}
}

// @Tags 树形结构管理
// @Summary 树形结构Mark列表
// @Description 树形结构
// @Produce  json
// @Security ApiKeyAuth
// //@Param name query string false "Portal名"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/sys/tree_node_mark [get]
func (r TreeNode) GetAllMark(c *gin.Context) {
	var obj []ShowMarkList
	var o = r.Option()
	o.Select = "DISTINCT mark"
	o.Joins = ""
	o.Order = "mark"
	o.Scan = true
	tx :=r.DB.Begin()
	defer func() {
		tx.Rollback()
	}()
	err := models.Get(tx,&obj, o, &obj)
	if err != nil {
		response.Error(c, err)
		return
	} else {
		var pageData = response.PageDataList{PageNumber: 1,PageSize:0,List:&obj,Total: int64(len(obj))}
		response.Success(c, pageData)
	}
}

// @Tags 树形结构管理
// @Summary 树形结构详细
// @Description 树形结构
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "ID"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/sys/tree_node/{id} [get]
func (r TreeNode) Get(c *gin.Context) {
	_ = c.Params.ByName("id")
}

// @Tags 树形结构管理
// @Summary 树形结构列表
// @Description 树形结构
// @Produce  json
// @Security ApiKeyAuth
// @Param name query string false "树形结构名"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/sys/tree_node [get]
func (r TreeNode) Query(c *gin.Context) {
	var param QuerySchema
	if err := c.ShouldBindQuery(&param); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	//fmt.Println(param)
	if treeDataSlice,err:=r.QueryTreeNode(param); err!=nil{
		response.Error(c, err)
		return
	} else {
		response.Success(c, treeDataSlice)
	}
}

func listToTreeV1(data []map[string]interface{}) []map[string]interface{} {
	res := make(map[uint]map[string]interface{})
	for _, v := range data {
		id := uint(v["id"].(float64))
		parentID := uint(v["parent_id"].(float64))

		if res[id] != nil {
			v["children"] = res[id]["children"]
			res[id] = v
		} else {
			v["children"] = []map[string]interface{}{}
			res[id] = v
		}
		if res[parentID] != nil {
			res[parentID]["children"] = append(
				res[parentID]["children"].([]map[string]interface{}),
				res[id],
			)
		} else {
			res[parentID] = map[string]interface{}{
				"children": []map[string]interface{}{
					res[id],
				},
			}
		}
	}
	return res[0]["children"].([]map[string]interface{})
}


func listToTree(data []map[string]interface{}) []map[string]interface{} {
	res := make(map[int]map[string]interface{})
	for _, v := range data {
		id := v["id"].(int)
		parentID := v["parent_id"].(int)
		if res[id] != nil {
			v["children"] = res[id]["children"]
			res[id] = v
		} else {
			v["children"] = []map[string]interface{}{}
			res[id] = v
		}
		if res[parentID] != nil {
			res[parentID]["children"] = append(
				res[parentID]["children"].([]map[string]interface{}),
				res[id],
			)
		} else {
			res[parentID] = map[string]interface{}{
				"children": []map[string]interface{}{
					res[id],
				},
			}
		}
	}
	return res[0]["children"].([]map[string]interface{})
}

// @Tags 树形结构管理
// @Summary 创建树形结构node
// @Description 树形结构
// @Produce  json
// @Security ApiKeyAuth
// @Param payload body PostSchema true "参数信息"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/sys/tree_node [post]
func (r TreeNode) Create(c *gin.Context) {
	var obj PostSchema
	if err := c.ShouldBindJSON(&obj); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	if err:=r.AddNode(obj); err!=nil{
		response.Error(c, err)
		return
	} else {
		response.CreateSuccess(c, "")
	}
}

// @Tags 树形结构管理
// @Summary 移动树形结构Node
// @Description 树形结构
// @Produce  json
// @Security ApiKeyAuth
// //@Param id path int true "ID"
// @Param payload body PutSchema true "参数信息"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/sys/tree_node [put]
func (r TreeNode) Update(c *gin.Context) {
	var obj PutSchema
	if err := c.ShouldBindJSON(&obj); err!=nil{
		response.Error(c, err)
		return
	}
	if err:=r.MoveNode(obj); err!=nil{
		response.Error(c, err)
		return
	} else {
		response.UpdateSuccess(c, "")
	}
}

// @Tags 树形结构管理
// @Summary 移动树形结构Node
// @Description 树形结构
// @Produce  json
// @Security ApiKeyAuth
// //@Param id path int true "ID"
// @Param payload body PutSchema true "参数信息"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/sys/tree_node/{id}/rename [put]
func (r TreeNode) Rename(c *gin.Context) {
	_id := c.Params.ByName("id")
	var obj RenameSchema
	if err := c.ShouldBindJSON(&obj); err!=nil{
		response.Error(c, err)
		return
	}
	var NodeObj tbs.TreeNode
	if err:=r.DB.First(&NodeObj,"id = ?", _id).Error;err!=nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			err=errors.New(fmt.Sprintf("该节点[%s]不存在", _id))
		}
		response.Error(c, err)
		return
	}
	if err:=r.DB.Model(&NodeObj).Update("name", obj.Name).Error; err!=nil{
		response.Error(c, err)
		return
	} else {
		response.UpdateSuccess(c, "")
	}
}

// @Tags 树形结构管理
// @Summary 删除树形结构node
// @Description 树形结构
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "ID"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/sys/tree_node/{id}/{mark} [delete]
func (r TreeNode) Delete(c *gin.Context) {
	_id := c.Params.ByName("id")
	mark := c.Params.ByName("mark")
	if err:=r.DeleteNode(_id, mark); err!=nil{
		response.Error(c, err)
		return
	} else {
		response.DeleteSuccess(c)
	}

}


func (r TreeNode) Option() models.Option {
	var o models.Option
	o.Select = "id, parent_id, lft, rgt,name,mark"
	//o.Select = "*"
	o.Order = "parent_id,lft ASC"
	return o
}

func (r TreeNode) QueryTreeNode(param QuerySchema) ([]map[string]interface{},error){
	var obj []ShowSchema
	var o = r.Option()
	o.Where = "mark = ?"
	o.Value = append(o.Value, param.Mark)
	//o.Order = "parent_id,lft ASC"
	o.Scan = true
	dbObj:= models.DBInt(&tbs.TreeNode{},o, r.DB)
	if err:=dbObj.Scan(&obj).Error;err!=nil{
		return nil, err
	}
	var MapData []map[string]interface{}
	var err error
	if MapData,err=convert.StructToMapSlice(obj); err!=nil{
		return nil, err
	}
	//fmt.Println("MapData:", MapData)
	var treeDataSlice  = make([]map[string]interface{},0)
	if len(MapData)>0{
		treeDataSlice= listToTreeV1(MapData)
	}
	return treeDataSlice, nil
}

func (r TreeNode) AddNode(obj PostSchema) error{
	tx :=r.DB.Begin()
	defer func() {
		tx.Rollback()
	}()

	var sameCount int64 = 0
	var ParentObj tbs.TreeNode
	if err:=tx.First(&ParentObj,"id = ? and mark = ?", obj.ParentID,obj.Mark).Error;
	err!=nil{
		return err
	}  // && !errors.Is(err, gorm.ErrRecordNotFound)
	if err:=tx.Model(tbs.TreeNode{}).Select("id").Where("`parent_id` = ? and name = ? and mark=?",
		obj.ParentID,obj.Name,obj.Mark).Count(&sameCount).Error;err!=nil && !errors.Is(err, gorm.ErrRecordNotFound){  //&& !gorm.IsRecordNotFoundError(err)
		return err
	}

	if obj.ParentID == 0 {
		var rootCount int64 = 0
		if err:=tx.Model(&tbs.TreeNode{}).Select("id").Where("lft = ? and mark=?", 1,obj.Mark).Count(
			&rootCount).Error;err!=nil && !errors.Is(err, gorm.ErrRecordNotFound){
			return err
		}			//  && !errors.Is(err, gorm.ErrRecordNotFound)
		if rootCount>0{
			return errors.New("根节点已存在")
		}
		var newNode = tbs.TreeNode{ParentID: 0, Name: obj.Name, Lft: 1, Rgt: 2,Mark: obj.Mark}
		if err:=tx.Create(&newNode).Error;err!=nil{
			return err
		}
	}else if sameCount > 0{
		return errors.New("已存在同名兄弟节点")
	}else if ParentObj.ID != 0 {
		//fmt.Println("new......", ParentObj.Rgt, ParentObj.Lft, ParentObj.Name)
		if err:=tx.Model(&tbs.TreeNode{}).Where("rgt >= ? and mark=?",ParentObj.Rgt,obj.Mark).UpdateColumn(
			"rgt", gorm.Expr("rgt + ?", 2)).Error;err!=nil{
			return err
		}
		if err:=tx.Model(&tbs.TreeNode{}).Where("lft >= ? and mark=?",ParentObj.Rgt,obj.Mark).UpdateColumn(
			"lft", gorm.Expr("lft + ?", 2)).Error;err!=nil{
			return err
		}
		var right = ParentObj.Rgt+1
		var newNode = tbs.TreeNode{
			ParentID: obj.ParentID, Name: obj.Name, Lft: ParentObj.Rgt, Rgt: right, Mark: obj.Mark}
		if err:=tx.Create(&newNode).Error;err!=nil{
			return err
		}
		//if err:=r.DB.Exec("set right1=(select `right` from tree_v2 where `id` = %s);" +
		//	"update tree_v2 set `right` = `right` + 2 where `right` >= right1;" +
		//	"update tree_v2 set `left` = `left` + 2 where `left` >= right1;" +
		//	"insert into tree_v2(`parent_id`, `name`, `left`, `right`) values(%s,%s, right1, right1 + 1);",
		//	obj.ParentID,obj.ParentID,obj.Name).Error;err!=nil{
		//	response.Error(c, err)
		//	return
		//}
	} else {
		return errors.New("父节点不存在(未创建或被删除)")
	}

	if err:=tx.Commit().Error;err!=nil{
		tx.Rollback()
		return err
	}
	return nil
}

func (r TreeNode) MoveNode(obj PutSchema) error {
	tx :=r.DB.Begin()
	defer func() {
		tx.Rollback()
	}()
	var SelfObj tbs.TreeNode
	// 获取移动对象的 left, right 值
	if err:=tx.First(&SelfObj,"id=? and mark=?", obj.SelfID,obj.Mark).Error;err!=nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			err=errors.New(fmt.Sprintf("当前移动节点[%d]不存在", obj.SelfID))
		}
		return err
	}
	//fmt.Println("SelfObj:", SelfObj)
	// 将需要移动的记录的 id 存入临时表, 以保证操作 left, right 值变化时这些记录不受影响
	var moveNodeIdList []uint
	if err:=tx.Model(&tbs.TreeNode{}).Where(
		"lft>=? and rgt<=? and mark=?", SelfObj.Lft,SelfObj.Rgt,obj.Mark).Pluck(
			"id",&moveNodeIdList).Error; err!=nil{
		return err
	}
	//fmt.Println("moveNodeIdList:", moveNodeIdList, obj.SiblingID)

	// 将被移动记录后面的记录往前移, 填充空缺位置
	var num=SelfObj.Rgt-SelfObj.Lft+1
	/*
		update tree_v2 set `left` = `left` - (self_right-self_left+1) where `left` > self_left and id not in
		(select id from tree_v2_self_ids);

		update tree_v2 set `right` = `right` - (self_right-self_left+1) where `right` > self_right and id not in
		(select id from tree_v2_self_ids);
	*/

	if err:=tx.Exec("update tree_node set `lft`=`lft`-? where `lft` > ? and id not in (?) and mark = ?",
		num, SelfObj.Lft, moveNodeIdList,obj.Mark).Error;err!=nil{
		return err
	}

	if err:=tx.Exec("update tree_node set `rgt`=`rgt`-? where `rgt` > ? and id not in (?) and mark = ?",
		num, SelfObj.Rgt, moveNodeIdList,obj.Mark).Error;err!=nil{
		return err
	}

	var ParentObj tbs.TreeNode
	// 获取父节点对象并检查是否纯在
	if err:=tx.First(&ParentObj,"id = ? and mark = ?", obj.ParentID,obj.Mark).Error;err!=nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			err=errors.New(fmt.Sprintf("父节点[%d]不存在", obj.ParentID))
		}
		return err
	}


	if obj.SiblingID == -1{
		fmt.Println("在末尾插入子节点")
		// 在末尾插入子节点
		/*
			update tree_v2 set `right` = `right` + (self_right-self_left+1) where `right` >= parent_right and id not in
			(select id from tree_v2_self_ids);

			update tree_v2 set `left` = `left` + (self_right-self_left+1) where `left` >= parent_right and id not in
			(select id from tree_v2_self_ids);

			update tree_v2 set `right`=`right` + (parent_right-self_left), `left`=`left` + (parent_right-self_left)
			where id in (select id from tree_v2_self_ids);
		*/
		//if err:=tx.Model(&tbs.TreeNode{}).Where("rgt >= ? and id not in (?)",ParentObj.Rgt, moveNodeIdList).UpdateColumn(
		//	"rgt", gorm.Expr("rgt + ?", num)).Error;err!=nil{
		//	response.Error(c, err)
		//	return
		//}
		//
		//if err:=tx.Model(&tbs.TreeNode{}).Where("lft >= ? and id not in (?)",ParentObj.Rgt, moveNodeIdList).UpdateColumn(
		//	"lft", gorm.Expr("lft + ?", num)).Error;err!=nil{
		//	response.Error(c, err)
		//	return
		//}

		if err:=tx.Exec("update tree_node set `rgt`=`rgt`+? where `rgt` >= ? and id not in (?) and mark = ?",
			num, ParentObj.Rgt, moveNodeIdList,obj.Mark).Error;err!=nil{
			return err
		}

		if err:=tx.Exec("update tree_node set `lft`=`lft`+? where `lft` >= ? and id not in (?) and mark = ?",
			num, ParentObj.Rgt, moveNodeIdList,obj.Mark).Error;err!=nil{
			return err
		}
		var tmpNum= ParentObj.Rgt-SelfObj.Lft
		if err:=tx.Exec("update tree_node set `rgt`=`rgt`+?, `lft`=`lft`+? where id in (?) and mark = ?",
			tmpNum, tmpNum, moveNodeIdList,obj.Mark).Error;err!=nil{
			return err
		}
		//if err:=tx.Model(&tbs.TreeNode{}).Where("id in (?)",moveNodeIdList).UpdateColumn(
		//	"rgt", gorm.Expr("rgt + ?", ParentObj.Rgt-SelfObj.Lft)).UpdateColumn(
		//	"lft", gorm.Expr("lft + ?", ParentObj.Rgt-SelfObj.Lft)).Error;err!=nil{
		//	response.Error(c, err)
		//	return
		//}

	}else if obj.SiblingID == 0{
		fmt.Println("在开头插入子节点")
		// 在开头插入子节点
		/*
			update tree_v2 set `right` = `right` + (self_right-self_left+1) where `right` > parent_left and id not in
			(select id from tree_v2_self_ids);

			update tree_v2 set `left` = `left` + (self_right-self_left+1) where `left` > parent_left and id not in
			(select id from tree_v2_self_ids);

			update tree_v2 set `right`=`right` - (self_left-parent_left-1), `left`=`left` - (self_left-parent_left-1)
			where id in (select id from tree_v2_self_ids)
		*/
		//if err:=tx.Model(&tbs.TreeNode{}).Where("rgt > ? and id not in (?)",ParentObj.Lft, moveNodeIdList).UpdateColumn(
		//	"rgt", gorm.Expr("rgt + ?", num)).Error;err!=nil{
		//	response.Error(c, err)
		//	return
		//}
		//
		//if err:=tx.Model(&tbs.TreeNode{}).Where("lft > ? and id not in (?)",ParentObj.Lft, moveNodeIdList).UpdateColumn(
		//	"lft", gorm.Expr("lft + ?", num)).Error;err!=nil{
		//	response.Error(c, err)
		//	return
		//}
		if err:=tx.Exec("update tree_node set `rgt`=`rgt`+? where `rgt` > ? and id not in (?) and mark = ?",
			num, ParentObj.Lft, moveNodeIdList,obj.Mark).Error;err!=nil{
			return err
		}

		if err:=tx.Exec("update tree_node set `lft`=`lft`+? where `lft` > ? and id not in (?) and mark = ?",
			num, ParentObj.Lft, moveNodeIdList,obj.Mark).Error;err!=nil{
			return err
		}

		var tmpNum= SelfObj.Lft-ParentObj.Lft-1
		//fmt.Println("111", tmpNum, SelfObj.Lft, ParentObj.Lft, SelfObj.Lft-ParentObj.Lft-1)
		if err:=tx.Exec("update tree_node set `rgt`=`rgt`-?, `lft`=`lft`-? where id in (?) and mark = ?",
			tmpNum, tmpNum, moveNodeIdList,obj.Mark).Error;err!=nil{
			return err
		}
		//if err:=tx.Model(&tbs.TreeNode{}).Where("id in (?)",moveNodeIdList).UpdateColumn(
		//	"rgt", gorm.Expr("rgt - ?", SelfObj.Lft-ParentObj.Lft-1)).UpdateColumn(
		//	"lft", gorm.Expr("lft - ?", SelfObj.Lft-ParentObj.Lft-1)).Error;err!=nil{
		//	response.Error(c, err)
		//	return
		//}
	}else {
		// 插入指定节点之后
		/*
			select `left`, `right`, `parent_id` into sibling_left, sibling_right, sibling_parent_id from tree_v2 where id = sibling_id;
			if parent_id != sibling_parent_id then
				SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = '指定的兄弟节点不在指定的父节点中';
			end if;
			update tree_v2 set `right` = `right` + (self_right-self_left+1) where `right` > sibling_right and id not in
			(select id from ctree_v2_self_ids);

			update tree_v2 set `left` = `left` + (self_right-self_left+1) where `left` > sibling_right and id not in
			(select id from tree_v2_self_ids);

			update tree_v2 set `right`=`right` - (self_left-sibling_right-1), `left`=`left` - (self_left-sibling_right-1)
			where id in (select id from tree_v2_self_ids);
		*/
		var SiblingObj tbs.TreeNode
		if err:=tx.First(&SiblingObj,"id = ? and mark = ?", obj.SiblingID,obj.Mark).Error;err!=nil{
			if errors.Is(err, gorm.ErrRecordNotFound){
				err=errors.New(fmt.Sprintf("指定插入节点[%d]不存在", obj.ParentID))
			}
			return err
		}
		if obj.ParentID != SiblingObj.ParentID {
			return errors.New(fmt.Sprintf("指定的兄弟节点[%d]不在指定的父节点中", obj.SiblingID))
		}

		if err:=tx.Exec("update tree_node set `rgt`=`rgt`+? where `rgt` > ? and id not in (?) and mark = ?",
			num, SiblingObj.Rgt, moveNodeIdList,obj.Mark).Error;err!=nil{
			return err
		}

		if err:=tx.Exec("update tree_node set `lft`=`lft`+? where `lft` > ? and id not in (?) and mark = ?",
			num, SiblingObj.Rgt, moveNodeIdList,obj.Mark).Error;err!=nil{
			return err
		}

		var tmpNum=SelfObj.Lft-SiblingObj.Rgt-1
		if err:=tx.Exec("update tree_node set `rgt`=`rgt`-?, `lft`=`lft`-? where id in (?) and mark = ?",
			tmpNum, tmpNum, moveNodeIdList,obj.Mark).Error;err!=nil{
			return err
		}
	}

	SelfObj.ParentID = obj.ParentID
	if err:=tx.Model(&tbs.TreeNode{BaseID: tbs.BaseID{ID: SelfObj.ID}}).Updates(map[string]interface{}{"parent_id": obj.ParentID}).Error;err!=nil{
		return err
	}

	if err:=tx.Commit().Error;err!=nil{
		tx.Rollback()
		return err
	}
	return nil
}

func (r TreeNode) DeleteNode(_id string, mark string) error{
	tx :=r.DB.Begin()
	defer func() {
		tx.Rollback()
	}()
	var NodeObj tbs.TreeNode
	if err:=tx.First(&NodeObj,"id = ? and mark = ?", _id, mark).Error;err!=nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			err=errors.New(fmt.Sprintf("该节点[%s]不存在", _id))
		}
		return err
	}

	if err:= tx.Delete(&tbs.TreeNode{},"lft >= ? and rgt <= ?", NodeObj.Lft, NodeObj.Rgt).Error; err!=nil{
		return err
	}
	var num=NodeObj.Rgt-NodeObj.Lft+1
	if err:=tx.Model(&tbs.TreeNode{}).Where("lft > ?",NodeObj.Lft).UpdateColumn(
		"lft", gorm.Expr("lft - ?", num)).Error;err!=nil{
		return err
	}

	if err:=tx.Model(&tbs.TreeNode{}).Where("rgt > ?",NodeObj.Rgt).UpdateColumn(
		"rgt", gorm.Expr("rgt - ?", num)).Error;err!=nil{
		return err
	}

	if err:=tx.Commit().Error;err!=nil{
		tx.Rollback()
		return err
	}
	return nil
}

func (r TreeNode) GetChildAll(_id string, mark string) error{
	var obj []ShowSchema
	var o = r.Option()
	o.Where = "left >= (select left from tree_node where id=?) and " +
		"right <= (select right from tree_node where id=?) and mark = ?"
	o.Value = append(o.Value, _id,_id,mark)
	//o.Order = "parent_id,lft ASC"
	o.Scan = true
	dbObj:= models.DBInt(&tbs.TreeNode{},o, r.DB)
	if err:=dbObj.Scan(&obj).Error;err!=nil{
		return err
	}
	// 根据节点id获取此节点的所有子孙节点(包含自己)
	// select * from tree_table where
	// left >= (select left from tree_table where id=???) and
	// right <= (select right from tree_table where id=???)
	return nil
}

func (r TreeNode) GetChild(_id string, mark string) error{
	var obj []ShowSchema
	var o = r.Option()
	o.Where = "parent_id = ? and mark = ?"
	o.Value = append(o.Value,_id,mark)
	//o.Order = "parent_id,lft ASC"
	o.Scan = true
	dbObj:= models.DBInt(&tbs.TreeNode{},o, r.DB)
	if err:=dbObj.Scan(&obj).Error;err!=nil{
		return err
	}
	return nil
}

func (r TreeNode) GetParent(_id string, mark string) error{
	var obj []ShowSchema
	var o = r.Option()
	o.Where = "left <= (select left from tree_node where id=?) and " +
		"right >= (select right from tree_node where id=?) and mark = ?"
	o.Value = append(o.Value, _id,_id,mark)
	//o.Order = "parent_id,lft ASC"
	o.Scan = true
	dbObj:= models.DBInt(&tbs.TreeNode{},o, r.DB)
	if err:=dbObj.Scan(&obj).Error;err!=nil{
		return err
	}
	// 根据节点id获取此节点的所有上级节点(包括自己)
	// select * from tree_table where
	// left <= (select left from tree_table where id=???) and
	// right >= (select right from tree_table where id=???)
	return nil
}