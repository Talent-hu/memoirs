package main

import (
	"fmt"
	"memoirs/config"
	_ "memoirs/docs"
	"memoirs/global"
	"memoirs/router"
)

// @title 后台系统
// @version 1.0.0
// @description gin开发的后台管理系统
func main() {
	err := config.NewApp()
	if err != nil {
		panic(err)
	}
	//_ = global.DB.AutoMigrate(
	//	&model.SubjectCategory{},
	//	&model.QuestionLabel{},
	//	&model.QuestionSelect{},
	//	&model.QuestionJudge{},
	//	&model.QuestionFillBack{},
	//	&model.QuestionSimple{},
	//)
	addr := fmt.Sprintf(":%s", global.AppConfig.NetConf.Port)
	server := router.InitRouter()
	_ = server.Run(addr)
}
