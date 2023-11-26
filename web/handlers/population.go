package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/whywehere/smart-city-backend/config/db"
	"github.com/whywehere/smart-city-backend/utils/response"
	"github.com/whywehere/smart-city-backend/web/dal/query"
	"go.uber.org/zap"
	"time"
)

func PopulationHandler(ctx *gin.Context) {
	year := time.Now().Year()
	opPeople := query.Use(db.DB).CityPersonNum
	population, err := opPeople.WithContext(ctx).Where(opPeople.Year.Gte(int32(year - 5))).Order(opPeople.Year).Find()
	if err != nil {
		zap.L().Error(err.Error())
		response.ResponseError(ctx, response.CodeServerBusy)
		return
	}
	response.ResponseSuccess(ctx, gin.H{
		"numArr": population,
	})
}
