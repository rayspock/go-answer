//Package models ... fetch data and interacts directly with database
package models

import (
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
	result := config.DB.Where("key = ?", key).First(answer)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return exception.NewWithError(http.StatusNotFound, result.Error)
	}
	if (result.Error != nil) {
		return result.Error
	}
	return nil
}

//GetAllAnswer ... 
func GetAllAnswer(answers *[]Answer) (err error) {
	result := config.DB.Find(answers)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return exception.NewWithError(http.StatusNotFound, result.Error)
	}
	if (result.Error != nil) {
		return result.Error
	}
	return nil
}

//GetAnswerHistoryByKey ...
func GetAnswerHistoryByKey(histories *[]History, key string) (err error) {
	result := config.DB.Select("event, data").Find(
		histories, "event IN (?) and key = ?", []string{"create", "update", "delete"}, key).Order(
		"create_date asc")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return exception.NewWithError(http.StatusNotFound, result.Error)
	}
	if result.Error != nil {
		return err
	}
	return nil
}

//UpdateAnswerByKey ...
func UpdateAnswerByKey(key, value string) (err error) {
	result := config.DB.Model(&Answer{}).Where("key = ?", key).Update("val", value)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return exception.NewWithError(http.StatusBadRequest, errors.New("the answer doesn't exist or has been deleted"))
	}
	return nil
}

//DeleteAnswerByKey ...
func DeleteAnswerByKey(key string) (err error) {
	result := config.DB.Where("key = ?", key).Delete(&Answer{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return exception.NewWithError(http.StatusBadRequest, errors.New("the answer doesn't exist or has been deleted"))
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
	history := History{Event: event, Key: data.Key, Data: *data, CreateDate: &today}
	result := config.DB.Create(&history)
	if result.Error != nil {
		return err
	}
	return nil
}
