package jaq

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AfsAppCheckRequest struct {
	requests.RpcRequest
	CallerName string `position:"Query" name:"CallerName"`
	Session    string `position:"Query" name:"Session"`
}

func (req *AfsAppCheckRequest) Invoke(client *sdk.Client) (resp *AfsAppCheckResponse, err error) {
	req.InitWithApiInfo("jaq", "2016-11-23", "AfsAppCheck", "", "")
	resp = &AfsAppCheckResponse{}
	err = client.DoAction(req, resp)
	return
}

type AfsAppCheckResponse struct {
	responses.BaseResponse
	ErrorCode int
	ErrorMsg  string
	Data      AfsAppCheckData
}

type AfsAppCheckData struct {
	SecondCheckResult int
}
