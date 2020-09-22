package net_work

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"new_WZRY_spider/check_error"
	"new_WZRY_spider/entity"
)

func GetHerosSlices(url string) (Heros []entity.Herolist) {
	// 创建浏览器客户端
	client := http.Client{}
	//构建请求
	request, err := http.NewRequest("GET", url, nil)
	//检查错误
	check_error.CheckErr(err)
	//向客户端发起请求,获取 response
	response, err := client.Do(request)
	//检查错误
	check_error.CheckErr(err)
	//读取response,获取网页的信息
	htmlByte, err := ioutil.ReadAll(response.Body)
	//检查错误
	check_error.CheckErr(err)
	//将json格式转换成 heros切片 方便存储和读取
	err = json.Unmarshal(htmlByte, &entity.Heros)
	//检查错误
	check_error.CheckErr(err)
	Heros = entity.Heros
	return Heros
}
