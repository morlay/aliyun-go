package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetMessageCallbackRequest struct {
	CallbackType         string `position:"Query" name:"CallbackType"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	CallbackSwitch       string `position:"Query" name:"CallbackSwitch"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EventTypeList        string `position:"Query" name:"EventTypeList"`
	CallbackURL          string `position:"Query" name:"CallbackURL"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (r SetMessageCallbackRequest) Invoke(client *sdk.Client) (response *SetMessageCallbackResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetMessageCallbackRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "SetMessageCallback", "", "")

	resp := struct {
		*responses.BaseResponse
		SetMessageCallbackResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetMessageCallbackResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetMessageCallbackResponse struct {
	RequestId string
}
