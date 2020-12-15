//Package controllers ... handle http requests for the Answer
package controllers

import (
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
// @Router /answer/{key} [post]
func CreateAnswerByKey(c *gin.Context) error {
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
	return nil
}