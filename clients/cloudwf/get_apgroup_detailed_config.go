package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetApgroupDetailedConfigRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *GetApgroupDetailedConfigRequest) Invoke(client *sdk.Client) (resp *GetApgroupDetailedConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetApgroupDetailedConfig", "", "")
	resp = &GetApgroupDetailedConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetApgroupDetailedConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
