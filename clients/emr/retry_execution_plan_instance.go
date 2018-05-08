package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RetryExecutionPlanInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Arguments       string `position:"Query" name:"Arguments"`
	Id              string `position:"Query" name:"Id"`
	RerunFail       string `position:"Query" name:"RerunFail"`
}

func (req *RetryExecutionPlanInstanceRequest) Invoke(client *sdk.Client) (resp *RetryExecutionPlanInstanceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "RetryExecutionPlanInstance", "", "")
	resp = &RetryExecutionPlanInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RetryExecutionPlanInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
