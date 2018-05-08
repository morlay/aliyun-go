package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PushMessageToAndroidRequest struct {
	requests.RpcRequest
	AppKey      int64  `position:"Query" name:"AppKey"`
	TargetValue string `position:"Query" name:"TargetValue"`
	Title       string `position:"Query" name:"Title"`
	Body        string `position:"Query" name:"Body"`
	JobKey      string `position:"Query" name:"JobKey"`
	Target      string `position:"Query" name:"Target"`
}

func (req *PushMessageToAndroidRequest) Invoke(client *sdk.Client) (resp *PushMessageToAndroidResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "PushMessageToAndroid", "", "")
	resp = &PushMessageToAndroidResponse{}
	err = client.DoAction(req, resp)
	return
}

type PushMessageToAndroidResponse struct {
	responses.BaseResponse
	RequestId common.String
	MessageId common.String
}
