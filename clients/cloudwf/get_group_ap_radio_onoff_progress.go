package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetGroupApRadioOnoffProgressRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetGroupApRadioOnoffProgressRequest) Invoke(client *sdk.Client) (response *GetGroupApRadioOnoffProgressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetGroupApRadioOnoffProgressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetGroupApRadioOnoffProgress", "", "")

	resp := struct {
		*responses.BaseResponse
		GetGroupApRadioOnoffProgressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetGroupApRadioOnoffProgressResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetGroupApRadioOnoffProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
