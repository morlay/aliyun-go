package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PushMessageToAndroidRequest struct {
	AppKey      int64  `position:"Query" name:"AppKey"`
	TargetValue string `position:"Query" name:"TargetValue"`
	Title       string `position:"Query" name:"Title"`
	Body        string `position:"Query" name:"Body"`
	JobKey      string `position:"Query" name:"JobKey"`
	Target      string `position:"Query" name:"Target"`
}

func (r PushMessageToAndroidRequest) Invoke(client *sdk.Client) (response *PushMessageToAndroidResponse, err error) {
	req := struct {
		*requests.RpcRequest
		PushMessageToAndroidRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "PushMessageToAndroid", "", "")

	resp := struct {
		*responses.BaseResponse
		PushMessageToAndroidResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PushMessageToAndroidResponse

	err = client.DoAction(&req, &resp)
	return
}

type PushMessageToAndroidResponse struct {
	RequestId string
	MessageId string
}
