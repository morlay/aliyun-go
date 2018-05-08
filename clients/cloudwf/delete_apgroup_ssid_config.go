package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
