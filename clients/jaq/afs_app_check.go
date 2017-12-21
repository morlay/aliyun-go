package jaq

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AfsAppCheckRequest struct {
	CallerName string `position:"Query" name:"CallerName"`
	Session    string `position:"Query" name:"Session"`
}

func (r AfsAppCheckRequest) Invoke(client *sdk.Client) (response *AfsAppCheckResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AfsAppCheckRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("jaq", "2016-11-23", "AfsAppCheck", "", "")

	resp := struct {
		*responses.BaseResponse
		AfsAppCheckResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AfsAppCheckResponse

	err = client.DoAction(&req, &resp)
	return
}

type AfsAppCheckResponse struct {
	ErrorCode int
	ErrorMsg  string
	Data      AfsAppCheckData
}

type AfsAppCheckData struct {
	SecondCheckResult int
}
