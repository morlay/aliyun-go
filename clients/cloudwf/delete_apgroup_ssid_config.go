package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteApgroupSsidConfigRequest struct {
	requests.RpcRequest
	ApgroupId int64 `position:"Query" name:"ApgroupId"`
	Id        int64 `position:"Query" name:"Id"`
}

func (req *DeleteApgroupSsidConfigRequest) Invoke(client *sdk.Client) (resp *DeleteApgroupSsidConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeleteApgroupSsidConfig", "", "")
	resp = &DeleteApgroupSsidConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteApgroupSsidConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
