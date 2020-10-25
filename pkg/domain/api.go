package domain

//API API
type API struct {
	ID          string `json:"id" gorm:"column:id;primary_key"`
	Name        string `json:"name" gorm:"column:name"`
	URL         string `json:"url" gorm:"column:url"`
	Description string `json:"description" gorm:"column:description"`
	ModelID     string `json:"model_id" gorm:"column:model_id"`
	CommonColumn
}
