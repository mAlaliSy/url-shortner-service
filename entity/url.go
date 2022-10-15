package entity

type Url struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Code     string `json:"code" gorm:"unique; not null"`
	Clicks   uint64 `json:"clicks" gorm:"clicks"`
}
