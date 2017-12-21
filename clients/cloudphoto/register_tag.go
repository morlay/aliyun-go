package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RegisterTagRequest struct {
	StoreName string `position:"Query" name:"StoreName"`
	Text      string `position:"Query" name:"Text"`
	TagKey    string `position:"Query" name:"TagKey"`
	Lang      string `position:"Query" name:"Lang"`
}

func (r RegisterTagRequest) Invoke(client *sdk.Client) (response *RegisterTagResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RegisterTagRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "RegisterTag", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		RegisterTagResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RegisterTagResponse

	err = client.DoAction(&req, &resp)
	return
}

type RegisterTagResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}
