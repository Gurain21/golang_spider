package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"new_WZRY_spider/entity"

	//_ "github.com/mahonia-master"
	"net/http"
	"new_WZRY_spider/check_error"
	"new_WZRY_spider/db_mysql"
	"new_WZRY_spider/net_work"
	"new_WZRY_spider/requHandler"
)

func main() {
	//	//打开数据库
	db_mysql.OpenDB()
	// 程序结束前关闭数据库
	defer db_mysql.HeroDB.Close()
	//打印数据库已存储的数据行数
	recordNum, err := db_mysql.QueryHerosNum()
	fmt.Println("数据库中已存储了", recordNum, "行数据")
	check_error.CheckErr(err)
	// 如果数据库中未存储数据  则获取 并 遍历 切片 Heros,将每一条数据存储入库
	if recordNum <= 0 {
		//开启事务  事务操作的特点：将批量操作作为一个整体，要么全部成功，要么回滚，执行失败。
		tx,err :=db_mysql.HeroDB.Begin()
		check_error.CheckErr(err)
		Heros := net_work.GetHerosSlices(entity.HerosUrl)
		for _, hero := range Heros { //遍历切片Heros
			err := db_mysql.SaveHero2MySql(hero) //向数据库里添加数据
			if err != nil {
						//回滚事务 特点：将批量操作作为一个整体，要么全部成功，要么回滚，执行失败。
						tx.Rollback()
						check_error.CheckErr(err)
					}
			//提交事务
			tx.Commit()
		}
	} else { //数据库中已经存储好了数据 展示登录页面
		http.HandleFunc("/", requHandler.Index)
		http.HandleFunc("/login", requHandler.Login)
		//静态文件服务 向互联网暴露登录页面所需要的图片和js文件
		http.Handle("/static/img/", http.StripPrefix("/static/img", http.FileServer(http.Dir("./static/img/"))))
		http.Handle("/static/js/", http.StripPrefix("/static/js", http.FileServer(http.Dir("./static/js/"))))
		fmt.Println("正在监听127.0.0.1:9000端口")
		//创建一个客户端口
		err = http.ListenAndServe("127.0.0.1:9000", nil)
		check_error.CheckErr(err)
	}
}
