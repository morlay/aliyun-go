package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetApgroupPortalConfigProgressRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetApgroupPortalConfigProgressRequest) Invoke(client *sdk.Client) (response *GetApgroupPortalConfigProgressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetApgroupPortalConfigProgressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetApgroupPortalConfigProgress", "", "")

	resp := struct {
		*responses.BaseResponse
		GetApgroupPortalConfigProgressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetApgroupPortalConfigProgressResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetApgroupPortalConfigProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
