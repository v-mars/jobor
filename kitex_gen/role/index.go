package role

func (p *RolePut) GetApiIdsInt() []int {
	var arrayInt = make([]int, 0)
	if p.ApiIds != nil {
		for _, v := range p.ApiIds.GetValues() {
			v := v
			arrayInt = append(arrayInt, int(v.GetNumberValue()))
		}
	}
	return arrayInt
}
func (p *RolePut) GetPathInt() []int {
	var arrayInt = make([]int, 0)
	if p.Path != nil {
		for _, v := range p.Path.GetValues() {
			v := v
			arrayInt = append(arrayInt, int(v.GetNumberValue()))
		}
	}
	return arrayInt
}
