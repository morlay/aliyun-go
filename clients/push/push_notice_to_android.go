package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PushNoticeToAndroidRequest struct {
	ExtParameters string `position:"Query" name:"ExtParameters"`
	AppKey        int64  `position:"Query" name:"AppKey"`
	TargetValue   string `position:"Query" name:"TargetValue"`
	Title         string `position:"Query" name:"Title"`
	Body          string `position:"Query" name:"Body"`
	JobKey        string `position:"Query" name:"JobKey"`
	Target        string `position:"Query" name:"Target"`
}

func (r PushNoticeToAndroidRequest) Invoke(client *sdk.Client) (response *PushNoticeToAndroidResponse, err error) {
	req := struct {
		*requests.RpcRequest
		PushNoticeToAndroidRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "PushNoticeToAndroid", "", "")

	resp := struct {
		*responses.BaseResponse
		PushNoticeToAndroidResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PushNoticeToAndroidResponse

	err = client.DoAction(&req, &resp)
	return
}

type PushNoticeToAndroidResponse struct {
	RequestId string
	MessageId string
}
