package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveTagRequest struct {
	TagName string `position:"Query" name:"TagName"`
	AppKey  int64  `position:"Query" name:"AppKey"`
}

func (r RemoveTagRequest) Invoke(client *sdk.Client) (response *RemoveTagResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveTagRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "RemoveTag", "", "")

	resp := struct {
		*responses.BaseResponse
		RemoveTagResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveTagResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveTagResponse struct {
	RequestId string
}
