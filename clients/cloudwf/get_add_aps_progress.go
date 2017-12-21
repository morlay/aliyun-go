package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetAddApsProgressRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetAddApsProgressRequest) Invoke(client *sdk.Client) (response *GetAddApsProgressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetAddApsProgressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetAddApsProgress", "", "")

	resp := struct {
		*responses.BaseResponse
		GetAddApsProgressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetAddApsProgressResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetAddApsProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
