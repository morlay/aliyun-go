package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetGroupApRepairProgressRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetGroupApRepairProgressRequest) Invoke(client *sdk.Client) (response *GetGroupApRepairProgressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetGroupApRepairProgressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetGroupApRepairProgress", "", "")

	resp := struct {
		*responses.BaseResponse
		GetGroupApRepairProgressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetGroupApRepairProgressResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetGroupApRepairProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
