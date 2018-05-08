package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyResourcePoolSchedulerTypeRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SchedulerType   string `position:"Query" name:"SchedulerType"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *ModifyResourcePoolSchedulerTypeRequest) Invoke(client *sdk.Client) (resp *ModifyResourcePoolSchedulerTypeResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyResourcePoolSchedulerType", "", "")
	resp = &ModifyResourcePoolSchedulerTypeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyResourcePoolSchedulerTypeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
