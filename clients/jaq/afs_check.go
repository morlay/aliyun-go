package jaq

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AfsCheckRequest struct {
	CallerName string `position:"Query" name:"CallerName"`
	Session    string `position:"Query" name:"Session"`
	Token      string `position:"Query" name:"Token"`
	Sig        string `position:"Query" name:"Sig"`
	Platform   int    `position:"Query" name:"Platform"`
	Scene      string `position:"Query" name:"Scene"`
}

func (r AfsCheckRequest) Invoke(client *sdk.Client) (response *AfsCheckResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AfsCheckRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("jaq", "2016-11-23", "AfsCheck", "", "")

	resp := struct {
		*responses.BaseResponse
		AfsCheckResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AfsCheckResponse

	err = client.DoAction(&req, &resp)
	return
}

type AfsCheckResponse struct {
	ErrorCode int
	ErrorMsg  string
	Data      bool
}
