package entity

type User struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-" gorm:"not null"`
}
