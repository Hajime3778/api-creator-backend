package repository

import (
	"errors"
	"log"

	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/database"
	"go.mongodb.org/mongo-driver/bson"
)

// APIServerRepository Interface
type APIServerRepository interface {
	Get(param string) error
	GetList(param string) error
	Create(body []byte) (string, error)
	Update() error
	Delete() error
}

type apiServerRepository struct {
	db *database.DB
}

// NewAPIServerRepository APIServerRepositoryインターフェイスを表すオブジェクトを作成します
func NewAPIServerRepository(db *database.DB) APIServerRepository {
	return &apiServerRepository{
		db: db,
	}
}

// Get APIServerを1件取得します
func (r *apiServerRepository) Get(param string) error {
	return errors.New("not inprement")
}

// GetList 複数のAPIServerを取得します
func (r *apiServerRepository) GetList(param string) error {
	return errors.New("not inprement")
}

// Create APIServerを追加します
func (r *apiServerRepository) Create(body []byte) (string, error) {

	mongoConn, ctx, cancel := r.db.NewMongoDBConnection()
	defer cancel()

	collection := mongoConn.Collection("test")

	var b interface{}

	err := bson.UnmarshalExtJSON(body, false, &b)
	res, err := collection.InsertOne(ctx, &b)
	if err != nil {
		return "", err
	}
	id := res.InsertedID
	log.Println(id)
	return "sample_id", errors.New("not inprement")
}

// Update APIServerを更新します
func (r *apiServerRepository) Update() error {
	return errors.New("not inprement")
}

// Delete APIServerを削除します
func (r *apiServerRepository) Delete() error {
	return errors.New("not inprement")
}
