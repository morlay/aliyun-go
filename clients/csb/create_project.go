package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateProjectRequest struct {
	requests.RpcRequest
	Data  string `position:"Body" name:"Data"`
	CsbId int64  `position:"Query" name:"CsbId"`
}

func (req *CreateProjectRequest) Invoke(client *sdk.Client) (resp *CreateProjectResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "CreateProject", "CSB", "")
	resp = &CreateProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateProjectResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      CreateProjectData
}

type CreateProjectData struct {
	Id int64
}
