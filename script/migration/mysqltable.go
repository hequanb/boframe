package main

import (
	"fmt"
	
	"boframe/models/usermodels"
	"boframe/settings"
	
	"boframe/settings/mysqlI"
)

func main() {
	// 初始化配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("init setting failed: %v \n", err)
		return
	}
	
	// 初始化MYSQL连接
	if err := mysqlI.Init(settings.Conf.MYSQLConfig); err != nil {
		fmt.Printf("init mysql failed: %v \n", err)
		return
	}
	defer mysqlI.Close()
	
	if err := mysqlI.DB().AutoMigrate(&usermodels.User{}); err != nil {
		fmt.Println("AutoMigrate failed: ", err)
		return
	}
	
	fmt.Println("done")
}
