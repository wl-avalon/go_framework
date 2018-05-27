package rpc

import (
	"net/http"
	"strings"
	"io/ioutil"
	"time"
	"github.com/gin-gonic/gin/json"
	"go_framework/request/models"
)

type RpcConfig struct {
	Domain map[string]string
}

var rpcConfig *RpcConfig

func InitRpc(domainMap map[string]string){
	rpcConfig = new(RpcConfig)
	rpcConfig.Domain = domainMap
}

func HttpRequest(request *models.Request, response *models.Response) error {
	time.Sleep(time.Second * 2)
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	url 			:= rpcConfig.Domain[request.GetServiceName()] + request.GetUrl()
	jsonParams,err	:= json.Marshal(request.GetParams())
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonParams)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	responseChan := response.GetResponseChannel()
	responseChan <- string(body)
	return nil
}