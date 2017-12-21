package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}
