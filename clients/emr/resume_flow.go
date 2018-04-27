package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResumeFlowRequest struct {
	requests.RpcRequest
	FlowInstanceId  string `position:"Query" name:"FlowInstanceId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *ResumeFlowRequest) Invoke(client *sdk.Client) (resp *ResumeFlowResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ResumeFlow", "", "")
	resp = &ResumeFlowResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResumeFlowResponse struct {
	responses.BaseResponse
	RequestId string
	Data      bool
}
