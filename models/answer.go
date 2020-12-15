//Package models ... fetch data and interacts directly with database
package models

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/rayspock/go-answer/config"
	"github.com/rayspock/go-answer/utils/exception"

	"github.com/jinzhu/gorm"
)

//GetAnswerByKey ...
func GetAnswerByKey(answer *Answer, key string) (err error) {
	var rows *sql.Rows
	rows, err = config.DB.Table("answer").Select(
		"key, value").Where("key = ?", key).Order(
		"create_date desc").Limit(1).Rows()
	defer rows.Close()
	if err != nil {
		return err
	}
	if rows.Next() {
		config.DB.ScanRows(rows, &answer)
	} else {
		return exception.New(http.StatusNotFound)
	}
	return nil
}

//GetAnswerHistoryByKey ...
func GetAnswerHistoryByKey(histories *[]History, key string) (err error) {
	var rows *sql.Rows
	rows, err = config.DB.Table("history").Select(
		"event, data").Where(
		"event IN ?", []string{"create", "update", "delete"}).Order(
		"create_date asc").Rows()
	defer rows.Close()
	if err != nil {
		return err
	}
	for rows.Next() {
		var history History
		config.DB.ScanRows(rows, &history)
		*histories = append(*histories, history)
	}
	return nil
}

//UpdateAnswerByKey ...
func UpdateAnswerByKey(key, value string) (err error) {
	err = config.DB.Model(&Answer{}).Where("key = ?", key).Update("value", value).Error
	if err != nil {
		return err
	}
	return nil
}

//DeleteAnswerByKey ...
func DeleteAnswerByKey(key string) (err error) {
	answer := new(Answer)
	result := config.DB.Take(&answer)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return exception.NewWithError(http.StatusNotFound, result.Error)
	}
	if result.Error != nil {
		return err
	}
	err = config.DB.Delete(&Answer{}).Where("key = ?", key).Error
	if err != nil {
		return err
	}
	return nil
}

//CreateAnswerByKey ...
func CreateAnswerByKey(key, value string) (err error) {
	answer := new(Answer)
	// check if answer exist
	result := config.DB.Where("key = ? AND value = ?", key, value).First(&answer)
	if result.RowsAffected > 0 {
		return exception.NewWithError(http.StatusBadRequest, errors.New("The answer is already exists"))
	}

	// create new answer
	newAnswer := Answer{Key: key, Value: value, CreateDate: time.Now()}
	result = config.DB.Create(&newAnswer)
	if result.Error != nil {
		return err
	}
	return nil
}
