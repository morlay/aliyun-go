package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type KillFlowRequest struct {
	requests.RpcRequest
	FlowInstanceId  string `position:"Query" name:"FlowInstanceId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *KillFlowRequest) Invoke(client *sdk.Client) (resp *KillFlowResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "KillFlow", "", "")
	resp = &KillFlowResponse{}
	err = client.DoAction(req, resp)
	return
}

type KillFlowResponse struct {
	responses.BaseResponse
	RequestId common.String
	Data      bool
}
