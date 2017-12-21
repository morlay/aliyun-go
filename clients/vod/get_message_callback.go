package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetMessageCallbackRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (r GetMessageCallbackRequest) Invoke(client *sdk.Client) (response *GetMessageCallbackResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetMessageCallbackRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "GetMessageCallback", "", "")

	resp := struct {
		*responses.BaseResponse
		GetMessageCallbackResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetMessageCallbackResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetMessageCallbackResponse struct {
	RequestId       string
	MessageCallback GetMessageCallbackMessageCallback
}

type GetMessageCallbackMessageCallback struct {
	CallbackSwitch string
	CallbackURL    string
	EventTypeList  string
}
