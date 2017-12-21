package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PushNoticeToiOSRequest struct {
	requests.RpcRequest
	ExtParameters string `position:"Query" name:"ExtParameters"`
	ApnsEnv       string `position:"Query" name:"ApnsEnv"`
	AppKey        int64  `position:"Query" name:"AppKey"`
	TargetValue   string `position:"Query" name:"TargetValue"`
	Title         string `position:"Query" name:"Title"`
	Body          string `position:"Query" name:"Body"`
	JobKey        string `position:"Query" name:"JobKey"`
	Target        string `position:"Query" name:"Target"`
}

func (req *PushNoticeToiOSRequest) Invoke(client *sdk.Client) (resp *PushNoticeToiOSResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "PushNoticeToiOS", "", "")
	resp = &PushNoticeToiOSResponse{}
	err = client.DoAction(req, resp)
	return
}

type PushNoticeToiOSResponse struct {
	responses.BaseResponse
	RequestId string
	MessageId string
}
