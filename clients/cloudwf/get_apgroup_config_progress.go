package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetApgroupConfigProgressRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetApgroupConfigProgressRequest) Invoke(client *sdk.Client) (response *GetApgroupConfigProgressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetApgroupConfigProgressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetApgroupConfigProgress", "", "")

	resp := struct {
		*responses.BaseResponse
		GetApgroupConfigProgressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetApgroupConfigProgressResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetApgroupConfigProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
