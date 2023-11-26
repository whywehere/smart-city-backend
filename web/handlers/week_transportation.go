package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/whywehere/smart-city-backend/config/db"
	"github.com/whywehere/smart-city-backend/utils/response"
	"github.com/whywehere/smart-city-backend/web/dal/query"
	"go.uber.org/zap"
	"time"
)

func WeekTransportationHandler(ctx *gin.Context) {
	now := time.Now()
	week := now.AddDate(0, 0, -7)
	opEvents := query.Use(db.DB).Event
	events, err := opEvents.WithContext(ctx).Where(opEvents.Time.Gt(week)).Find()
	if err != nil {
		zap.L().Error(err.Error())
		response.ResponseError(ctx, response.CodeServerBusy)
		return
	}
	response.ResponseSuccess(ctx, gin.H{
		"transArr": events,
	})
}
