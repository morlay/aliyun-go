package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteJobExecutionPlanParamRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
	Id              int64 `position:"Query" name:"Id"`
}

func (req *DeleteJobExecutionPlanParamRequest) Invoke(client *sdk.Client) (resp *DeleteJobExecutionPlanParamResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteJobExecutionPlanParam", "", "")
	resp = &DeleteJobExecutionPlanParamResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteJobExecutionPlanParamResponse struct {
	responses.BaseResponse
	RequestId common.String
}
