package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetGroupApRadioConfigProgressRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetGroupApRadioConfigProgressRequest) Invoke(client *sdk.Client) (response *GetGroupApRadioConfigProgressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetGroupApRadioConfigProgressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetGroupApRadioConfigProgress", "", "")

	resp := struct {
		*responses.BaseResponse
		GetGroupApRadioConfigProgressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetGroupApRadioConfigProgressResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetGroupApRadioConfigProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
