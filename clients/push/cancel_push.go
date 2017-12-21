package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CancelPushRequest struct {
	MessageId string `position:"Query" name:"MessageId"`
	AppKey    int64  `position:"Query" name:"AppKey"`
}

func (r CancelPushRequest) Invoke(client *sdk.Client) (response *CancelPushResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CancelPushRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "CancelPush", "", "")

	resp := struct {
		*responses.BaseResponse
		CancelPushResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CancelPushResponse

	err = client.DoAction(&req, &resp)
	return
}

type CancelPushResponse struct {
	RequestId string
}
