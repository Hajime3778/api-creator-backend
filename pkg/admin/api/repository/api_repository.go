package repository

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Hajime3778/api-creator-backend/pkg/domain"

	"github.com/jinzhu/gorm"
)

// APIRepository Interface
type APIRepository interface {
	GetAll() ([]domain.API, error)
	GetByID(id string) (domain.API, error)
	GetByURL(url string) (domain.API, error)
	Create(api domain.API) (string, error)
	Update(api domain.API) error
	Delete(id string, methods []domain.Method, model domain.Model) error
}

type apiRepository struct {
	db *gorm.DB
}

// NewAPIRepository APIRepositoryインターフェイスを表すオブジェクトを作成します
func NewAPIRepository(db *gorm.DB) APIRepository {
	return &apiRepository{
		db: db,
	}
}

// GetAll すべてのAPIを取得します
func (r *apiRepository) GetAll() ([]domain.API, error) {
	apis := []domain.API{}
	err := r.db.Find(&apis).Error

	return apis, err
}

// GetByID APIを1件取得します
func (r *apiRepository) GetByID(id string) (domain.API, error) {
	api := domain.API{}
	err := r.db.Where("id = ?", id).First(&api).Error

	return api, err
}

// GetByURL APIを1件取得します
func (r *apiRepository) GetByURL(url string) (domain.API, error) {
	api := domain.API{}
	err := r.db.Raw("SELECT * FROM `apis` WHERE '" + url + "' like CONCAT('%', url, '%')").Scan(&api).Error

	return api, err
}

// Create APIを作成します
func (r *apiRepository) Create(api domain.API) (string, error) {
	err := r.db.Create(&api).Error
	id := api.ID
	return id, err
}

// Update APIを更新します
func (r *apiRepository) Update(api domain.API) error {
	targetAPI := domain.API{}

	err := r.db.Where("id = ?", api.ID).First(&targetAPI).Error
	if err != nil {
		return err
	}

	return r.db.Save(&api).Error
}

// Delete APIを削除します(関連するメソッド、モデルも含めて)
func (r *apiRepository) Delete(id string, methods []domain.Method, model domain.Model) error {
	api := domain.API{}
	api.ID = id

	r.db.Transaction(func(tx *gorm.DB) error {
		// for _, m := range methods {
		// 	method := domain.Method{ID: m.ID}
		// 	if err := tx.Delete(&method).Error; err != nil {
		// 		return err
		// 	}
		// }
		// if err := tx.Delete(&model).Error; err != nil {
		// 	return err
		// }
		// if err := tx.Delete(&api).Error; err != nil {
		// 	return err
		// }
		// localで実行する場合 removeCollectionRequest("http://localhost:9000/" + model.Name)
		if err := removeCollectionRequest("http://api-server/" + model.Name); err != nil {
			return err
		}
		return nil
	})
	return nil
}

func removeCollectionRequest(url string) error {
	request, err := http.NewRequest("DELETE", url+"/remove-target-collection", nil)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	log.Println(string(b))
	return nil
}
