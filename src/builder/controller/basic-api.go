package controller

import (
	"builder/constant"
	"builder/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

	// health
	addRequestMapping(
		RequestMapper{
			Method:  "GET",
			Path:    "/health",
			Request: health,
		},
	)
}

/*
	Request Mapping Functions
*/

// health
// @Summary health check api
// @Description builder의 health를 체크할 목적의 api
// @Name health
// @Produce  json
// @Router /health [get]
// @Success 200 {object} model.BasicResult
func health(c *gin.Context) {
	c.JSON(http.StatusOK, &model.BasicResult{
		Code:    constant.ResultSuccess,
		Message: "taco-registry-builder is healthy",
	})
}