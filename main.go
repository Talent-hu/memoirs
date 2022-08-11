package main

import (
	"fmt"
	"memoirs/config"
	_ "memoirs/docs"
	"memoirs/global"
	"memoirs/model/auth"
	"memoirs/model/bank"
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
	//CreateTable()
	addr := fmt.Sprintf(":%s", global.AppConfig.NetConf.Port)
	server := router.InitRouter()
	_ = server.Run(addr)
}

func CreateTable() {
	_ = global.DB.AutoMigrate(
		//&auth.Area{},
		//&auth.SysDict{},
		//&auth.SysDictItem{},
		//&auth.User{},
		//&auth.UserRole{},
		//&auth.Role{},
		//&auth.RoleMenu{},
		&auth.Menu{},
		&bank.SubjectCategory{},
		&bank.QuestionLabel{},
		&bank.QuestionSelect{},
		&bank.QuestionJudge{},
		&bank.QuestionFillBack{},
		&bank.QuestionSimple{},
	)
}
