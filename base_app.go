package go_framework

import (
	"go_framework/request"
	"fmt"
	"time"
	"go_framework/components"
)

func init() {

}

func Run() error {
	err := InitConfig("/Users/avalonspace/Documents/Source Code/go-path/src/go_framework/server.ini")
	if err != nil {
		return err
	}

	//test()
	for {
		go test()
		time.Sleep(1 * time.Millisecond)
	}

	return nil
}

func test(){
	noteFrequency := map[string]interface{} {
		"C0": 16.35,
	}
	//1 开始请求
	fmt.Println("开始请求")
	response, err:= request.AsyncReq("passport", "/study-palace/passport/commit/login", noteFrequency)
	if err != nil {
		fmt.Println("请求失败,错误信息:" + err.Error())
	}

	//2 代码请求结束,协程请求中
	fmt.Println("代码请求结束,协程请求中")
	C0, err := response.Get("C0")

	//3 请求结束，拿到返回值
	if err != nil {
		fmt.Println("请求失败,错误信息:" + err.Error())
	}else{
		C0string, err := components.TurnNumberToString(C0)
		if err != nil {
			fmt.Println("返回结果转为字符串失败")
		}else{
			fmt.Println("返回结果:" + C0string)
		}
	}
}