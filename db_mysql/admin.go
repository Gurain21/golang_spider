package db_mysql

func QueryAdmin(name string, pwd string) (int, error) {
	//查询数据库中 的admin_name 和 admin_pwd 数据
	row := HeroDB.QueryRow("select count(admin_name) admin_num from wzry_admin where admin_name = ? and admin_pwd = ?",
		name, pwd)
	//创建admin_num 用于判断 数据库中是否存在 该管理员账户
	var admin_num int
	err := row.Scan(&admin_num)
	if err != nil {
		return 0, err
	}
	return admin_num, nil
}
