package requHandler

import (
	"html/template"
	"net/http"
)

//处理请求的函数
func Index(writer http.ResponseWriter, request *http.Request) {
	//writer.Write([]byte("hello every one"))//views/admin.html
	temp, err := template.ParseFiles("./views/index.html") //E:\GoProjects\src\0604work_Spider\views\admin.html
	if err != nil {
		//writer.Write([]byte(err.Error()))
		errorTmp, _ := template.ParseFiles("./views/error.html")
		errorTmp.Execute(writer, err.Error())
		return
	}
	//解析模板正常
	temp.Execute(writer, nil)
}
