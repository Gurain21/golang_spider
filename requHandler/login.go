package requHandler

import (
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"new_WZRY_spider/check_error"
	"new_WZRY_spider/db_mysql"
	"new_WZRY_spider/entity"
)

/**
 * 处理管理员登录
 */
func Login(writer http.ResponseWriter, request *http.Request) {
	//解析并获取前端页面传递的管理员登录的相关信息，然后到数据库中查询
	//1、参数解析
	err := request.ParseForm()
	if err != nil {
		tmpt, _ := template.ParseFiles("./views/error.html")
		tmpt.Execute(writer, err.Error())
		return
	}
	//获取表单提交的用户名和密码
	adminName := request.FormValue("user_name")
	adminPwd := request.FormValue("user_pwd")
	//如果两个值有一个为空,则返回错误页面
	if adminName == "" || adminPwd == "" {
		tmpt, _ := template.ParseFiles("./views/error.html")
		tmpt.Execute(writer, "用户名或者密码为空，请检查后重新尝试")
		return
	}
	//根据这个两个值，到数据库中进行匹配
	admin_num, err := db_mysql.QueryAdmin(adminName, adminPwd)
	if err != nil {
		tmpt, _ := template.ParseFiles("./views/error.html")
		tmpt.Execute(writer, "用户名或者密码为空，请检查后重新尝试")
		return
	}
	//查询到了用户
	if admin_num > 0 {
		//解析模板文件
		temp, err := template.ParseFiles("./views/home.html")
		//检查错误
		check_error.CheckErr(err)
		//从数据库中获取heros切片
		Heros, err := db_mysql.QueryAllHeros()
		//检查错误
		check_error.CheckErr(err)
		//创建一个名为showdata结构体,将管理员名称,所有的heros信息,放到这个结构体中
		showData := entity.HomeShowData{AdminName: adminName, AllHeros: Heros}
		//将结构体showData传入模板
		temp.Execute(writer, showData)
	}
}
