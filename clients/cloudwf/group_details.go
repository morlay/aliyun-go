package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GroupDetailsRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r GroupDetailsRequest) Invoke(client *sdk.Client) (response *GroupDetailsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GroupDetailsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GroupDetails", "", "")

	resp := struct {
		*responses.BaseResponse
		GroupDetailsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GroupDetailsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GroupDetailsResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
