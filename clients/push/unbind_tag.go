package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UnbindTagRequest struct {
	TagName   string `position:"Query" name:"TagName"`
	ClientKey string `position:"Query" name:"ClientKey"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	KeyType   string `position:"Query" name:"KeyType"`
}

func (r UnbindTagRequest) Invoke(client *sdk.Client) (response *UnbindTagResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UnbindTagRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "UnbindTag", "", "")

	resp := struct {
		*responses.BaseResponse
		UnbindTagResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UnbindTagResponse

	err = client.DoAction(&req, &resp)
	return
}

type UnbindTagResponse struct {
	RequestId string
}
