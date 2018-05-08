package hsm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Remark          string `position:"Query" name:"Remark"`
}

func (req *ModifyInstanceRequest) Invoke(client *sdk.Client) (resp *ModifyInstanceResponse, err error) {
	req.InitWithApiInfo("hsm", "2018-01-11", "ModifyInstance", "hsm", "")
	resp = &ModifyInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstanceResponse struct {
	responses.BaseResponse
	RequestId string
}
