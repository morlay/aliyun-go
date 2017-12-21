package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BatchChangeGroupApNameRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
}

func (r BatchChangeGroupApNameRequest) Invoke(client *sdk.Client) (response *BatchChangeGroupApNameResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BatchChangeGroupApNameRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "BatchChangeGroupApName", "", "")

	resp := struct {
		*responses.BaseResponse
		BatchChangeGroupApNameResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BatchChangeGroupApNameResponse

	err = client.DoAction(&req, &resp)
	return
}

type BatchChangeGroupApNameResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
