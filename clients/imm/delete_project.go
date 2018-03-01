package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteProjectRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
}

func (req *DeleteProjectRequest) Invoke(client *sdk.Client) (resp *DeleteProjectResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DeleteProject", "imm", "")
	resp = &DeleteProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteProjectResponse struct {
	responses.BaseResponse
	RequestId string
}
