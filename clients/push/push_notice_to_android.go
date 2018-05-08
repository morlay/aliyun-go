package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PushNoticeToAndroidRequest struct {
	requests.RpcRequest
	ExtParameters string `position:"Query" name:"ExtParameters"`
	AppKey        int64  `position:"Query" name:"AppKey"`
	TargetValue   string `position:"Query" name:"TargetValue"`
	Title         string `position:"Query" name:"Title"`
	Body          string `position:"Query" name:"Body"`
	JobKey        string `position:"Query" name:"JobKey"`
	Target        string `position:"Query" name:"Target"`
}

func (req *PushNoticeToAndroidRequest) Invoke(client *sdk.Client) (resp *PushNoticeToAndroidResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "PushNoticeToAndroid", "", "")
	resp = &PushNoticeToAndroidResponse{}
	err = client.DoAction(req, resp)
	return
}

type PushNoticeToAndroidResponse struct {
	responses.BaseResponse
	RequestId common.String
	MessageId common.String
}
