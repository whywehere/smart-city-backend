package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/whywehere/smart-city-backend/config/db"
	"github.com/whywehere/smart-city-backend/utils/response"
	"github.com/whywehere/smart-city-backend/web/dal/query"
	"go.uber.org/zap"
)

func SourceHandler(ctx *gin.Context) {
	opSource := query.Use(db.DB).Source
	sources, err := opSource.WithContext(ctx).Find()
	if err != nil {
		zap.L().Error(err.Error())
		response.ResponseError(ctx, response.CodeServerBusy)
		return
	}
	response.ResponseSuccess(ctx, gin.H{
		"numArr": sources,
	})
}
