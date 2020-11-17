package repository

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	ctx context.Context
	db  *mongo.Database
}

// NewAPIServerRepository APIServerRepositoryインターフェイスを表すオブジェクトを作成します
func NewAPIServerRepository(ctx context.Context, db *mongo.Database) APIServerRepository {
	return &apiServerRepository{
		ctx: ctx,
		db:  db,
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
	collection := r.db.Collection("test")

	var b interface{}

	err := bson.UnmarshalExtJSON(body, false, &b)
	res, err := collection.InsertOne(r.ctx, &b)
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
