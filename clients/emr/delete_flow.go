package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteFlowRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DeleteFlowRequest) Invoke(client *sdk.Client) (resp *DeleteFlowResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteFlow", "", "")
	resp = &DeleteFlowResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteFlowResponse struct {
	responses.BaseResponse
	RequestId common.String
	Data      bool
}
