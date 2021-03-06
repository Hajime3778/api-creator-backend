package repository

import (
	"errors"
	"net/http"

	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// APIServerRepository Interface
type APIServerRepository interface {
	Get(modelName string, key string, param interface{}) (interface{}, int, error)
	GetList(modelName string, key string, param interface{}) (interface{}, int, error)
	Create(modelName string, key string, body []byte) (interface{}, int, error)
	Update(modelName string, key string, body []byte) (interface{}, int, error)
	Delete(modelName string, key string, param interface{}) (interface{}, int, error)
	RemoveCollection(modelName string) (interface{}, int, error)
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
func (r *apiServerRepository) Get(modelName string, key string, param interface{}) (interface{}, int, error) {
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
func (r *apiServerRepository) GetList(modelName string, key string, param interface{}) (interface{}, int, error) {
	mongoConn, ctx, cancel := r.db.NewMongoDBConnection()
	defer cancel()

	collection := mongoConn.Collection(modelName)

	// TODO:条件指定する場合以下が参考になる。
	// https://qiita.com/nishina555/items/9e20211e8d6f12fdb7b7#%E9%83%A8%E5%88%86%E4%B8%80%E8%87%B4%E6%A4%9C%E7%B4%A2%E6%AD%A3%E8%A6%8F%E8%A1%A8%E7%8F%BE
	request := bson.D{}
	if key != "" && param != "" {
		request = bson.D{{Key: key, Value: param}}
	}

	option := options.Find()
	// _idを除外
	option.SetProjection(bson.D{{Key: "_id", Value: 0}})

	var response []bson.M
	cur, err := collection.Find(ctx, request, option)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	for cur.Next(ctx) {
		var doc bson.M
		if err = cur.Decode(&doc); err != nil {
			return "", http.StatusInternalServerError, err
		}
		response = append(response, doc)
	}

	if response == nil {
		return "", http.StatusNotFound, errors.New("record not found")
	}

	return response, http.StatusOK, nil
}

// Create APIServerを追加します
func (r *apiServerRepository) Create(modelName string, keyName string, body []byte) (interface{}, int, error) {

	mongoConn, ctx, cancel := r.db.NewMongoDBConnection()
	defer cancel()

	collection := mongoConn.Collection(modelName)

	var b bson.M

	err := bson.UnmarshalExtJSON(body, false, &b)
	_, err = collection.InsertOne(ctx, &b)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	return b, http.StatusCreated, nil
}

// Update APIServerを更新します
func (r *apiServerRepository) Update(modelName string, keyName string, body []byte) (interface{}, int, error) {
	mongoConn, ctx, cancel := r.db.NewMongoDBConnection()
	defer cancel()

	collection := mongoConn.Collection(modelName)

	var requestBody bson.M
	var updateModel bson.D

	err := bson.UnmarshalExtJSON(body, false, &requestBody)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	// TODO: UniqueKey設定できるようにする
	id := requestBody[keyName].(interface{})

	err = bson.UnmarshalExtJSON(body, false, &updateModel)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{{Key: "$set", Value: updateModel}}

	_, err = collection.UpdateOne(ctx, filter, &update)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	return requestBody, http.StatusOK, nil
}

// Delete APIServerを削除します
func (r *apiServerRepository) Delete(modelName string, key string, param interface{}) (interface{}, int, error) {
	mongoConn, ctx, cancel := r.db.NewMongoDBConnection()
	defer cancel()

	collection := mongoConn.Collection(modelName)

	request := bson.M{
		key: param,
	}

	_, err := collection.DeleteOne(ctx, request)
	if err == mongo.ErrNoDocuments {
		return "", http.StatusNotFound, errors.New("record not found")
	} else if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return "", http.StatusNoContent, nil
}

// RemoveCollection Collectionを削除します
func (r *apiServerRepository) RemoveCollection(modelName string) (interface{}, int, error) {
	mongoConn, ctx, cancel := r.db.NewMongoDBConnection()
	defer cancel()

	collection := mongoConn.Collection(modelName)

	if err := collection.Drop(ctx); err != nil {
		return "", http.StatusInternalServerError, err
	}

	return "", http.StatusNoContent, nil
}
