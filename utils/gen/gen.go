package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"

	"strings"
)

func main() {

	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	gormdb, _ := gorm.Open(mysql.Open("root:cq194068971@(127.0.0.1:3306)/smart_city?charset=utf8mb4&parseTime=True&loc=Local"))
	g := gen.NewGenerator(gen.Config{

		OutPath: "./web/dal/query",

		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,

		/* Mode: gen.WithoutContext,*/

		//if you want the nullable field generation property to be pointer type, set FieldNullable true

		/* FieldNullable: true,*/

	})
	g.WithTableNameStrategy(func(tableName string) (targetTableName string) {

		if strings.HasPrefix(tableName, "_") { //忽略下划线开头的表

			return ""

		}

		return tableName

	})

	g.UseDB(gormdb)

	g.ApplyBasic(g.GenerateAllTable()...) //同步数据库所有表

	g.Execute()

}
