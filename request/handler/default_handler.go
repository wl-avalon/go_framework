package handler

type paramsItem interface{}
type IGetParam interface {
	getParams() map[string]paramsItem
}
type IResponseHandler interface {
	responseHandler() map[string]paramsItem
}

type Handler struct {
}