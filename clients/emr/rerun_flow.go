package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RerunFlowRequest struct {
	requests.RpcRequest
	FlowInstanceId  string `position:"Query" name:"FlowInstanceId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	ReRunFail       string `position:"Query" name:"ReRunFail"`
}

func (req *RerunFlowRequest) Invoke(client *sdk.Client) (resp *RerunFlowResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "RerunFlow", "", "")
	resp = &RerunFlowResponse{}
	err = client.DoAction(req, resp)
	return
}

type RerunFlowResponse struct {
	responses.BaseResponse
	RequestId common.String
	Data      bool
}
