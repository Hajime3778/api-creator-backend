package domain

//Model Model
type Model struct {
	ID          string `json:"id" gorm:"column:id;primary_key"`
	APIID       string `json:"apiId" gorm:"column:api_id"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	Schema      string `json:"schema" gorm:"column:schema"`
	CommonColumn
}
