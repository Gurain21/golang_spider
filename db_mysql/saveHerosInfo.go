package db_mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"new_WZRY_spider/check_error"
	"new_WZRY_spider/entity"
)

//创建一个变量指针,用于进行数据库操作
var HeroDB *sql.DB

//向数据库存储数据
func SaveHero2MySql(hero entity.Herolist) (error) {
	//打开数据库
	OpenDB()
	//向数据库添加数据
	_, err := HeroDB.Exec("insert into "+
		"heros("+
		"ename,cname,title,pay_type,new_type,hero_type,skin_name)"+
		"values("+
		"?, ?, ?, ?, ?, ?, ?)",
		hero.Ename, hero.Cname, hero.Title, hero.Pay_type, hero.New_type, hero.Hero_type, hero.Skin_name)
	return err

}
func QueryHerosNum() (int, error) {
	//1、连接并打开数据库
	OpenDB()

	//查询数据库内数据 ename 的总行数
	rows := HeroDB.QueryRow("select count(ename) recordnum from heros")
	var recordNum int
	err := rows.Scan(&recordNum)
	if err != nil {
		return 0, err
	}
	return recordNum, nil
}

//读取数据库中的数据,并返回一个heros切片
func QueryAllHeros() ([]entity.Herolist, error) {
	//查询数据
	rows, err := HeroDB.Query("select * from heros")
	if err != nil {
		return nil, err
	}
	heros := make([]entity.Herolist, 0)
	//迭代查询数据
	for rows.Next() {
		var hero entity.Herolist
		//读取每一行数据,并将它们写入到 hero 中
		err = rows.Scan(&hero.Ename, &hero.Cname, &hero.Title, &hero.Pay_type, &hero.New_type, &hero.Hero_type, &hero.Skin_name)
		if err != nil {
			return nil, err
		}
		//将每一个 hero 存储到切片heros中
		heros = append(heros, hero)
	}
	//返回heros 和 nil
	return heros, nil
}
func OpenDB() {
	//如果数据库已经打开了 ,直接return
	if HeroDB != nil {
		return
	}
	//打开数据库
	database, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/wzry?charset=utf8")
	//检查错误
	check_error.CheckErr(err)
	//给HeroDB传值
	HeroDB = database

}
