package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RegisterTagRequest struct {
	requests.RpcRequest
	StoreName string `position:"Query" name:"StoreName"`
	Text      string `position:"Query" name:"Text"`
	TagKey    string `position:"Query" name:"TagKey"`
	Lang      string `position:"Query" name:"Lang"`
}

func (req *RegisterTagRequest) Invoke(client *sdk.Client) (resp *RegisterTagResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "RegisterTag", "cloudphoto", "")
	resp = &RegisterTagResponse{}
	err = client.DoAction(req, resp)
	return
}

type RegisterTagResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
}
