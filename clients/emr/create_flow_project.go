package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateFlowProjectRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Description     string `position:"Query" name:"Description"`
}

func (req *CreateFlowProjectRequest) Invoke(client *sdk.Client) (resp *CreateFlowProjectResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateFlowProject", "", "")
	resp = &CreateFlowProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateFlowProjectResponse struct {
	responses.BaseResponse
	RequestId common.String
	Id        common.String
}
