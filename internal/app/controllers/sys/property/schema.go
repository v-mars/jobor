package property

type Role struct {
	ID        uint  	    `json:"id"`
	Name      string        `gorm:"column:name;" json:"name"`
}