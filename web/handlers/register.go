package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/whywehere/smart-city-backend/config/db"
	"github.com/whywehere/smart-city-backend/utils/response"
	"github.com/whywehere/smart-city-backend/web/dal/model"
	"github.com/whywehere/smart-city-backend/web/dal/query"
	"go.uber.org/zap"
)

func RegisterHandler(ctx *gin.Context) {
	var newUser *model.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		zap.L().Error(err.Error())
		response.ResponseError(ctx, response.CodeServerBusy)
	}

	if err := query.Use(db.DB).User.WithContext(ctx).Create(newUser); err != nil {
		zap.L().Error(err.Error())
		response.ResponseError(ctx, response.CodeServerBusy)
	}
	response.ResponseSuccess(ctx, "注册成功")
}
