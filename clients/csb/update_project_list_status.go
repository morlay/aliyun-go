package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateProjectListStatusRequest struct {
	requests.RpcRequest
	Data  string `position:"Body" name:"Data"`
	CsbId int64  `position:"Query" name:"CsbId"`
}

func (req *UpdateProjectListStatusRequest) Invoke(client *sdk.Client) (resp *UpdateProjectListStatusResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "UpdateProjectListStatus", "CSB", "")
	resp = &UpdateProjectListStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateProjectListStatusResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
}
