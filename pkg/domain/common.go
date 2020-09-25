package domain

import "time"

// CommonColumn すべてのDBに共通する項目
type CommonColumn struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at" sql:"not null;type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" sql:"not null;type:datetime"`
}

// CreatedResponse 作成時の返却値
type CreatedResponse struct {
	ID int `json:"id"`
}
