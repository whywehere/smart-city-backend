package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/whywehere/smart-city-backend/config/db"
	"github.com/whywehere/smart-city-backend/utils/response"
	"github.com/whywehere/smart-city-backend/web/dal/query"
	"go.uber.org/zap"
)

func SelectAllProblemsHandler(ctx *gin.Context) {
	opProblem := query.Use(db.DB).Problem
	problems, err := opProblem.WithContext(ctx).Group(opProblem.Type).Find()
	if err != nil {
		zap.L().Error(err.Error())
		response.ResponseError(ctx, response.CodeServerBusy)
	}
	response.ResponseSuccess(ctx, gin.H{
		"numArr": problems,
	})

}

func SelectAllByAreaHandler(ctx *gin.Context) {
	opProblem := query.Use(db.DB).Problem
	problems, err := opProblem.WithContext(ctx).Select(opProblem.ID, opProblem.Num.Sum()).Group(opProblem.ID).Find()
	if err != nil {
		zap.L().Error(err.Error())
		response.ResponseError(ctx, response.CodeServerBusy)
	}
	response.ResponseSuccess(ctx, gin.H{
		"numArr": problems,
	})
}

func SelectAirQualityHandler(ctx *gin.Context) {
	opCityAirQuality := query.Use(db.DB).CityAirQuality
	airArr, err := opCityAirQuality.WithContext(ctx).Find()
	if err != nil {
		zap.L().Error(err.Error())
		response.ResponseError(ctx, response.CodeServerBusy)
	}
	response.ResponseSuccess(ctx, gin.H{
		"airArr": airArr,
	})

}

// SelectSourceHandler 查询各个区域各种资源使用情况
func SelectSourceHandler(ctx *gin.Context) {
	opSource := query.Use(db.DB).Source
	sourceArr, err := opSource.WithContext(ctx).Select(opSource.Type, opSource.Num.GroupConcat()).Group(opSource.Type).Find()
	if err != nil {
		zap.L().Error(err.Error())
		response.ResponseError(ctx, response.CodeServerBusy)
	}
	response.ResponseSuccess(ctx, gin.H{
		"sourceArr": sourceArr,
	})

}

// SelectPersons 查询上海市总人口及增长率
func SelectPersons(ctx *gin.Context) {
	OpPerson := query.Use(db.DB).CityPersonNum
	personNums, err := OpPerson.WithContext(ctx).Order(OpPerson.Year).Find()
	if err != nil {
		zap.L().Error(err.Error())
		response.ResponseError(ctx, response.CodeServerBusy)
	}
	response.ResponseSuccess(ctx, gin.H{
		"infoArr": personNums,
	})
}

// SelectProblemsNumByYearHandler 按照年份查询四种问题的总数
func SelectProblemsNumByYearHandler(ctx *gin.Context) {
	opProblem := query.Use(db.DB).Problem
	problems, err := opProblem.WithContext(ctx).Select(opProblem.Type, opProblem.Num.Sum()).Group(opProblem.Type, opProblem.Year).Order(opProblem.Year).Find()
	if err != nil {
		zap.L().Error(err.Error())
		response.ResponseError(ctx, response.CodeServerBusy)
	}
	response.ResponseSuccess(ctx, gin.H{
		"numArr": problems,
	})
}
