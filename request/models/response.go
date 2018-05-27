package models

import (
	"errors"
	"encoding/json"
)

type errorContent struct{
	errorCode			string
	errorMessage		string
	errorUserMessage	string
}

type data interface{}

type Response struct{
	init		bool
	httpCode	string
	error		errorContent
	data		data
	responseCh	chan interface{}
}

func NewResponse() (*Response, error) {
	response 			:= new(Response)
	response.init 		= false
	response.responseCh	= make(chan interface{})
	return response, nil
}

func (response *Response) GetResponseChannel() chan interface{} {
	return response.responseCh
}

func (response *Response) Get(key string) (interface{}, error) {
	if !response.init {
		err := response.waitResponse()
		if err != nil {
			return nil, err
		}
	}
	value, ok := response.data.(map[string]interface{})[key]
	if !ok {
		err := errors.New("返回值中不存在" + key + "这个key")
		return nil, err
	}
	return value, nil
}

func (response *Response) waitResponse() error {
	select {
	case body := <-response.responseCh:
		response.init = true
		err := json.Unmarshal([]byte(body.(string)), &response.data)
		if err != nil {
			err := errors.New("返回值json反序列化失败,body is:" + body.(string))
			return err
		}
	}
	return nil
}