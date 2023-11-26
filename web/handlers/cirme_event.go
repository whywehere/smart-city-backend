package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/whywehere/smart-city-backend/config/db"
	"github.com/whywehere/smart-city-backend/utils/response"
	"github.com/whywehere/smart-city-backend/web/dal/query"
	"go.uber.org/zap"
)

func CrimeEventHandler(ctx *gin.Context) {
	opEvents := query.Use(db.DB).Event
	events, err := opEvents.WithContext(ctx).Where(opEvents.Type.Eq(2)).Find()
	if err != nil {
		zap.L().Error(err.Error())
		response.ResponseError(ctx, response.CodeServerBusy)
		return
	}
	response.ResponseSuccess(ctx, gin.H{
		"transArr": events,
	})
}
