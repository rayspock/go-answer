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

const (
	layout = "2006-01-02"
)

//GetAnswerByKey ...
func GetAnswerByKey(answer *Answer, key string) (err error) {
	result := config.DB.First(answer)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return exception.NewWithError(http.StatusNotFound, result.Error)
	}
	return nil
}

//GetAnswerHistoryByKey ...
func GetAnswerHistoryByKey(histories *[]History, key string) (err error) {
	var rows *sql.Rows
	rows, err = config.DB.Table("history").Select(
		"event, data").Where("event IN (?)", []string{"create", "update", "delete"}).Order(
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
	err = config.DB.Model(&Answer{}).Where("key = ?", key).Update("val", value).Error
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
func CreateAnswerByKey(answer *Answer) (err error) {
	answerCheck := new(Answer)
	// check if answer exist
	result := config.DB.Where("key = ?", answer.Key).First(&answerCheck)
	if result.RowsAffected > 0 {
		return exception.NewWithError(http.StatusBadRequest, errors.New("the answer already exists"))
	}

	// create new answer
	result = config.DB.Create(answer)
	if result.Error != nil {
		return err
	}
	return nil
}

//SaveToHistory ... Save data manipulation log to history
func SaveToHistory(event string, data *Answer) (err error) {
	today := time.Now()
	history := History{Event: event, Key: data.Key, Data: *data, CreateDate: today}
	result := config.DB.Create(&history)
	if result.Error != nil {
		return err
	}
	return nil
}
