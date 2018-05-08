package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type KillExecutionPlanInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *KillExecutionPlanInstanceRequest) Invoke(client *sdk.Client) (resp *KillExecutionPlanInstanceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "KillExecutionPlanInstance", "", "")
	resp = &KillExecutionPlanInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type KillExecutionPlanInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
