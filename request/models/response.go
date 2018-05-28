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
	init			bool
	httpCode		string
	err				error
	data			data
	responseCh		chan interface{}
	responseErrCh	chan error
}

func NewResponse() (*Response, error) {
	response 				:= new(Response)
	response.init			= false
	response.responseCh		= make(chan interface{})
	response.responseErrCh	= make(chan error)
	return response, nil
}

func (response *Response) GetResponseChannel() chan interface{} {
	return response.responseCh
}

func (response *Response) GetResponseErrorChannel() chan error {
	return response.responseErrCh
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
		err :=json.Unmarshal([]byte(body.(string)), &response.data)
		if err != nil {
			response.err = errors.New("返回值无法按json解析")
			return err
		}
		response.init = true
		return nil
	case err := <-response.responseErrCh:
		response.err = err
		return err
	}
}