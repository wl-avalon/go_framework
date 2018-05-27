package request

import (
	"go_framework/request/models"
	"go_framework/request/rpc"
)

/**
 * 异步请求
 * @param serviceName 	业务名称
 * @param interfaceName	接口名称
 * @param params     	业务参数
 * @param response		响应对象
 * @return error		异常
 */
func AsyncReq(serviceName string, interfaceName string, params map[string]interface{}) (*models.Response, error){
	//1 创建请求对象
	request, reqErr := models.NewRequest(serviceName, interfaceName, params)
	if reqErr != nil {
		return nil, reqErr
	}

	//2 创建响应对象
	response, resErr := models.NewResponse()
	if resErr != nil {
		return nil, resErr
	}

	//3 发送请求
	go rpc.HttpRequest(request, response)
	return response, nil
}