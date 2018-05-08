package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteApgroupConfigRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *DeleteApgroupConfigRequest) Invoke(client *sdk.Client) (resp *DeleteApgroupConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeleteApgroupConfig", "", "")
	resp = &DeleteApgroupConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteApgroupConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
