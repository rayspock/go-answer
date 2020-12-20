//Package controllers ... handle http requests for the Answer
package controllers

import (
	"net/http"

	"github.com/rayspock/go-answer/models"
	"github.com/rayspock/go-answer/utils/exception"

	"github.com/gin-gonic/gin"
)

// GetAnswerByKey godoc
// @Summary get answer by Key
// @Description get answer (returns the latest answer for the given key)
// @Accept  json
// @Produce  json
// @Param key path string true "Answer key"
// @Success 200 {object} models.Answer
// @Failure 400 {object} exception.HTTPError
// @Failure 404 {object} exception.HTTPError
// @Failure 500 {object} exception.HTTPError
// @Router /answer/{key} [get]
func GetAnswerByKey(c *gin.Context) error {
	var payload models.ParamPayload
	if err := c.ShouldBindUri(&payload); err != nil {
		return exception.NewWithError(http.StatusBadRequest, err)
	}
	var answer models.Answer
	err := models.GetAnswerByKey(&answer, payload.Key)
	if err != nil {
		return err
	}
	models.SaveToHistory("get", &answer)
	c.JSON(http.StatusOK, answer)
	return nil
}

// GetAnswerHistoryByKey godoc
// @Summary get answer's history by key
// @Description get history for given key (returns an array of events in chronological order)
// @Accept  json
// @Produce  json
// @Param key path string true "Answer key"
// @Success 200 {object} []models.History
// @Failure 400 {object} exception.HTTPError
// @Failure 404 {object} exception.HTTPError
// @Failure 500 {object} exception.HTTPError
// @Router /answer/{key}/history [get]
func GetAnswerHistoryByKey(c *gin.Context) error {
	var payload models.ParamPayload
	if err := c.ShouldBindUri(&payload); err != nil {
		return exception.NewWithError(http.StatusBadRequest, err)
	}
	var histories []models.History
	err := models.GetAnswerHistoryByKey(&histories, payload.Key)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, histories)
	return nil
}

// UpdateAnswerByKey godoc
// @Summary update answer
// @Description update answer
// @Accept  json
// @Produce  json
// @Param key path string true "Answer key"
// @Param body body models.AnswerPayload true "Answer value"
// @Success 200 {string} string "success"
// @Failure 400 {object} exception.HTTPError
// @Failure 404 {object} exception.HTTPError
// @Failure 500 {object} exception.HTTPError
// @Router /answer/{key} [put]
func UpdateAnswerByKey(c *gin.Context) error {
	var payload models.ParamPayload
	var answerPayload models.AnswerPayload
	if err := c.ShouldBindUri(&payload); err != nil {
		return exception.NewWithError(http.StatusBadRequest, err)
	}
	if err := c.ShouldBindJSON(&answerPayload); err == nil {
		err := models.UpdateAnswerByKey(payload.Key, answerPayload.Value)
		if err != nil {
			return err
		}
		models.SaveToHistory("update", &models.Answer{Key: payload.Key, Val: answerPayload.Value})
		c.String(http.StatusOK, "success")
	} else {
		return exception.NewWithError(http.StatusBadRequest, err)
	}
	return nil
}

// CreateAnswerByKey godoc
// @Summary create answer
// @Description create answer
// @Accept  json
// @Produce  json
// @Param body body models.Answer true "Answer"
// @Success 200 {string} string "success"
// @Failure 400 {object} exception.HTTPError
// @Failure 404 {object} exception.HTTPError
// @Failure 500 {object} exception.HTTPError
// @Router /answer/ [post]
func CreateAnswerByKey(c *gin.Context) error {
	var answer models.Answer
	if err := c.ShouldBindJSON(&answer); err == nil {
		err := models.CreateAnswerByKey(&answer)
		if err != nil {
			return err
		}
		models.SaveToHistory("create", &answer)
		c.String(http.StatusOK, "success")
	} else {
		return exception.NewWithError(http.StatusBadRequest, err)
	}
	return nil
}

// DeleteAnswerByKey godoc
// @Summary delete answer
// @Description delete answer
// @Accept  json
// @Produce  json
// @Param key path string true "Answer key"
// @Success 200 {string} string "success"
// @Failure 400 {object} exception.HTTPError
// @Failure 404 {object} exception.HTTPError
// @Failure 500 {object} exception.HTTPError
// @Router /answer/{key} [delete]
func DeleteAnswerByKey(c *gin.Context) error {
	var payload models.ParamPayload
	if err := c.ShouldBindUri(&payload); err != nil {
		return exception.NewWithError(http.StatusBadRequest, err)
	}
	err := models.DeleteAnswerByKey(payload.Key)
	if err != nil {
		return err
	}
	models.SaveToHistory("delete", &models.Answer{Key: payload.Key})
	c.String(http.StatusOK, "success")
	return nil
}
