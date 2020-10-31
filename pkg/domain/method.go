package domain

//Method Method
type Method struct {
	ID               string `json:"id" gorm:"column:id;primary_key"`
	APIID            string `json:"apiId" gorm:"column:api_id;"`
	Type             string `json:"type" gorm:"column:type"`
	URL              string `json:"url" gorm:"column:url"`
	Description      string `json:"description" gorm:"column:description"`
	RequestParameter string `json:"requestParameter" gorm:"column:request_parameter"`
	RequestModelID   string `json:"requestModelId" gorm:"column:request_model_id"`
	ResponseModelID  string `json:"responseModelId" gorm:"column:response_model_id"`
	IsArray          bool   `json:"isArray" gorm:"column:is_array"`
	CommonColumn
}
