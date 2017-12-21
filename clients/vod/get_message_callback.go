package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetMessageCallbackRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *GetMessageCallbackRequest) Invoke(client *sdk.Client) (resp *GetMessageCallbackResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "GetMessageCallback", "", "")
	resp = &GetMessageCallbackResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetMessageCallbackResponse struct {
	responses.BaseResponse
	RequestId       string
	MessageCallback GetMessageCallbackMessageCallback
}

type GetMessageCallbackMessageCallback struct {
	CallbackSwitch string
	CallbackURL    string
	EventTypeList  string
}
