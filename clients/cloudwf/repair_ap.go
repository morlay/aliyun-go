package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RepairApRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *RepairApRequest) Invoke(client *sdk.Client) (resp *RepairApResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "RepairAp", "", "")
	resp = &RepairApResponse{}
	err = client.DoAction(req, resp)
	return
}

type RepairApResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
