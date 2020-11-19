package repository

import (
	"errors"
	"log"
	"net/http"

	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// APIServerRepository Interface
type APIServerRepository interface {
	Get(modelName string, param string) (interface{}, int, error)
	GetList(param string) error
	Create(modelName string, body []byte) (interface{}, int, error)
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
func (r *apiServerRepository) Get(modelName string, param string) (interface{}, int, error) {
	mongoConn, ctx, cancel := r.db.NewMongoDBConnection()
	defer cancel()

	collection := mongoConn.Collection(modelName)

	request := bson.M{
		"id": param,
	}
	option := options.FindOne()
	// _idを除外
	option.SetProjection(bson.M{"_id": 0})

	var response bson.M
	err := collection.FindOne(ctx, request, option).Decode(&response)
	if err == mongo.ErrNoDocuments {
		return "", http.StatusNotFound, errors.New("record not found")
	} else if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return response, http.StatusOK, nil
}

// GetList 複数のAPIServerを取得します
func (r *apiServerRepository) GetList(param string) error {
	return errors.New("not inprement")
}

// Create APIServerを追加します
func (r *apiServerRepository) Create(modelName string, body []byte) (interface{}, int, error) {

	mongoConn, ctx, cancel := r.db.NewMongoDBConnection()
	defer cancel()

	collection := mongoConn.Collection(modelName)

	var b bson.M

	err := bson.UnmarshalExtJSON(body, false, &b)
	res, err := collection.InsertOne(ctx, &b)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	log.Println(res.InsertedID)
	return b, http.StatusCreated, nil
}

// Update APIServerを更新します
func (r *apiServerRepository) Update() error {
	return errors.New("not inprement")
}

// Delete APIServerを削除します
func (r *apiServerRepository) Delete() error {
	return errors.New("not inprement")
}
