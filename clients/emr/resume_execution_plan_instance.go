package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResumeExecutionPlanInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *ResumeExecutionPlanInstanceRequest) Invoke(client *sdk.Client) (resp *ResumeExecutionPlanInstanceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ResumeExecutionPlanInstance", "", "")
	resp = &ResumeExecutionPlanInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResumeExecutionPlanInstanceResponse struct {
	responses.BaseResponse
	RequestId string
}
