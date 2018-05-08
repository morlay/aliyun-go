package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyExecutionPlanBasicInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Id              string `position:"Query" name:"Id"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *ModifyExecutionPlanBasicInfoRequest) Invoke(client *sdk.Client) (resp *ModifyExecutionPlanBasicInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyExecutionPlanBasicInfo", "", "")
	resp = &ModifyExecutionPlanBasicInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyExecutionPlanBasicInfoResponse struct {
	responses.BaseResponse
	RequestId common.String
}
