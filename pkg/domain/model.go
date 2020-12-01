package domain

import (
	"encoding/json"
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

//Model Model
type Model struct {
	ID          string `json:"id" gorm:"column:id;primary_key"`
	APIID       string `json:"apiId" gorm:"column:api_id"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	Schema      string `json:"schema" gorm:"column:schema"`
	CommonColumn
}

// ValidateSchema JsonSchemaを検証します。
func (m *Model) ValidateSchema() error {
	sl := gojsonschema.NewSchemaLoader()
	sl.Validate = true
	err := sl.AddSchemas(gojsonschema.NewStringLoader(m.Schema))
	if err != nil {
		return err
	}

	_, err = m.GetKeyNames()
	if err == nil {
		return err
	}

	return nil
}

// GetKeyNames jsonschemaからkey名を取得します。
func (m *Model) GetKeyNames() ([]string, error) {
	var keyNames []string
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(m.Schema), &jsonMap)
	if err != nil {
		return nil, err
	}

	keys := jsonMap["keys"].([]interface{})

	if keys == nil {
		return keyNames, nil
	}

	properties := jsonMap["properties"].(map[string]interface{})
	missingPropertyName := ""

	for _, key := range keys {
		// 存在しない項目がKeyに指定されていた場合
		if properties[key.(string)] == nil {
			if missingPropertyName != "" {
				missingPropertyName = missingPropertyName + ", "
			}
			missingPropertyName = missingPropertyName + key.(string)
		} else {
			keyNames = append(keyNames, key.(string))
		}
	}

	if missingPropertyName != "" {
		return nil, errors.New("存在しない項目" + missingPropertyName + "がkeysに指定されています")
	}

	// TODO: 複数キー指定は今は対応しない
	if len(keyNames) > 1 {
		return nil, errors.New("keysが複数指定されています")
	}

	return keyNames, nil
}
