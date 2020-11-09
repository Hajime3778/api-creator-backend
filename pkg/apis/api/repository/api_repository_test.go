package repository_test

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/Hajime3778/api-creator-backend/pkg/apis/api/repository"
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
	apiId, _ := uuid.NewRandom()
	mock, db := setUpMockDB()

	query := regexp.QuoteMeta("SELECT * FROM `apis`")
	rows := sqlmock.NewRows([]string{"id", "name", "url", "description", "created_at", "updated_at"}).
		AddRow(apiId.String(), "name", "url", "description", time.Now(), time.Now())
	mock.ExpectQuery(query).WillReturnRows(rows)

	apiRepository := repository.NewAPIRepository(db)

	api, err := apiRepository.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, api)
}

func TestGetByID(t *testing.T) {
	mock, db := setUpMockDB()
	apiId, _ := uuid.NewRandom()

	query := regexp.QuoteMeta("SELECT * FROM `apis` WHERE (id = ?) ORDER BY `apis`.`id` ASC LIMIT 1")
	rows := sqlmock.NewRows([]string{"id", "name", "url", "description", "created_at", "updated_at"}).
		AddRow(apiId.String(), "name", "url", "description", time.Now(), time.Now())
	mock.ExpectQuery(query).WillReturnRows(rows)

	apiRepository := repository.NewAPIRepository(db)

	api, err := apiRepository.GetByID(apiId.String())
	assert.NoError(t, err)
	assert.NotNil(t, api)
}

func TestCreate(t *testing.T) {
	mock, db := setUpMockDB()
	apiId, _ := uuid.NewRandom()

	mockAPI := domain.API{}
	mockAPI.ID = apiId.String()
	mockAPI.Name = "name"
	mockAPI.Description = "test"
	mockAPI.CreatedAt = time.Time{}
	mockAPI.UpdatedAt = time.Time{}

	mock.ExpectBegin()
	query := regexp.QuoteMeta("INSERT INTO `apis` (`id`,`name`,`url`,`description`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)")
	mock.ExpectExec(query).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	apiRepository := repository.NewAPIRepository(db)

	_, err := apiRepository.Create(mockAPI)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	mock, db := setUpMockDB()
	apiId, _ := uuid.NewRandom()

	mockAPI := domain.API{}
	mockAPI.ID = apiId.String()
	mockAPI.Name = "name"
	mockAPI.Description = "test"

	selectQuery := regexp.QuoteMeta("SELECT * FROM `apis` WHERE (id = ?) ORDER BY `apis`.`id` ASC LIMIT 1")
	selectRows := sqlmock.NewRows([]string{"id", "name", "url", "description", "created_at", "updated_at"}).
		AddRow(apiId.String(), "name", "url", "description", mockAPI.CreatedAt, mockAPI.UpdatedAt)
	mock.ExpectQuery(selectQuery).WillReturnRows(selectRows)

	mock.ExpectBegin()
	query := regexp.QuoteMeta("UPDATE `apis` SET `name` = ?, `url` = ?, `description` = ?,  `updated_at` = ? WHERE `apis`.`id` = ?")
	mock.ExpectExec(query).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	apiRepository := repository.NewAPIRepository(db)

	err := apiRepository.Update(mockAPI)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	mock, db := setUpMockDB()
	apiId, _ := uuid.NewRandom()

	mock.ExpectBegin()
	query := regexp.QuoteMeta("DELETE FROM `apis` WHERE `apis`.`id` = ?")
	mock.ExpectExec(query).WithArgs(apiId.String()).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	apiRepository := repository.NewAPIRepository(db)

	err := apiRepository.Delete(apiId.String())
	assert.NoError(t, err)
}
