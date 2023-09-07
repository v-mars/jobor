package user

import (
	"encoding/json"
	"jobor/pkg/utils"
)

func (i *Userinfo) IsAdmin() bool {
	if utils.InOfStr("admin", i.Roles) {
		return true
	} else {
		return false
	}
}

func (i *Userinfo) IsDevOps() bool {
	if utils.InOfStr("devops", i.Roles) {
		return true
	} else {
		return false
	}
}

func (i *Userinfo) IsOps() bool {
	if utils.InOfStr("ops", i.Roles) {
		return true
	} else {
		return false
	}
}

func (i *Userinfo) IsQa() bool {
	if utils.InOfStr("qa", i.Roles) {
		return true
	} else {
		return false
	}
}

func (i *Userinfo) IsDeveloper() bool {
	if utils.InOfStr("developer", i.Roles) {
		return true
	} else {
		return false
	}
}

func (i *Userinfo) JsonString() string {
	marshal, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return string(marshal)
}

func (i *Userinfo) Map() (mapRes map[string]interface{}) {
	err := utils.AnyToAny(i, &mapRes)
	if err != nil {
		panic(err)
	}
	return mapRes
}

func (p *PutUserReq) GetRoleIdsInt() []int {
	var arrayInt = make([]int, 0)
	if p.RoleIds != nil {
		for _, v := range p.RoleIds.GetValues() {
			v := v
			arrayInt = append(arrayInt, int(v.GetNumberValue()))
		}
	}
	return arrayInt
}
