package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SuspendExecutionPlanInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *SuspendExecutionPlanInstanceRequest) Invoke(client *sdk.Client) (resp *SuspendExecutionPlanInstanceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "SuspendExecutionPlanInstance", "", "")
	resp = &SuspendExecutionPlanInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type SuspendExecutionPlanInstanceResponse struct {
	responses.BaseResponse
	RequestId string
}
