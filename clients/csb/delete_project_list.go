package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteProjectListRequest struct {
	requests.RpcRequest
	Data  string `position:"Body" name:"Data"`
	CsbId int64  `position:"Query" name:"CsbId"`
}

func (req *DeleteProjectListRequest) Invoke(client *sdk.Client) (resp *DeleteProjectListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "DeleteProjectList", "CSB", "")
	resp = &DeleteProjectListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteProjectListResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
}
