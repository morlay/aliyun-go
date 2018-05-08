package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteProjectRequest struct {
	requests.RpcRequest
	CsbId     int64 `position:"Query" name:"CsbId"`
	ProjectId int64 `position:"Query" name:"ProjectId"`
}

func (req *DeleteProjectRequest) Invoke(client *sdk.Client) (resp *DeleteProjectResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "DeleteProject", "CSB", "")
	resp = &DeleteProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteProjectResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
}
