//Package controllers ... handle http requests for the Answer
package controllers

import (
	"net/http"

	"github.com/rayspock/go-answer/models"

	"github.com/gin-gonic/gin"
)

// GetAnswerByKey godoc
// @Summary get answer by Key
// @Description get answer (returns the latest answer for the given key)
// @Accept  json
// @Produce  json
// @Param key path string true "Answer key"
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /answer/{key} [get]
func GetAnswerByKey(c *gin.Context) error {
	key := c.Param("key")
	var answer models.Answer
	err := models.GetAnswerByKey(&answer, key)
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
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /answer/{key}/history [get]
func GetAnswerHistoryByKey(c *gin.Context) error {
	key := c.Param("key")
	var histories []models.History
	err := models.GetAnswerHistoryByKey(&histories, key)
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
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /answer/{key} [put]
func UpdateAnswerByKey(c *gin.Context) error {
	var payload models.AnswerPayload
	key := c.Param("key")
	c.BindJSON(&payload)
	err := models.UpdateAnswerByKey(key, payload.Value)
	if err != nil {
		return err
	}
	models.SaveToHistory("update", &models.Answer{Key: key, Val: payload.Value})
	c.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
	return nil
}

// CreateAnswerByKey godoc
// @Summary create answer
// @Description create answer
// @Accept  json
// @Produce  json
// @Param key path string true "Answer key"
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /answer/ [post]
func CreateAnswerByKey(c *gin.Context) error {
	var answer models.Answer
	c.BindJSON(&answer)
	err := models.CreateAnswerByKey(&answer)
	if err != nil {
		return err
	}
	models.SaveToHistory("create", &answer)
	c.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
	return nil
}

// DeleteAnswerByKey godoc
// @Summary delete answer
// @Description delete answer
// @Accept  json
// @Produce  json
// @Param key path string true "Answer key"
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /answer/{key} [delete]
func DeleteAnswerByKey(c *gin.Context) error {
	key := c.Param("key")
	err := models.DeleteAnswerByKey(key)
	if err != nil {
		return err
	}
	models.SaveToHistory("delete", &models.Answer{Key: key})
	c.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
	return nil
}
