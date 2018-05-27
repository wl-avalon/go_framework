package models

import (
	"go_framework/request/handler"
	"errors"
	"io"
	"crypto/md5"
	"github.com/gin-gonic/gin/json"
	"fmt"
)

type requestOptions struct{
	retry				int	//重试次数
	retryTimes			int	//当前重试次数
	requestMode			int	//请求方式 0:惰性请求 1:同步请求
	timeoutMS			int	//超时时间，单位:毫秒
	connectTimeoutMS	int //连接超时时间，单位:毫秒
}

type Request struct {
	url 		string					//实际要请求的URL
	serviceName	string					//服务名称
	params		map[string]interface{}	//请求参数
	handler		*handler.Handler		//handler结构体，封装参数或者处理返回值
	options		*requestOptions			//请求的一些选项
	hash		string					//请求哈希值
	hashOffset	int						//哈希冲突偏移值
}

func (req *Request) GetUrl() string {
	return req.url
}

func (req *Request) GetServiceName() string {
	return req.serviceName
}

func (req *Request) GetParams() map[string]interface{} {
	return req.params
}


func NewRequest(serviceName string, url string, params map[string]interface{}) (*Request, error){
	request 			:= new (Request)
	request.serviceName = serviceName
	request.params		= params
	request.url 		= url
	request.hashOffset  = 0
	request.options		= new(requestOptions)
	request.handler		= new(handler.Handler)

	var err error
	request.hash, err	= hashParams(params)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func hashParams(params map[string]interface{}) (string, error){
	//1 按map的Key升序排序
	//sortedKey := make([]string, 0)
	//for key := range params {
	//	sortedKey = append(sortedKey, key)
	//}
	//sort.Strings(sortedKey)

	//2 json后生成哈希 TODO:这里可能由于map乱序导致问题，后续需要保证严格顺序
	resString, err := json.Marshal(params)
	if err != nil {
		error := errors.New("请求参数生成哈希源json串时失败")
		return "", error
	}

	//3 生成哈希值
	h := md5.New()
	io.WriteString(h, string(resString))
	hashString := fmt.Sprintf("%x\n", h.Sum(nil))
	return hashString, nil
}