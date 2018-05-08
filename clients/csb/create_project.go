package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      CreateProjectData
}

type CreateProjectData struct {
	Id common.Long
}
