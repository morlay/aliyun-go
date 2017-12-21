package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetApgroupSsidConfigProgressRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetApgroupSsidConfigProgressRequest) Invoke(client *sdk.Client) (response *GetApgroupSsidConfigProgressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetApgroupSsidConfigProgressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetApgroupSsidConfigProgress", "", "")

	resp := struct {
		*responses.BaseResponse
		GetApgroupSsidConfigProgressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetApgroupSsidConfigProgressResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetApgroupSsidConfigProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
