package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteFlowProjectUserRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	UserName        string `position:"Query" name:"UserName"`
}

func (req *DeleteFlowProjectUserRequest) Invoke(client *sdk.Client) (resp *DeleteFlowProjectUserResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteFlowProjectUser", "", "")
	resp = &DeleteFlowProjectUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteFlowProjectUserResponse struct {
	responses.BaseResponse
	RequestId string
	Data      bool
}
