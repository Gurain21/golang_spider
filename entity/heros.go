package entity

var HerosUrl = "https://pvp.qq.com/web201605/js/herolist.json" //json的网址
type Herolist struct {                                         //存储英雄数据
	//"ename": 105,
	//"cname": "廉颇",
	//"title": "正义爆轰",
	//"pay_type": 10,
	//"new_type": 0,
	//"hero_type": 3,
	//"skin_name": "正义爆轰|地狱岩魂"
	Ename     int
	Cname     string
	Title     string
	Pay_type  int
	New_type  int
	Hero_type int
	Skin_name string
}

var Heros []Herolist
