package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RepairGroupApRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *RepairGroupApRequest) Invoke(client *sdk.Client) (resp *RepairGroupApResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "RepairGroupAp", "", "")
	resp = &RepairGroupApResponse{}
	err = client.DoAction(req, resp)
	return
}

type RepairGroupApResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
