package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitFlowRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Conf            string `position:"Query" name:"Conf"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	FlowId          string `position:"Query" name:"FlowId"`
}

func (req *SubmitFlowRequest) Invoke(client *sdk.Client) (resp *SubmitFlowResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "SubmitFlow", "", "")
	resp = &SubmitFlowResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitFlowResponse struct {
	responses.BaseResponse
	RequestId string
	Id        string
}
