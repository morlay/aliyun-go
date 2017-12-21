package crm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteContacterRequest struct {
	requests.RpcRequest
	ContacterId int64 `position:"Query" name:"ContacterId"`
}

func (req *DeleteContacterRequest) Invoke(client *sdk.Client) (resp *DeleteContacterResponse, err error) {
	req.InitWithApiInfo("Crm", "2015-03-24", "DeleteContacter", "", "")
	resp = &DeleteContacterResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteContacterResponse struct {
	responses.BaseResponse
	Success       bool
	ResultCode    string
	ResultMessage string
}
