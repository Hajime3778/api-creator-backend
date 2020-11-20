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
	Get(modelName string, key string, param string) (interface{}, int, error)
	GetList(modelName string, key string, param string) (interface{}, int, error)
	Create(modelName string, body []byte) (interface{}, int, error)
	Update() (interface{}, int, error)
	Delete() (interface{}, int, error)
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
func (r *apiServerRepository) Get(modelName string, key string, param string) (interface{}, int, error) {
	mongoConn, ctx, cancel := r.db.NewMongoDBConnection()
	defer cancel()

	collection := mongoConn.Collection(modelName)

	request := bson.M{
		key: param,
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
func (r *apiServerRepository) GetList(modelName string, key string, param string) (interface{}, int, error) {
	mongoConn, ctx, cancel := r.db.NewMongoDBConnection()
	defer cancel()

	collection := mongoConn.Collection(modelName)

	request := bson.D{}
	if key != "" && param != "" {
		request = bson.D{{Key: key, Value: param}}
	}

	option := options.Find()
	// _idを除外
	option.SetProjection(bson.D{{Key: "_id", Value: 0}})

	var response []bson.M
	cur, err := collection.Find(ctx, request, option)
	if err == mongo.ErrNoDocuments {
		return "", http.StatusNotFound, errors.New("record not found")
	} else if err != nil {
		return "", http.StatusInternalServerError, err
	}

	for cur.Next(ctx) {
		var doc bson.M
		if err = cur.Decode(&doc); err != nil {
			return "", http.StatusInternalServerError, err
		}
		response = append(response, doc)
	}

	return response, http.StatusOK, nil
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
func (r *apiServerRepository) Update() (interface{}, int, error) {
	return "", http.StatusNotImplemented, nil
}

// Delete APIServerを削除します
func (r *apiServerRepository) Delete() (interface{}, int, error) {
	return "", http.StatusNotImplemented, nil
}
