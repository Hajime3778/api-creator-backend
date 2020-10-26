package repository_test

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/Hajime3778/api-creator-backend/pkg/apis/model/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/database"
	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func setUpMockDB() (sqlmock.Sqlmock, *database.DB) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return strings.Replace(defaultTableName, "_data_table", "", 1)
	}
	d, mock, _ := sqlmock.New()
	db := new(database.DB)
	db.Connection, _ = gorm.Open("mysql", d)

	return mock, db
}

func TestGetAll(t *testing.T) {
	mock, db := setUpMockDB()
	modelId, _ := uuid.NewRandom()
	apiId, _ := uuid.NewRandom()

	query := regexp.QuoteMeta("SELECT * FROM `models`")
	rows := sqlmock.NewRows([]string{
		"id", "api_id", "type", "url", "description",
		"request_parameter", "request_model_id", "response_model_id",
		"is_array", "created_at", "updated_at",
	}).
		AddRow(modelId.String(), apiId.String(), "GET", "url", "description", "id", "", "", false, time.Now(), time.Now())
	mock.ExpectQuery(query).WillReturnRows(rows)

	modelRepository := repository.NewModelRepository(db)

	model, err := modelRepository.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, model)
}

func TestGetByID(t *testing.T) {
	mock, db := setUpMockDB()
	modelId, _ := uuid.NewRandom()
	apiId, _ := uuid.NewRandom()

	query := regexp.QuoteMeta("SELECT * FROM `models` WHERE (id = ?) ORDER BY `models`.`id` ASC LIMIT 1")
	rows := sqlmock.NewRows([]string{
		"id", "api_id", "type", "url", "description",
		"request_parameter", "request_model_id", "response_model_id",
		"is_array", "created_at", "updated_at",
	}).
		AddRow(modelId.String(), apiId.String(), "GET", "url", "description", "id", "", "", false, time.Now(), time.Now())
	mock.ExpectQuery(query).WillReturnRows(rows)

	modelRepository := repository.NewModelRepository(db)

	model, err := modelRepository.GetByID(modelId.String())
	assert.NoError(t, err)
	assert.NotNil(t, model)
}

func TestCreate(t *testing.T) {
	mock, db := setUpMockDB()
	modelId, _ := uuid.NewRandom()
	apiId, _ := uuid.NewRandom()

	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.APIID = apiId.String()
	mockModel.Type = "GET"
	mockModel.URL = "url"
	mockModel.Description = "test"
	mockModel.RequestParameter = ""
	mockModel.RequestModelID = ""
	mockModel.ResponseModelID = ""
	mockModel.IsArray = false
	mockModel.CreatedAt = time.Time{}
	mockModel.UpdatedAt = time.Time{}

	mock.ExpectBegin()
	query := regexp.QuoteMeta("INSERT INTO `models` (`id`,`api_id`,`type`,`url`,`description`,`request_parameter`,`request_model_id`,`response_model_id`,`is_array`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")
	mock.ExpectExec(query).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	modelRepository := repository.NewModelRepository(db)

	_, err := modelRepository.Create(mockModel)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	mock, db := setUpMockDB()
	modelId, _ := uuid.NewRandom()
	apiId, _ := uuid.NewRandom()

	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.APIID = apiId.String()
	mockModel.Type = "GET"
	mockModel.URL = "url"
	mockModel.Description = "test"
	mockModel.RequestParameter = ""
	mockModel.RequestModelID = ""
	mockModel.ResponseModelID = ""
	mockModel.IsArray = false
	mockModel.CreatedAt = time.Time{}
	mockModel.UpdatedAt = time.Time{}

	selectQuery := regexp.QuoteMeta("SELECT * FROM `models` WHERE (id = ?) ORDER BY `models`.`id` ASC LIMIT 1")
	selectRows := sqlmock.NewRows([]string{
		"id", "api_id", "type", "url", "description",
		"request_parameter", "request_model_id", "response_model_id",
		"is_array", "created_at", "updated_at",
	}).AddRow(modelId.String(), apiId.String(), "GET", "url", "description", "id", "", "", false, time.Now(), time.Now())

	mock.ExpectQuery(selectQuery).WillReturnRows(selectRows)

	mock.ExpectBegin()
	query := regexp.QuoteMeta("UPDATE `models` SET `api_id` = ?, `type` = ?, `url` = ?, `description` = ?, `request_parameter` = ?, `request_model_id` = ?, `response_model_id` = ?, `is_array` = ?, `updated_at` = ? WHERE `models`.`id` = ?")
	mock.ExpectExec(query).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	modelRepository := repository.NewModelRepository(db)

	err := modelRepository.Update(mockModel)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	mock, db := setUpMockDB()
	modelId, _ := uuid.NewRandom()

	mock.ExpectBegin()
	query := regexp.QuoteMeta("DELETE FROM `models` WHERE `models`.`id` = ?")
	mock.ExpectExec(query).WithArgs(modelId.String()).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	modelRepository := repository.NewModelRepository(db)

	err := modelRepository.Delete(modelId.String())
	assert.NoError(t, err)
}
