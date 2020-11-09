package repository_test

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/Hajime3778/api-creator-backend/pkg/apis/model/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func setUpMockDB() (sqlmock.Sqlmock, *gorm.DB) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return strings.Replace(defaultTableName, "_data_table", "", 1)
	}
	d, mock, _ := sqlmock.New()
	conn, _ := gorm.Open("mysql", d)

	return mock, conn
}

func TestGetAll(t *testing.T) {
	mock, db := setUpMockDB()
	modelId, _ := uuid.NewRandom()

	query := regexp.QuoteMeta("SELECT * FROM `models`")
	rows := sqlmock.NewRows([]string{"id", "name", "description", "schema", "created_at", "updated_at"}).
		AddRow(modelId.String(), "test", "description", "schema", time.Now(), time.Now())
	mock.ExpectQuery(query).WillReturnRows(rows)

	modelRepository := repository.NewModelRepository(db)

	model, err := modelRepository.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, model)
}

func TestGetByID(t *testing.T) {
	mock, db := setUpMockDB()
	modelId, _ := uuid.NewRandom()

	query := regexp.QuoteMeta("SELECT * FROM `models` WHERE (id = ?) ORDER BY `models`.`id` ASC LIMIT 1")
	rows := sqlmock.NewRows([]string{"id", "name", "description", "schema", "created_at", "updated_at"}).
		AddRow(modelId.String(), "test", "description", "schema", time.Now(), time.Now())
	mock.ExpectQuery(query).WillReturnRows(rows)

	modelRepository := repository.NewModelRepository(db)

	model, err := modelRepository.GetByID(modelId.String())
	assert.NoError(t, err)
	assert.NotNil(t, model)
}

func TestCreate(t *testing.T) {
	mock, db := setUpMockDB()
	modelId, _ := uuid.NewRandom()

	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.Name = "name"
	mockModel.Description = "description"
	mockModel.Schema = ""
	mockModel.CreatedAt = time.Time{}
	mockModel.UpdatedAt = time.Time{}

	mock.ExpectBegin()
	query := regexp.QuoteMeta("INSERT INTO `models` (`id`,`api_id`,`name`,`description`,`schema`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)")
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

	mockModel := domain.Model{}
	mockModel.ID = modelId.String()
	mockModel.APIID = ""
	mockModel.Name = "name"
	mockModel.Description = "description"
	mockModel.Schema = ""
	mockModel.CreatedAt = time.Time{}
	mockModel.UpdatedAt = time.Time{}

	selectQuery := regexp.QuoteMeta("SELECT * FROM `models` WHERE (id = ?) ORDER BY `models`.`id` ASC LIMIT 1")
	selectRows := sqlmock.NewRows([]string{"id", "api_id", "name", "description", "schema", "created_at", "updated_at"}).AddRow(modelId.String(), "api_id", "name", "description", "schema", time.Now(), time.Now())

	mock.ExpectQuery(selectQuery).WillReturnRows(selectRows)

	mock.ExpectBegin()
	query := regexp.QuoteMeta("UPDATE `models` SET `api_id` = ?, `name` = ?, `description` = ?, `schema` = ?, `updated_at` = ? WHERE `models`.`id` = ?")
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
