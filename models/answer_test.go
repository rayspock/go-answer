package models

import (
	"database/sql"
	"database/sql/driver"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/rayspock/go-answer/config"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestGetAnswerByKey(t *testing.T) {
	var mock sqlmock.Sqlmock
	initMockDB(&mock, t)

	tests := []struct {
		name      string
		key       string
		mock      func()
		expect    *Answer
		expectErr bool
	}{
		{
			//When everything works as expected
			name: "OK",
			key:  "name",
			mock: func() {
				//We added one row
				rows := sqlmock.NewRows([]string{"key", "val"}).AddRow("name", "john")
				const sql = `SELECT * FROM "answer" WHERE (key = $1) LIMIT 1`
				mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs("name").WillReturnRows(rows)
			},
			expect: &Answer{
				Key: "name",
				Val: "john",
			},
		},
		{
			//When the key does not exist
			name: "Not found",
			key:  "name",
			mock: func() {
				rows := sqlmock.NewRows([]string{"key", "val"})
				const sql = `SELECT * FROM "answer" WHERE (key = $1) LIMIT 1`
				mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs("name").WillReturnRows(rows)
			},
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			var result Answer
			err := GetAnswerByKey(&result, tt.key)
			if err != nil {
				if tt.expectErr {
					assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())
				} else {
					t.Errorf("GetAnswerByKey() error new = %v", err)
					return
				}
			}
			if err == nil && !reflect.DeepEqual(result, *tt.expect) {
				t.Errorf("Received = %v, expected: %v", result, *tt.expect)
			}
		})
	}
}

func TestGetAllAnswer(t *testing.T) {
	var mock sqlmock.Sqlmock
	initMockDB(&mock, t)

	tests := []struct {
		name      string
		key       string
		mock      func()
		expect    *[]Answer
		expectErr bool
	}{
		{
			//When everything works as expected
			name: "OK",
			mock: func() {
				//We added one row
				rows := sqlmock.NewRows([]string{"key", "val"}).AddRow("name", "john")
				const sql = `SELECT * FROM "answer"`
				mock.ExpectQuery(regexp.QuoteMeta(sql)).WillReturnRows(rows)
			},
			expect: &[]Answer{{
				Key: "name",
				Val: "john",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			var result []Answer
			err := GetAllAnswer(&result)
			if err != nil {
				if tt.expectErr {
					assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())
				} else {
					t.Errorf("GetAllAnswer() error new = %v", err)
					return
				}
			}
			if err == nil && !reflect.DeepEqual(result, *tt.expect) {
				t.Errorf("Received = %v, expected: %v", result, *tt.expect)
			}
		})
	}
}

func TestGetAnswerHistoryByKey(t *testing.T) {
	var mock sqlmock.Sqlmock
	initMockDB(&mock, t)

	tests := []struct {
		name      string
		key       string
		mock      func()
		expect    *[]History
		expectErr bool
	}{
		{
			//When everything works as expected
			name: "OK",
			key:  "name",
			mock: func() {
				//We added one row
				rows := sqlmock.NewRows([]string{"event", "data"}).AddRow("create", []byte(`{
					"key": "name",
					"value": "John"
				}`)).AddRow("delete", []byte(`{
					"key": "name"
				}`))
				const sql = `SELECT event, data FROM "history" WHERE (event IN ($1,$2,$3) and key = $4)`
				mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs("create", "update", "delete", "name").WillReturnRows(rows)
			},
			expect: &[]History{
				{
					Event: "create",
					Data: Answer{
						Key: "name",
						Val: "John",
					},
				},
				{
					Event: "delete",
					Data: Answer{
						Key: "name",
					},
				},
			},
		},
		{
			//When the key does not exist
			name: "Not found",
			key:  "name",
			mock: func() {
				const sql = `SELECT event, data FROM "history" WHERE (event IN ($1,$2,$3) and key = $4)`
				mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs("create", "update", "delete", "name").WillReturnRows(sqlmock.NewRows(nil))
			},
			expect: &[]History{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			var result []History
			err := GetAnswerHistoryByKey(&result, tt.key)
			if err != nil {
				if tt.expectErr {
					assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())
				} else {
					t.Errorf("GetAnswerHistoryByKey() error new = %v", err)
					return
				}
			}
			if err == nil && !reflect.DeepEqual(result, *tt.expect) {
				t.Errorf("Received = %v, expected: %v", result, *tt.expect)
			}
		})
	}
}

func TestUpdateAnswerByKey(t *testing.T) {
	var mock sqlmock.Sqlmock
	initMockDB(&mock, t)

	tests := []struct {
		name      string
		key       string
		value     string
		mock      func()
		expect    *Answer
		expectErr bool
	}{
		{
			//When everything works as expected
			name:  "OK",
			key:   "name",
			value: "Ray",
			mock: func() {
				const sqlUpdate = `UPDATE "answer" SET "val" = $1 WHERE (key = $2)`

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(sqlUpdate)).
					WithArgs("Ray", "name").
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := UpdateAnswerByKey(tt.key, tt.value)
			if err != nil {
				t.Errorf("UpdateAnswerByKey() error new = %v", err)
				return
			}
		})
	}
}

func TestDeleteAnswerByKey(t *testing.T) {
	var mock sqlmock.Sqlmock
	initMockDB(&mock, t)

	tests := []struct {
		name      string
		key       string
		mock      func()
		expect    *Answer
		expectErr bool
	}{
		{
			//When everything works as expected
			name: "OK",
			key:  "name",
			mock: func() {
				const sql = `DELETE FROM "answer" WHERE (key = $1)`

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(sql)).
					WithArgs("name").
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := DeleteAnswerByKey(tt.key)
			if err != nil {
				t.Errorf("DeleteAnswerByKey() error new = %v", err)
				return
			}
		})
	}
}

func TestCreateAnswerByKey(t *testing.T) {
	var mock sqlmock.Sqlmock
	initMockDB(&mock, t)

	tests := []struct {
		name      string
		answer    *Answer
		mock      func()
		expect    *Answer
		expectErr bool
	}{
		{
			//When everything works as expected
			name: "OK",
			answer: &Answer{
				Key: "name",
				Val: "John",
			},
			mock: func() {
				const sqlInsert = `INSERT INTO "answer" ("key","val") VALUES ($1,$2)`
				const sql = `SELECT * FROM "answer" WHERE (key = $1) LIMIT 1`

				mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs("name").WillReturnRows(sqlmock.NewRows(nil))
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).
					WithArgs("name", "John").
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := CreateAnswerByKey(tt.answer)
			if err != nil {
				t.Errorf("CreateAnswerByKey() error new = %v", err)
				return
			}
		})
	}
}

func TestSaveToHistory(t *testing.T) {
	var mock sqlmock.Sqlmock
	initMockDB(&mock, t)

	tests := []struct {
		name      string
		event     string
		answer    *Answer
		mock      func()
		expect    *Answer
		expectErr bool
	}{
		{
			//When everything works as expected
			name:  "OK",
			event: "create",
			answer: &Answer{
				Key: "name",
				Val: "John",
			},
			mock: func() {
				const sqlInsert = `INSERT INTO "history" ("event","key","data","create_date") VALUES ($1,$2,$3,$4)`

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).
					WithArgs("create", "name", &Answer{
						Key: "name",
						Val: "John",
					}, AnyTime{}).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := SaveToHistory(tt.event, tt.answer)
			if err != nil {
				t.Errorf("SaveToHistory() error new = %v", err)
				return
			}
		})
	}
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func initMockDB(mock *sqlmock.Sqlmock, t *testing.T) {
	var db *sql.DB
	var err error
	db, *mock, err = sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	config.DB, err = gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("Status: %s", err)
	}
}
