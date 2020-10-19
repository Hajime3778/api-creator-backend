package domain

//API API
type API struct {
	ID          string `json:"id" gorm:"column:id;primary_key"`
	Name        string `json:"name" gorm:"column:name" sql:"not null;type:varchar(40)"`
	URL         string `json:"url" gorm:"column:url" sql:"not null;type:varchar(40)"`
	Description string `json:"description" gorm:"column:description" sql:"not null;type:varchar(200)"`
	CommonColumn
}
