package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteFlowProjectRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DeleteFlowProjectRequest) Invoke(client *sdk.Client) (resp *DeleteFlowProjectResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteFlowProject", "", "")
	resp = &DeleteFlowProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteFlowProjectResponse struct {
	responses.BaseResponse
	RequestId common.String
	Data      bool
}
