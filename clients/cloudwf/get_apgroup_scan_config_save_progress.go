package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetApgroupScanConfigSaveProgressRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetApgroupScanConfigSaveProgressRequest) Invoke(client *sdk.Client) (response *GetApgroupScanConfigSaveProgressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetApgroupScanConfigSaveProgressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetApgroupScanConfigSaveProgress", "", "")

	resp := struct {
		*responses.BaseResponse
		GetApgroupScanConfigSaveProgressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetApgroupScanConfigSaveProgressResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetApgroupScanConfigSaveProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
