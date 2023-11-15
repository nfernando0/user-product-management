package models

type Product struct {
	ID     int          `json:"id" gorm:"primary_key:auto_increment"`
	Name   string       `json:"name" form:"name" gorm:"type: varchar(255)"`
	Desc   string       `json:"desc" form:"desc" gorm:"type: varchar(255)"`
	Price  int          `json:"price" form:"price" gorm:"type: int"`
	Stock  int          `json:"stock" form:"stock" gorm:"type: int"`
	UserId int          `json:"-" form:"user_id"`
	User   UserResponse `json:"user"`
}
