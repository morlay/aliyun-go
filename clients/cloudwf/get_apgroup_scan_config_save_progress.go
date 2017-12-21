package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetApgroupScanConfigSaveProgressRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *GetApgroupScanConfigSaveProgressRequest) Invoke(client *sdk.Client) (resp *GetApgroupScanConfigSaveProgressResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetApgroupScanConfigSaveProgress", "", "")
	resp = &GetApgroupScanConfigSaveProgressResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetApgroupScanConfigSaveProgressResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
