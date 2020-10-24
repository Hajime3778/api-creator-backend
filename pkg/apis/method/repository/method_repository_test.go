package repository_test

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/Hajime3778/api-creator-backend/pkg/apis/method/repository"
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
	methodId, _ := uuid.NewRandom()
	apiId, _ := uuid.NewRandom()

	query := regexp.QuoteMeta("SELECT * FROM `methods`")
	rows := sqlmock.NewRows([]string{
		"id", "api_id", "type", "url", "description",
		"request_parameter", "request_model_id", "response_model_id",
		"is_array", "created_at", "updated_at",
	}).
		AddRow(methodId.String(), apiId.String(), "GET", "url", "description", "id", "", "", false, time.Now(), time.Now())
	mock.ExpectQuery(query).WillReturnRows(rows)

	methodRepository := repository.NewMethodRepository(db)

	method, err := methodRepository.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, method)
}

func TestGetByID(t *testing.T) {
	mock, db := setUpMockDB()
	methodId, _ := uuid.NewRandom()
	apiId, _ := uuid.NewRandom()

	query := regexp.QuoteMeta("SELECT * FROM `methods` WHERE (id = ?) ORDER BY `methods`.`id` ASC LIMIT 1")
	rows := sqlmock.NewRows([]string{
		"id", "api_id", "type", "url", "description",
		"request_parameter", "request_model_id", "response_model_id",
		"is_array", "created_at", "updated_at",
	}).
		AddRow(methodId.String(), apiId.String(), "GET", "url", "description", "id", "", "", false, time.Now(), time.Now())
	mock.ExpectQuery(query).WillReturnRows(rows)

	methodRepository := repository.NewMethodRepository(db)

	method, err := methodRepository.GetByID(methodId.String())
	assert.NoError(t, err)
	assert.NotNil(t, method)
}

func TestCreate(t *testing.T) {
	mock, db := setUpMockDB()
	methodId, _ := uuid.NewRandom()
	apiId, _ := uuid.NewRandom()

	mockMethod := domain.Method{}
	mockMethod.ID = methodId.String()
	mockMethod.APIID = apiId.String()
	mockMethod.Type = "GET"
	mockMethod.URL = "url"
	mockMethod.Description = "test"
	mockMethod.RequestParameter = ""
	mockMethod.RequestModelID = ""
	mockMethod.ResponseModelID = ""
	mockMethod.IsArray = false
	mockMethod.CreatedAt = time.Time{}
	mockMethod.UpdatedAt = time.Time{}

	mock.ExpectBegin()
	query := regexp.QuoteMeta("INSERT INTO `methods` (`id`,`api_id`,`type`,`url`,`description`,`request_parameter`,`request_model_id`,`response_model_id`,`is_array`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")
	mock.ExpectExec(query).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	methodRepository := repository.NewMethodRepository(db)

	_, err := methodRepository.Create(mockMethod)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	mock, db := setUpMockDB()
	methodId, _ := uuid.NewRandom()
	apiId, _ := uuid.NewRandom()

	mockMethod := domain.Method{}
	mockMethod.ID = methodId.String()
	mockMethod.APIID = apiId.String()
	mockMethod.Type = "GET"
	mockMethod.URL = "url"
	mockMethod.Description = "test"
	mockMethod.RequestParameter = ""
	mockMethod.RequestModelID = ""
	mockMethod.ResponseModelID = ""
	mockMethod.IsArray = false
	mockMethod.CreatedAt = time.Time{}
	mockMethod.UpdatedAt = time.Time{}

	selectQuery := regexp.QuoteMeta("SELECT * FROM `methods` WHERE (id = ?) ORDER BY `methods`.`id` ASC LIMIT 1")
	selectRows := sqlmock.NewRows([]string{
		"id", "api_id", "type", "url", "description",
		"request_parameter", "request_model_id", "response_model_id",
		"is_array", "created_at", "updated_at",
	}).AddRow(methodId.String(), apiId.String(), "GET", "url", "description", "id", "", "", false, time.Now(), time.Now())

	mock.ExpectQuery(selectQuery).WillReturnRows(selectRows)

	mock.ExpectBegin()
	query := regexp.QuoteMeta("UPDATE `methods` SET `api_id` = ?, `type` = ?, `url` = ?, `description` = ?, `request_parameter` = ?, `request_model_id` = ?, `response_model_id` = ?, `is_array` = ?, `updated_at` = ? WHERE `methods`.`id` = ?")
	mock.ExpectExec(query).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	methodRepository := repository.NewMethodRepository(db)

	err := methodRepository.Update(mockMethod)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	mock, db := setUpMockDB()
	methodId, _ := uuid.NewRandom()

	mock.ExpectBegin()
	query := regexp.QuoteMeta("DELETE FROM `methods` WHERE `methods`.`id` = ?")
	mock.ExpectExec(query).WithArgs(methodId.String()).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	methodRepository := repository.NewMethodRepository(db)

	err := methodRepository.Delete(methodId.String())
	assert.NoError(t, err)
}
