package entity

type User struct {
	BaseModel
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-" gorm:"not null"`
}
