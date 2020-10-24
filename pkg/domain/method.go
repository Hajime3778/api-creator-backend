package domain

//Method Method
type Method struct {
	ID               string `json:"id" gorm:"column:id;primary_key"`
	APIID            string `json:"api_id" gorm:"column:api_id;"`
	Type             string `json:"type" gorm:"column:type"`
	URL              string `json:"url" gorm:"column:url"`
	Description      string `json:"description" gorm:"column:description"`
	RequestParameter string `json:"request_parameter" gorm:"column:request_parameter"`
	RequestModelID   string `json:"request_model_id" gorm:"column:request_model_id"`
	ResponseModelID  string `json:"response_model_id" gorm:"column:response_model_id"`
	IsArray          bool   `json:"is_array" gorm:"column:is_array"`
	CommonColumn
}
