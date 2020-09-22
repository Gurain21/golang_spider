package check_error

import "fmt"

func CheckErr(err error) { //检查错误
	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("程序出现异常:", ins.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}
