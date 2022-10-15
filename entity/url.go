package entity

type Url struct {
	BaseModel
	ID       uint64 `json:"id" gorm:"primaryKey"`
	UserId   uint64 `json:"userId"`
	Redirect string `json:"redirect" gorm:"not null"`
	Code     string `json:"code" gorm:"unique; not null"`
	Clicks   uint64 `json:"clicks" gorm:"clicks"`
	User     User   `json:"-"`
}
