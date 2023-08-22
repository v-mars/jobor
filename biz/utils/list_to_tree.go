package utils

func ListToTree(data []map[string]interface{}) map[int]map[string]interface{} {
	res := make(map[int]map[string]interface{})
	for _, v := range data {
		v := v
		name := v["name"].(string)
		id := int(v["id"].(float64))
		parentID := int(v["parent_id"].(float64))

		v["name_path"] = []string{name}
		v["path"] = []int{id}
		if res[id] != nil {
			v["children"] = res[id]["children"]
			res[id] = v
		} else {
			//res[id]["children"] = []map[string]interface{}{}
			//res[id]["children"] = make([]map[string]interface{}, 0)
			res[id] = v
		}
		if res[parentID] != nil {
			namePaths := res[parentID]["name_path"].([]string)
			paths := res[parentID]["path"].([]int)
			v["path"] = append([]int{}, paths...)
			v["path"] = append(v["path"].([]int), id)
			v["name_path"] = append([]string{}, namePaths...)
			v["name_path"] = append(v["name_path"].([]string), name)
			if res[parentID]["children"] != nil {
				res[parentID]["children"] = append(res[parentID]["children"].([]map[string]interface{}), res[id])
			} else {
				res[parentID]["children"] = []map[string]interface{}{res[id]}
			}
		} else {
			res[parentID] = map[string]interface{}{
				"path":      []int{},
				"name_path": []string{},
				"children": []map[string]interface{}{
					res[id],
				},
			}
			//if parentID > 0 {
			//	res[parentID]["path"] = []int{parentID}
			//}
		}
	}
	//res[0]["children"].([]map[string]interface{})
	return res
}

func GetChildMethod(childIds *[]int, children []map[string]interface{}) {
	for _, c := range children {
		c := c
		*childIds = append(*childIds, int(c["id"].(float64)))
		if c["children"] != nil && c["children"].([]map[string]interface{}) != nil {
			GetChildMethod(childIds, c["children"].([]map[string]interface{}))
		}
	}
}

func listToTree(data []map[string]interface{}) map[int]map[string]interface{} {
	res := make(map[int]map[string]interface{})
	for _, v := range data {
		id := int(v["id"].(float64))
		parentID := int(v["parent_id"].(float64))

		if res[id] != nil {
			v["children"] = res[id]["children"]
			res[id] = v
		} else {
			//v["children"] = []map[string]interface{}{}
			res[id] = v
		}
		if res[parentID] != nil {
			if res[parentID]["children"] != nil {
				res[parentID]["children"] = append(
					res[parentID]["children"].([]map[string]interface{}),
					res[id],
				)
			} else {
				res[parentID]["children"] = []map[string]interface{}{res[id]}
			}

		} else {
			res[parentID] = map[string]interface{}{
				"children": []map[string]interface{}{
					res[id],
				},
			}
		}
	}
	//res[0]["children"].([]map[string]interface{})

	return res
}

// ListToTreeYuanShi
/**
 * 原始方法
 */
func ListToTreeYuanShi(data []map[string]interface{}) map[int]map[string]interface{} {
	res := make(map[int]map[string]interface{})
	for _, v := range data {
		id := int(v["id"].(float64))
		parentID := int(v["parent_id"].(float64))
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
	//return res[0]["children"].([]map[string]interface{})
	return res
}
