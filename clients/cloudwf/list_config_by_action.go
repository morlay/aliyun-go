package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListConfigByActionRequest struct {
	SearchName string `position:"Query" name:"SearchName"`
	Limit      int    `position:"Query" name:"Limit"`
	ActionName string `position:"Query" name:"ActionName"`
}

func (r ListConfigByActionRequest) Invoke(client *sdk.Client) (response *ListConfigByActionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListConfigByActionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListConfigByAction", "", "")

	resp := struct {
		*responses.BaseResponse
		ListConfigByActionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListConfigByActionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListConfigByActionResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
