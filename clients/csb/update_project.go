package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateProjectRequest struct {
	requests.RpcRequest
	Data  string `position:"Body" name:"Data"`
	CsbId int64  `position:"Query" name:"CsbId"`
}

func (req *UpdateProjectRequest) Invoke(client *sdk.Client) (resp *UpdateProjectResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "UpdateProject", "CSB", "")
	resp = &UpdateProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateProjectResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
}
