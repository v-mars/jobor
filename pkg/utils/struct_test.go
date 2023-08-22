package utils

import (
	"encoding/json"
	"fmt"
)

func Json6() {
	type Stbb struct {
		Sdb string `json:"sdb"`
		Sdc string `json:"sdc"`
	}
	type Stru struct {
		ID int    `json:"id"`
		Cd string `json:"cd"`
		CC string
		Stbb
		AA string `json:"aa"`
		BB string `json:"-"`
		St struct {
			Sta string `json:"sta"`
			Stb string `json:"stb"`
		} `json:"st"`
	}
	var st = Stru{AA: "a1", BB: "b1", CC: "c1", Cd: "cd1", ID: 10}
	st.St.Sta = "sta1"
	st.Sdb = "sdb1"
	marshal, err := json.Marshal(st)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("marshal:", string(marshal))
	var mm map[string]interface{}
	err = json.Unmarshal(marshal, &mm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("marshal:", mm)
	var names []string
	var names2 []string
	err = GetStructFieldsByReflect(st, &names, &names2)
	if err != nil {
		fmt.Println("names err:", err)
		return
	}
	fmt.Println("names:", names)
}
