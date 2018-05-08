package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SuspendFlowRequest struct {
	requests.RpcRequest
	FlowInstanceId  string `position:"Query" name:"FlowInstanceId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *SuspendFlowRequest) Invoke(client *sdk.Client) (resp *SuspendFlowResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "SuspendFlow", "", "")
	resp = &SuspendFlowResponse{}
	err = client.DoAction(req, resp)
	return
}

type SuspendFlowResponse struct {
	responses.BaseResponse
	RequestId common.String
	Data      bool
}
