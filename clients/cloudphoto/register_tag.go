package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code      string
	Message   string
	RequestId string
	Action    string
}
