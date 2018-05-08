package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteCredentialsListRequest struct {
	requests.RpcRequest
	Data        string `position:"Body" name:"Data"`
	IgnoreDauth string `position:"Query" name:"IgnoreDauth"`
	Force       string `position:"Query" name:"Force"`
}

func (req *DeleteCredentialsListRequest) Invoke(client *sdk.Client) (resp *DeleteCredentialsListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "DeleteCredentialsList", "CSB", "")
	resp = &DeleteCredentialsListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCredentialsListResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
}
