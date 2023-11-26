package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/whywehere/smart-city-backend/config/db"
	"github.com/whywehere/smart-city-backend/utils/jwt"
	"github.com/whywehere/smart-city-backend/utils/response"
	"github.com/whywehere/smart-city-backend/web/dal/model"
	"github.com/whywehere/smart-city-backend/web/dal/query"
	"go.uber.org/zap"
)

func LoginHandler(ctx *gin.Context) {
	var loginUser *model.User
	if err := ctx.ShouldBindJSON(&loginUser); err != nil {
		zap.L().Error(err.Error())
		response.ResponseError(ctx, response.CodeServerBusy)
		return
	}
	OpUser := query.Use(db.DB).User
	isExist, err := OpUser.WithContext(ctx).Where(OpUser.Acount.Eq(loginUser.Acount), OpUser.Password.Eq(loginUser.Password)).Count()
	if err != nil {
		zap.L().Error(err.Error())
		response.ResponseError(ctx, response.CodeServerBusy)
		return
	}
	if isExist == 0 {
		response.ResponseError(ctx, response.CodeUserNotExist)
		return
	}

	accessToken, refreshToken, err := jwt.GenToken(uint64(loginUser.ID), loginUser.Username)
	response.ResponseSuccess(ctx, gin.H{
		"user_id":       loginUser.ID,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}
